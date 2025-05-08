package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/middleware"
)

type RoomHandler struct {
	DB        *sqlx.DB
	JWTSecret string
	Repo      *RoomRepo
}

const (
	ctxRoomName   = "room_name"
	ctxGuestName  = "guest_name"
	ctxIssuedDate = "issued_date"
)

func RegisterRoutes(r fiber.Router, h *RoomHandler) {
	if h.Repo == nil {
		h.Repo = &RoomRepo{DB: h.DB}
	}
	r.Get("/today", h.ListTodayRooms)
	r.Post("/", h.CreateRoom)
	r.Post("/join", h.JoinRoom)
	r.Get("/verify-token", middleware.JWTGuard(h.JWTSecret), h.CheckToken)
	r.Delete("/:name", middleware.JWTGuard(h.JWTSecret), h.DeleteRoom)
}
