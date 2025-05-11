package chat

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/stacktemple/realtime-chat/server/repository"
)

type BroadcastMessage struct {
	RoomName  string
	GuestName string
	Type      string
	Content   string
	Time      time.Time
}

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan BroadcastMessage
	Messages   *repository.MessageRepository
	Mutex      sync.Mutex
}

func NewHub(messages *repository.MessageRepository) *Hub {
	return &Hub{
		Rooms:      make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan BroadcastMessage),
		Messages:   messages,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			if h.Rooms[client.RoomName] == nil {
				h.Rooms[client.RoomName] = make(map[*Client]bool)
			}
			h.Rooms[client.RoomName][client] = true
			h.Mutex.Unlock()

		case client := <-h.Unregister:
			h.Mutex.Lock()
			if clients, ok := h.Rooms[client.RoomName]; ok {
				delete(clients, client)
				close(client.Send)
				if len(clients) == 0 {
					delete(h.Rooms, client.RoomName)
				}
			}
			h.Mutex.Unlock()

		case msg := <-h.Broadcast:
			// Save to DB
			go func(m BroadcastMessage) {
				err := h.Messages.AddMessage(m.RoomName, m.GuestName, m.Content, m.Type)
				if err != nil {
					log.Println("DB insert error:", err)
				}
			}(msg)

			// Prepare JSON
			data, _ := json.Marshal(map[string]interface{}{
				"guest_name": msg.GuestName,
				"content":    msg.Content,
				"type":       msg.Type,
				"sent_at":    msg.Time.Format(time.RFC3339),
			})

			h.Mutex.Lock()
			for client := range h.Rooms[msg.RoomName] {
				select {
				case client.Send <- data:
				default:
					close(client.Send)
					delete(h.Rooms[msg.RoomName], client)
				}
			}
			h.Mutex.Unlock()
		}
	}
}
