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

	// httpServer := &http.Server{
	// 	Addr:    ":8000",
	// }
	// TODO http server graceful shutdown
	// httpServer.Shutdown(context.Background())

	http.ListenAndServe(":8090", nil)
}
