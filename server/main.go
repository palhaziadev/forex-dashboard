package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/palhaziadev/forex-dashboard/server/internal"
)

var basePath = "/api" // TODO config

// TODO try gin

func main() {
	fmt.Printf("Current Unix Time2: %v\n", time.Now().Unix())
	// time.Sleep(20 * time.Second)

	srv := server.NewServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv.Start(basePath)
	log.Println("Server Started2")
	<-done
	log.Println("Server Stopped")
	srv.Shutdown()
	log.Println("Server Exited Properly")
	defer os.Exit(0)
}
