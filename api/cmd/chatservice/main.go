package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network/internal/model"
)

func main() {
	cs := model.NewChatServer()
	http.HandleFunc("/ws", cs.HandleWS)
	fmt.Println("CHAT SERVICE STARTED...")
	log.Fatal(http.ListenAndServe(":30000", nil))
}
