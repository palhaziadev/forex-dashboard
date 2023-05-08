package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/palhaziadev/forex-dashboard/dbservice/internal/controller"
)

type Server struct {
	controller *controller.Controller
	httpServer *http.Server
}

func NewServer(controller *controller.Controller) *Server { // NewServer needs to return pointer? create interface for return type???
	return &Server{
		controller: controller, // TODO maybe add controller as a param from main.go
		httpServer: &http.Server{
			Addr: ":8092",
		},
	}
}

func (server *Server) KillApp(w http.ResponseWriter, r *http.Request) {
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}

func (server *Server) RegisterHandlers(apiBasePath string) func() {
	returnFunction := func() {
		http.Handle(fmt.Sprintf("%s/currency/seedTestData", apiBasePath), http.HandlerFunc(server.controller.SeedTestData))
		http.Handle(fmt.Sprintf("%s/%s", apiBasePath, "kill"), http.HandlerFunc(server.KillApp))
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		// extra handling here
		// db connection close?
		fmt.Println("++++++++++++extra handling here+++++++++++")
		cancel()
	}()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
