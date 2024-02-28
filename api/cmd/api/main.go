package main

import (
	"log"
	"social-network/api/internal/app/server"
)

func main() {
	config := server.NewConfig()
	err := config.ReadConfig("api/config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	log.Fatal(server.Start(config))
}
