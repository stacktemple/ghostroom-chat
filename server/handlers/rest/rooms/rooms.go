package rooms

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/handlers/socket/chat"
	"github.com/stacktemple/realtime-chat/server/middleware"
	"github.com/stacktemple/realtime-chat/server/repository"
)

type RoomHandler struct {
	JWTSecret string
	Repo      *repository.RoomRepository
	MsgRepo   *repository.MessageRepository
	ChatHub   *chat.Hub
}

const (
	ctxRoomName   = "room_name"
	ctxGuestName  = "guest_name"
	ctxIssuedDate = "issued_date"
)

func RegisterRoutes(r fiber.Router, h *RoomHandler) {
	r.Get("/today", h.ListTodayRooms)
	r.Post("/", h.CreateRoom)
	r.Post("/join", h.JoinRoom)
	r.Get("/verify-token", middleware.JWTGuard(h.JWTSecret), h.CheckToken)
	r.Delete("/:name", middleware.JWTGuard(h.JWTSecret), h.DeleteRoom)
}
