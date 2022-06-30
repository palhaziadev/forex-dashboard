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

	server.RegisterHandlers(basePath)()
	go server.MqConsumerTest()

	http.ListenAndServe(":8090", nil)
}
