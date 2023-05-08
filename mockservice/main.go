package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/palhaziadev/forex-dashboard/mockservice/internal"
	"github.com/palhaziadev/forex-dashboard/mockservice/internal/controller"
	"github.com/palhaziadev/forex-dashboard/mockservice/internal/generator"
	"github.com/palhaziadev/forex-dashboard/mockservice/internal/service"
)

var basePath = "/api" // TODO config

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	fmt.Printf("Current Unix Time: 1")
	// time.Sleep(125 * time.Second)

	svc := service.NewMockService(generator.NewMockGenerator())
	controller := controller.NewController(svc)
	server := server.NewServer(controller)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server.Start(basePath)
	log.Println("Server Started")
	<-done
	log.Println("Server Stopped")
	server.Shutdown()
	log.Println("Server Exited Properly")
	defer os.Exit(0)
}
