package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RoomHandler struct {
	DB        *sqlx.DB
	JWTSecret string
	repo      *RoomRepo
}

func RegisterRoutes(r fiber.Router, h *RoomHandler) {
	if h.repo == nil {
		h.repo = &RoomRepo{DB: h.DB}
	}
	r.Get("/today", h.ListTodayRooms)
	r.Post("/", h.CreateRoom)
	r.Post("/join", h.JoinRoom)
	r.Get("/verify-token", h.CheckToken)
}
