package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/palhaziadev/forex-dashboard/mockservice/internal/generator"
	"github.com/palhaziadev/forex-dashboard/mockservice/internal/service"
)

var basePath = "/api" // TODO config

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(25 * time.Second)

	svc := service.NewMockService(generator.NewMockGenerator())

	// TODO move to NewMockService??
	svc.RegisterHandlers(basePath)()
	go svc.SendData()

	http.ListenAndServe(":8091", nil)
}
