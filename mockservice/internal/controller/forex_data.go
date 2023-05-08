package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/palhaziadev/forex-dashboard/mockservice/internal/service"
	httpUtils "github.com/palhaziadev/forex-dashboard/pkg/httpUtils"
)

type Controller struct {
	service *service.MockService
}

func NewController(service *service.MockService) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) HandleData(w http.ResponseWriter, r *http.Request) {
	httpUtils.SetupResponse(&w, r)
	var data = controller.service.GenerateInitialData("EURUSD", 5) // TODO move this out
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
