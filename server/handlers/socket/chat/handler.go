package chat

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/auth"
)

func NewHandler(db *sqlx.DB, jwtSecret string, hub *Hub) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		token := c.Query("token")
		claims, err := auth.ParseToken(jwtSecret, token)
		if err != nil {
			log.Println("invalid token:", err)
			c.Close()
			return
		}

		roomName := claims["room_name"].(string)
		guest := claims["guest_name"].(string)

		client := &Client{
			Conn:     c,
			Send:     make(chan []byte, 256),
			RoomName: roomName,
			Guest:    guest,
			Hub:      hub,
		}

		hub.Register <- client
		go client.WritePump()
		client.ReadPump()
	}
}
