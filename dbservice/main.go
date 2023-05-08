package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	server "github.com/palhaziadev/forex-dashboard/dbservice/internal"
	"github.com/palhaziadev/forex-dashboard/dbservice/internal/controller"
)

var basePath = "/api" // TODO config

func dbConnect() *gorm.DB {
	dsn := "host=postgres user=adam password=password dbname=forex_dashboard port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func main() {
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	db := dbConnect()

	ctrl := controller.NewController(db)
	server := server.NewServer(ctrl)

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
