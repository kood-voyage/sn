package main

import (
	"log"
	"social-network/internal/app/server"
)

// @title Social Network API
// @version 0.1
// @description api server for social network

// @host localhost:8080
// @BasePath /

func main() {
	config := server.NewConfig()
	err := config.ReadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	log.Fatal(server.Start(config))
}
