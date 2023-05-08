package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"

	config "github.com/palhaziadev/forex-dashboard/pkg/config"
	"github.com/palhaziadev/forex-dashboard/pkg/models"
	mqUtils "github.com/palhaziadev/forex-dashboard/pkg/rabbitmq"
	"github.com/palhaziadev/forex-dashboard/server/internal/controller"
	"github.com/streadway/amqp"
	"golang.org/x/net/websocket"
)

// TODO dependency injection?
type candleStickData struct {
	Open  float64 `json:"open"`
	Close float64 `json:"close"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	X     int64   `json:"x"` //timestamp
}

// var messageChan chan *testData = make(chan *testData)
// var messageChan chan *[]models.CurrencyData = make(chan *[]models.CurrencyData)
var messageChan chan *[]candleStickData = make(chan *[]candleStickData)

type Server struct {
	controller *controller.Controller
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{
		controller: controller.NewController(), // TODO maybe add controller as a param from main.go
		httpServer: &http.Server{
			Addr: ":8090",
		},
	}
}

func (server *Server) consumerHandler(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		fmt.Printf("Error occurred in RMQ consumer: %s", err)
	}

	fmt.Printf("Message received on '%s' queue: %s", queue, msg.Body)
	data := []models.CurrencyData{}
	err = json.Unmarshal(msg.Body, &data)
	if err != nil {
		fmt.Printf("Error occurred in consumer Unmarshal: %s", err)
	}

	// TODO put calculation to its own place
	candleStick := []candleStickData{}
	for _, v := range data {
		newItem := candleStickData{Open: v.Open, Close: v.Close, High: v.High, Low: v.Low, X: time.Now().UnixMilli()}
		candleStick = append(candleStick, newItem)
		// fmt.Printf("%s -> %s\n", k, v)
	}
	////

	fmt.Printf("----------\n\n: %v------------\n\n", data)
	// go func() { messageChan <- &data }()
	select {
	case messageChan <- &candleStick:
		//if can't write to the channel drop the message
		fmt.Printf("----------\n\n: dropper msg%v------------\n\n", data)
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
			fmt.Printf("Message received ws: '%s'\n\n%v\n\n\n", b, msg)
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

func (server *Server) KillApp(w http.ResponseWriter, r *http.Request) {
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}

func (server *Server) RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		http.Handle("/websocket", websocket.Handler(server.testWebsocketHandler))
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, server.controller.Path), http.HandlerFunc(server.controller.HandleForexList))
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "kill"), http.HandlerFunc(server.KillApp))
	}
	return returnFunction
}

func (server *Server) Start(apiBasePath string) {
	// rabbitmq shutdown?
	go server.MqConsumerTest()
	server.RegisterHandlers(apiBasePath)()

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %s\n", err)
		}
	}()
}

func (server *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		fmt.Println("++++++++++++extra handling here+++++++++++")
		cancel()
	}()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
