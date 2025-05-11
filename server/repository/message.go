package repository

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	DB *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{DB: db}
}

type Message struct {
	GuestName string    `db:"guest_name" json:"guest_name"`
	Content   string    `db:"content" json:"content"`
	Type      string    `db:"type" json:"type"`
	SentAt    time.Time `db:"sent_at" json:"sent_at"`
}

var validTypes = map[string]bool{
	"create": true,
	"join":   true,
	"text":   true,
}

func (r *MessageRepository) AddMessage(roomName, guestName, content, msgType string) error {

	if !validTypes[msgType] {
		return errors.New("invalid message type")
	}

	_, err := r.DB.Exec(`
		INSERT INTO messages (room_id, guest_name, content, type)
		SELECT id, $1, $2, $3 FROM rooms WHERE name = $4
	`, guestName, content, msgType, roomName)
	return err
}

func (r *MessageRepository) GetMessages(roomName string, limit int) ([]Message, error) {
	var messages []Message
	query := `
		SELECT guest_name, content, type, sent_at
		FROM messages
		WHERE room_id = (SELECT id FROM rooms WHERE name = $1)
		ORDER BY sent_at DESC
		LIMIT $2
	`
	err := r.DB.Select(&messages, query, roomName, limit)
	return messages, err
}
