package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/palhaziadev/forex-dashboard/mockservice/internal/controller"
)

type Server struct {
	controller *controller.Controller
	httpServer *http.Server
}

func NewServer(controller *controller.Controller) *Server {
	return &Server{
		controller: controller,
		httpServer: &http.Server{
			Addr: ":8091",
		},
	}
}

func (server *Server) RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		dataPath := "forexData"
		dataHandler := http.HandlerFunc(server.controller.HandleData)
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, dataPath), dataHandler)
	}
	return returnFunction
}

func (server *Server) Start(apiBasePath string) {
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
		// rabbitmq shutdown?
		fmt.Println("++++++++++++extra handling here+++++++++++")
		cancel()
	}()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
