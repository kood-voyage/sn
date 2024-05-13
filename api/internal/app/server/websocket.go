package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"social-network/internal/model"
	"social-network/internal/store"
	"social-network/pkg/validator"

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
	store   store.Store
}

func NewChatServer(store store.Store) *ChatService {
	return &ChatService{
		store: store,
	}
}

type Payload struct {
	Type     string `json:"type" validate:"required|contains:message,status,notification"`
	Address  string `json:"address" validate:"required|contains:group,direct,broadcast"`
	ID       string `json:"id" validate:"required"`
	SourceID string `json:"source_id" validate:"required"`
	Data     any    `json:"data"`
}

// need middleware to grab user_id
func (cs *ChatService) HandleWS(w http.ResponseWriter, r *http.Request) {
	sourceID, ok := r.Context().Value(ctxUserID).(string)
	if !ok {
		fmt.Println("unauthorized")
		return
	}

	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// append client to keep track of clients
	client := NewClient(sourceID, conn)
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
	cs.getOnlineUsers(*client)
	cs.sendUserStatus(*client, true)
	cs.Clients = append(cs.Clients, client)
}

func (c *Client) wsRecieveLoop(cs *ChatService) {
	fmt.Printf("New client %+v connected...\n", c)
	defer func() {
		fmt.Printf("Client %+v is leaving...\n", c)
		cs.sendUserStatus(*c, false)
		cs.RemoveClient(c)
	}()
	for {
		var body Payload
		if err := c.conn.ReadJSON(&body); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				return
			}
			log.Println("Read error:", err)
			continue
		}
		fmt.Printf("Recieved message: %+v\n", body)
		if err := validator.Validate(body); err != nil {
			// validate the payload
			c.conn.WriteJSON(fmt.Sprintf("bad payload - %v", err))
		} else {
			users, err := cs.getUsers(body)
			if err != nil {
				fmt.Println("THIS IS AN ERROR IN wsRecieveLoop", err)
			}
			if err := cs.writeToUsers(users, body); err != nil {
				fmt.Println(" THIS IS AN ERROR IN WSRECIEVE LOOP", err)
			}
		}
	}
}

func (cs *ChatService) getUsers(p Payload) ([]*Client, error) {
	switch p.Address {
	case "broadcast":
		return cs.Clients, nil
	case "group":
		members, err := cs.store.Chat().GetUsers(model.Chat{ID: p.ID})
		if err != nil {
			return nil, err
		}
		var client []*Client
		for _, c := range cs.Clients {
			for _, member := range members {
				if c.id == member.ID {
					client = append(client, c)
				}
			}
		}
		return client, nil
	case "direct":
		client, err := cs.getClient(p.ID)
		if err != nil {
			return nil, err
		}
		return []*Client{client}, nil
	default:
		return nil, errors.New("unknown address")
	}
}

func (cs *ChatService) writeToUsers(clients []*Client, p Payload) error {
	for _, c := range clients {
		if err := c.conn.WriteJSON(p); err != nil {
			return err
		}
	}
	return nil
}

func (cs *ChatService) sendUserStatus(client Client, status bool) {
	// get current user chats list
	// check all the user id-s from the chats list if we have a connection with specific id send that user according status
	userSendList, err := cs.store.Chat().GetChatsForUser(client.id)
	if err != nil {
		fmt.Println("SOMETHING WENT WRONG WHILE CALCULATING STATUSES")
		return
	}
	for _, c := range cs.Clients {
		for _, user := range userSendList {
			if c.id == user.ID {
				c.conn.WriteJSON(Payload{
					Type:     "status",
					Address:  "direct",
					ID:       c.id,
					SourceID: client.id,
					Data:     status,
				})
			}
		}
	}
}

func (cs *ChatService) getOnlineUsers(client Client) {
	// get current user chats list
	// check all the use id-s from the follow list if we hahve a connection with specific id send current user all the information about users -- online or offline
	userSendList, err := cs.store.Chat().GetChatsForUser(client.id)
	if err != nil {
		fmt.Println("SOMETHING WENT WRONG WHILE CALCULATING STATUSES")
		return
	}
	var sendList []model.User

	for _, c := range cs.Clients {
		for _, user := range userSendList {
			if c.id == user.ID {
				sendList = append(sendList, *user)
			}
		}
	}
	r := Payload{
		Type:     "status",
		Address:  "direct",
		ID:       client.id,
		SourceID: "server",
		Data:     sendList,
	}
	if err := client.conn.WriteJSON(r); err != nil {
		fmt.Println("getOnlineUsers ERROR ---> ", err)
	}
}

func (cs *ChatService) getClient(client_id string) (*Client, error) {
	for _, client := range cs.Clients {
		if client.id == client_id {
			return client, nil
		}
	}
	return nil, errors.New("client not found")
}
