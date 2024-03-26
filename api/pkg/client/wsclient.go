package client

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	Client *url.URL
}

type ChatClient interface {
	//all the endpoints
	Connect(w http.ResponseWriter, r *http.Request) error
}

func NewClient(wsaddr string) *Client {
	return &Client{
		Client: parseURL(wsaddr),
	}
}

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

func (h *Client) Connect(w http.ResponseWriter, r *http.Request) error {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Connect to the target WebSocket server
	targetConn, _, err := websocket.DefaultDialer.Dial(h.Client.String(), nil)
	if err != nil {
		return err
	}
	defer targetConn.Close()

	// Proxy messages between the client and the target server
	go proxyMessages(conn, targetConn)
	proxyMessages(targetConn, conn)
	return nil
}
func proxyMessages(src, dst *websocket.Conn) {
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			src.Close()
			dst.Close()
			return
		}
		err = dst.WriteMessage(messageType, message)
		if err != nil {
			src.Close()
			dst.Close()
			return
		}
	}
}
