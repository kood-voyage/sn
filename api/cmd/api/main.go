package main

import (
	"fmt"
	"log"
	"social-network/internal/app/server"
	"social-network/pkg/validator"
)

func main() {
	config := server.NewConfig()
	err := config.ReadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	test()
	log.Fatal(server.Start(config))
}

func test() {
	type user struct {
		Email    string `json:"email" validate:"required|min_len:5|max_len:25|email"`
		Username string `json:"username" validate:"min_len:5|integer"`
		Age      int    `json:"age"`
		Testfield []string `validate:"min_len:1"`
	}

	user1 := user{
		Email:    "johndoe@gmail.com",
		Username: "22222",
		Age:      2,
		Testfield: []string{"hello", "world"},
	}
	fmt.Println(validator.Validate(user1))
}
