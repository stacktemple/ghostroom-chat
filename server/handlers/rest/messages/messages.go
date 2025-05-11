package messages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/middleware"
	"github.com/stacktemple/realtime-chat/server/repository"
)

type MessageHandler struct {
	DB        *sqlx.DB
	JWTSecret string
	Repo      *repository.MessageRepository
}

func RegisterRoutes(r fiber.Router, h *MessageHandler) {
	r.Get("/:roomName", middleware.JWTGuard(h.JWTSecret), h.ListMessages)
}
