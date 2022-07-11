package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/palhaziadev/forex-dashboard/mockservice/internal/generator"
	config "github.com/palhaziadev/forex-dashboard/pkg/config"
	mqUtils "github.com/palhaziadev/forex-dashboard/pkg/rabbitmq"
)

// constructor for controller?
// https://app.pluralsight.com/course-player?clipId=b19bc9d6-4dac-449e-8ad6-a8e164cb5c75
// add/remove and slice loop
// https://app.pluralsight.com/course-player?clipId=4993f957-89c2-49a3-a4a3-59ce1fb1b23f
// controller functions json encoder
// https://app.pluralsight.com/course-player?clipId=8c3815e3-c168-4e4a-8441-838ef6691ee1

var healthCheckPath = "health"
var dataPath = "data"

type dataResponse struct {
	Data int `json:"data"`
}

type Message struct {
	Message interface{}
}

type healthStatusResponse struct {
	Status string `json:"status"`
}

const (
	healthSatusOk string = "healthy1"
)

type MockService struct {
	generator *generator.MockGenerator
}

func NewMockService(generator *generator.MockGenerator) *MockService {
	return &MockService{
		generator: generator,
	}
}

func (service *MockService) handleData(w http.ResponseWriter, r *http.Request) {
	var data = dataResponse{Data: rand.Intn(100)}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(data)
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

func (service *MockService) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	var healthCheck = healthStatusResponse{Status: healthSatusOk}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(string(healthCheck.Status) + " mockservice")
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

func (service *MockService) sendRabbitMessage(msg Message) {
	connectionString := config.GetEnvVar("RMQ_URL")

	rmqProducer := mqUtils.RMQProducer{
		Queue:            config.GetEnvVar("TEST_QUEUE"),
		ConnectionString: connectionString,
	}

	marshalledMsg, err := json.Marshal(msg.Message)
	if err != nil {
		log.Fatal(err)
	}
	rmqProducer.PublishMessage("text/plain", []byte(marshalledMsg))
}

func (service *MockService) SendData() {
	for {
		//send pock data every 2 seconds
		service.sendRabbitMessage(Message{Message: service.generator.GenerateNumber()})
		time.Sleep(2 * time.Second)
	}
}

func (service *MockService) RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		healthCheckHandler := http.HandlerFunc(service.handleHealthCheck)
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, healthCheckPath), healthCheckHandler)
		dataHandler := http.HandlerFunc(service.handleData)
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, dataPath), dataHandler)
	}
	return returnFunction
}
