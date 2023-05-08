package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/palhaziadev/forex-dashboard/mockservice/internal/generator"
	config "github.com/palhaziadev/forex-dashboard/pkg/config"
	"github.com/palhaziadev/forex-dashboard/pkg/models"
	mqUtils "github.com/palhaziadev/forex-dashboard/pkg/rabbitmq"
)

// constructor for controller?
// https://app.pluralsight.com/course-player?clipId=b19bc9d6-4dac-449e-8ad6-a8e164cb5c75
// add/remove and slice loop
// https://app.pluralsight.com/course-player?clipId=4993f957-89c2-49a3-a4a3-59ce1fb1b23f
// controller functions json encoder
// https://app.pluralsight.com/course-player?clipId=8c3815e3-c168-4e4a-8441-838ef6691ee1

var dataPath = "data"

// type dataResponse struct {
// 	Data int `json:"data"`
// }

type Message struct {
	Message interface{}
}
type MockService struct {
	generator *generator.MockGenerator
}

func NewMockService(generator *generator.MockGenerator) *MockService {
	return &MockService{
		generator: generator,
	}
}

func (service *MockService) sendRabbitMessage(msg Message) error {
	connectionString := config.GetEnvVar("RMQ_URL")

	rmqProducer := mqUtils.RMQProducer{
		Queue:            config.GetEnvVar("TEST_QUEUE"),
		ConnectionString: connectionString,
	}

	marshalledMsg, err := json.Marshal(msg.Message)
	if err != nil {
		log.Println(err)
	}
	err = rmqProducer.PublishMessage("text/plain", []byte(marshalledMsg))
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (service *MockService) SendData(isRabbitMQ bool) {
	initialData := service.GenerateInitialData("EURUSD", 30) // TODO move this out
	lastElement := initialData[len(initialData)-1]
	// firstElement := []models.CurrencyData{{CurrencyPair: lastElement.CurrencyPair, Value: 2}}
	// service.sendRabbitMessage(Message{Message: firstElement})
	if isRabbitMQ { // TODO needed?
		err := service.sendRabbitMessage(Message{Message: initialData})
		if err != nil {
			log.Println(err)
			return
		}
		for {
			// TODO fix CurrencyPair, and High, Low, Close
			nextElement := []models.CurrencyData{{CurrencyPair: lastElement.CurrencyPair, Open: service.GenerateNextValue(lastElement)}}
			// []models.CurrencyData
			//send mock data every 2 seconds
			service.sendRabbitMessage(Message{Message: nextElement})
			time.Sleep(2 * time.Second)
			lastElement = nextElement[0]
		}
	} else {
		log.Println("---------------------------------------")
		for i := 0; i < len(initialData); i++ {
			fmt.Println(initialData[i])
		}
		for {
			// TODO fix CurrencyPair, and High, Low, Close
			nextElement := []models.CurrencyData{{CurrencyPair: lastElement.CurrencyPair, Open: service.GenerateNextValue(lastElement)}}
			// []models.CurrencyData
			//send mock data every 2 seconds
			service.sendRabbitMessage(Message{Message: nextElement})
			time.Sleep(2 * time.Second)
			lastElement = nextElement[0]
		}
		// log.Println("---------------------------------------")
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (service *MockService) JsonData() {
	rand.Seed(time.Now().UnixNano())
	tmp := make([]models.CurrencyData, 20)
	for i := range tmp {
		tmp[i].CurrencyPair = "EURUSD"
		tmp[i].Meta.BaseCurrency = "EUR"
		tmp[i].Meta.QuoteCurrency = "USD"
		if i == 0 {
			log.Println(1 + roundFloat(rand.Float64()*(0.1), 5))
			// High, Low, Close?
			tmp[i].Open = 1 + roundFloat(rand.Float64()*(0.1), 5)
		} else {
			newValue := roundFloat(rand.Float64()*(0.001-0), 5)
			isMinus := rand.Intn(2) == 1
			log.Println(isMinus)
			if isMinus {
				newValue = math.Copysign(newValue, -1)
				log.Println(newValue)
			}
			tmp[i].Open = roundFloat(tmp[i-1].Open+newValue, 5)
		}
	}
	fmt.Println(tmp)

	content, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("mockData.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (service *MockService) GenerateInitialData(currency string, size int) []models.CurrencyData {
	rand.Seed(time.Now().UnixNano())
	tmp := make([]models.CurrencyData, size)
	for i := range tmp {
		tmp[i].CurrencyPair = currency
		if i == 0 {
			log.Println(1 + roundFloat(rand.Float64()*(0.1), 5))
			tmp[i].Open = 1 + roundFloat(rand.Float64()*(0.1), 5)
		} else {
			tmp[i].Open = service.GenerateNextValue(tmp[i-1])
		}
	}
	// fmt.Println(tmp)
	return tmp
}

func (service *MockService) GenerateNextValue(prevValue models.CurrencyData) float64 {
	newValue := roundFloat(rand.Float64()*(0.001-0), 5)
	isMinus := rand.Intn(2) == 1
	log.Println(isMinus)
	if isMinus {
		newValue = math.Copysign(newValue, -1)
		// log.Println(newValue)
	}
	return roundFloat(prevValue.Open+newValue, 5)
}
