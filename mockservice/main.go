package main

import (
	"fmt"
	"net/http"
	"time"

	service "github.com/palhaziadev/forex-dashboard/mockservice/internal"
)

var basePath = "/api" // TODO config

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(25 * time.Second)

	service.RegisterHandlers(basePath)()
	go service.SendData()

	http.ListenAndServe(":8091", nil)
}
