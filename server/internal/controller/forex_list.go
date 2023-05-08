package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
	Path string
}

func NewController() *Controller {
	return &Controller{
		Path: "forexList",
	}
}

// TODO add new endpoint and db connect in server or main??

// TODO put this to pkg
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (controller *Controller) HandleForexList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	// var healthCheck = healthStatusResponse{Status: healthSatusOk}
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal("{EURUSD: true}")
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
