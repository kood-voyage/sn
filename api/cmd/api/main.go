package main

import (
	"log"
	"social-network/internal/app/config"
	"social-network/internal/app/server"
)

// @title Social Network API
// @version 0.1
// @description api server for social network

// @host localhost:8080
// @BasePath /

func main() {
	cfg := config.NewConfig()
	err := cfg.ReadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error reading cfg file: %s\n", err)
	}

	log.Fatal(server.Start(cfg))
}
