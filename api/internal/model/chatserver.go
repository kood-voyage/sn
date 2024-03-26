package model

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	id   string
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		id:   id,
	}
}

type ChatService struct {
	sync.Mutex
	Clients []*Client
}

func NewChatServer() *ChatService {
	return &ChatService{}
}

type RequestBody struct {
	Type     string `json:"type"`
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
	Message  string `json:"message"`
}

// need middleware to grab user_id
func (cs *ChatService) HandleWS(w http.ResponseWriter, r *http.Request) {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := "testid"
	//append clients to keep track of clients
	client := NewClient(id, conn)
	cs.AddClient(client)
	fmt.Printf("Clients connected: %+v\n", cs.Clients)

	go client.wsRecieveLoop(cs)
}

func (cs *ChatService) RemoveClient(client *Client) error {
	cs.Lock()
	defer cs.Unlock()
	for i, c := range cs.Clients {
		if c == client {
			cs.Clients = append(cs.Clients[:i], cs.Clients[i+1:]...)
			return nil
		}
		if i == len(cs.Clients) {
			return errors.New("client not found")
		}
	}
	return nil
}

func (cs *ChatService) AddClient(client *Client) {
	cs.Lock()
	defer cs.Unlock()
	cs.Clients = append(cs.Clients, client)
}

func (c *Client) wsRecieveLoop(cs *ChatService) {
	fmt.Printf("New client %+v connected...\n", c)
	defer func() {
		fmt.Printf("Client %+v is leaving...\n", c)
		cs.RemoveClient(c)
	}()
	for {
		var body RequestBody
		if err := c.conn.ReadJSON(&body); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}
			log.Println("Read error:", err)
			continue
		}
		fmt.Printf("Recieved message: %+v\n", body)
		c.conn.WriteJSON(body)
	}
}
