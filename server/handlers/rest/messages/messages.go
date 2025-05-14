package messages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/middleware"
	"github.com/stacktemple/realtime-chat/server/repository"
)

const (
	ctxRoomName   = "room_name"
	ctxGuestName  = "guest_name"
	ctxIssuedDate = "issued_date"
)

type MessageHandler struct {
	DB        *sqlx.DB
	JWTSecret string
	Repo      *repository.MessageRepository
}

func RegisterRoutes(r fiber.Router, h *MessageHandler) {
	r.Get("/", middleware.JWTGuard(h.JWTSecret), h.ListMessages)
}
