package rooms

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/auth"
	"github.com/stacktemple/realtime-chat/server/handlers/socket/chat"
)

type JoinRoomPayload struct {
	Name      string `json:"name"`
	Password  string `json:"password,omitempty"`
	GuestName string `json:"guest_name"`
}

func (h *RoomHandler) JoinRoom(c *fiber.Ctx) error {
	var payload JoinRoomPayload
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}

	if payload.Name == "" || payload.GuestName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing room name or guest name")
	}

	// get room
	room, err := h.Repo.GetRoomByNameToday(payload.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Database error: "+err.Error())
	}
	if room == nil {
		return fiber.NewError(fiber.StatusNotFound, "Room not found")
	}

	// check password if needed
	if room.NeedPass {

		if room.PasswordHash == nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Missing password hash for protected room")
		}

		match := auth.ComparePassword(*room.PasswordHash, payload.Password)
		if !match {
			return fiber.NewError(fiber.StatusUnauthorized, "Incorrect password")
		}
	}

	// check guest exists
	exists, err := h.Repo.GuestExistsToday(room.ID, payload.GuestName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "DB error: "+err.Error())
	}
	if exists {
		return fiber.NewError(fiber.StatusConflict, "Guest name already used in this room today")
	}

	// add guest
	if err := h.Repo.AddGuest(room.ID, payload.GuestName, false); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Add guest failed: "+err.Error())
	}

	// create token
	loc, _ := time.LoadLocation("Asia/Bangkok")
	issued_date := time.Now().In(loc).Format("2006-01-02")

	claims := map[string]any{
		"room_name":   payload.Name,
		"guest_name":  payload.GuestName,
		"issued_date": issued_date,
	}
	token, err := auth.CreateToken(h.JWTSecret, claims, 24)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Token error: "+err.Error())
	}

	h.ChatHub.Broadcast <- chat.BroadcastMessage{
		RoomName:  payload.Name,
		GuestName: payload.GuestName,
		Type:      "join",
		Content:   "joined the room",
		Time:      time.Now(),
	}

	return c.JSON(fiber.Map{
		"message":     "Joined room",
		"token":       token,
		"issued_date": issued_date,
	})
}
