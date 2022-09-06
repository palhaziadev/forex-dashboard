package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/palhaziadev/forex-dashboard/pkg/config"
	mqUtils "github.com/palhaziadev/forex-dashboard/pkg/rabbitmq"
	"github.com/palhaziadev/forex-dashboard/server/internal/controller"
	"github.com/streadway/amqp"
	"golang.org/x/net/websocket"
)

// TODO dependency injection?
type testData struct {
	Number int
}

var messageChan chan *testData = make(chan *testData)

var healthCheckPath = "health"

type healthStatusResponse struct {
	Status string `json:"status"`
}

const (
	healthSatusOk string = "healthy3"
)

type Server struct {
	controller *controller.Controller
}

func NewServer() *Server {
	return &Server{
		controller: controller.NewController(), // TODO maybe add controller as a param from main.go
	}
}

func (server *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	var healthCheck = healthStatusResponse{Status: healthSatusOk}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(string(healthCheck.Status) + " server")
		log.Println(string(j))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, string(j))
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (server *Server) consumerHandler(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		fmt.Printf("Error occurred in RMQ consumer: %s", err)
	}

	fmt.Printf("Message received on '%s' queue: %s", queue, string(msg.Body))
	data := testData{}
	err = json.Unmarshal(msg.Body, &data)
	if err != nil {
		fmt.Printf("Error occurred in consumer Unmarshal: %s", err)
	}

	// go func() { messageChan <- &data }()
	select {
	case messageChan <- &data:
		//if can't write to the channel drop the message
		//TODO check if first or last message is sent when the frontend connects
	default:
	}
}

func (server *Server) MqConsumerTest() {
	connectionString := config.GetEnvVar("RMQ_URL")

	exampleQueue := mqUtils.RMQConsumer{
		Queue:            config.GetEnvVar("TEST_QUEUE"),
		ConnectionString: connectionString,
		MsgHandler:       server.consumerHandler,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()
	// Multiple listeners can be specified here

	<-forever
}

func (server *Server) testWebsocketHandler(ws *websocket.Conn) {

	// if the server recieves error from the client close the ws server
	done := make(chan struct{})
	go func(c *websocket.Conn) {
		for {
			var msg struct{}
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				fmt.Printf("************************recieve error %s\n", msg)
				break
			}
			fmt.Printf("recieved message %s\n", msg)
		}
		close(done)
	}(ws)
loop:
	for {
		select {
		case <-done:
			fmt.Println("connection was closed, lets break out of here")
			break loop
		default:
			msg := <-messageChan
			b, err := json.Marshal(msg)
			fmt.Printf("Message received: '%s'\n", b)
			websocket.JSON.Send(ws, msg)
			if err != nil {
				fmt.Printf("--------------handler error: '%s'\n", err)
				break
			}
		}
	}
	fmt.Println("closing the connection")
	defer ws.Close()
}

func (server *Server) RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		healthCheckHandler := http.HandlerFunc(server.handleHealthCheck)
		http.Handle("/websocket", websocket.Handler(server.testWebsocketHandler))
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, healthCheckPath), healthCheckHandler)
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, server.controller.Path), http.HandlerFunc(server.controller.HandleForexList))
	}
	return returnFunction
}
