package main

import (
	"fmt"
	"net/http"
	"time"

	server "github.com/palhaziadev/forex-dashboard/server/internal"
)

var basePath = "/api" // TODO config

// TODO try gin

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(20 * time.Second)

	srv := server.NewServer()

	// TODO move to NewServer?
	srv.RegisterHandlers(basePath)()
	go srv.MqConsumerTest()

	http.ListenAndServe(":8090", nil)
}
