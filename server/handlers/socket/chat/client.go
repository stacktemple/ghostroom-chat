package chat

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	RoomName string
	Guest    string
	Hub      *Hub
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		}
		if err := c.Conn.ReadJSON(&msg); err != nil {
			log.Println("read error:", err)
			break
		}

		if len(strings.TrimSpace(msg.Content)) == 0 {
			continue
		}

		// Allow only "text" type for now
		if msg.Type != "text" {
			log.Println("unsupported message type:", msg.Type)
			continue
		}

		c.Hub.Broadcast <- BroadcastMessage{
			RoomName:  c.RoomName,
			GuestName: c.Guest,
			Type:      msg.Type,
			Content:   msg.Content,
			Time:      time.Now(),
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
