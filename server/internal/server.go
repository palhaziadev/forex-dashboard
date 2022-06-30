package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/palhaziadev/forex-dashboard/pkg/config"
	mqUtils "github.com/palhaziadev/forex-dashboard/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"golang.org/x/net/websocket"
)

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

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
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

func consumerHandler(queue string, msg amqp.Delivery, err error) {
	if err != nil {
		fmt.Printf("Error occurred in RMQ consumer: %s", err)
	}

	fmt.Printf("Message received on '%s' queue: %s", queue, string(msg.Body))
	data := testData{}
	err = json.Unmarshal(msg.Body, &data)
	if err != nil {
		fmt.Printf("Error occurred in consumer Unmarshal: %s", err)
	}

	go func() { messageChan <- &data }()
}

func MqConsumerTest() {
	connectionString := config.GetEnvVar("RMQ_URL")

	exampleQueue := mqUtils.RMQConsumer{
		Queue:            config.GetEnvVar("TEST_QUEUE"),
		ConnectionString: connectionString,
		MsgHandler:       consumerHandler,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()
	// Multiple listeners can be specified here

	<-forever
}

func testWebsocketHandler(ws *websocket.Conn) {
	for {
		msg := <-messageChan
		b, err := json.Marshal(msg)
		fmt.Printf("Message received: '%s'\n", b)
		websocket.JSON.Send(ws, msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		healthCheckHandler := http.HandlerFunc(handleHealthCheck)
		http.Handle("/websocket", websocket.Handler(testWebsocketHandler))
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, healthCheckPath), healthCheckHandler)
	}
	return returnFunction
}
