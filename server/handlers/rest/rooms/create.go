package rooms

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/auth"
)

type CreateRoomPayload struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	GuestName string `json:"guest_name"`
}

const (
	MinNameLength      = 3
	MaxNameLength      = 32
	MinGuestNameLength = 3
	MaxGuestNameLength = 32
	MinPasswordLength  = 4
	MaxPasswordLength  = 12
)

func (h *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	var payload CreateRoomPayload
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}
	if payload.Name == "" || payload.GuestName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Room name must be between "+strconv.Itoa(MinNameLength)+" and "+strconv.Itoa(MaxNameLength)+" characters")
	}

	// check length
	if len(payload.Name) < MinNameLength || len(payload.Name) > MaxGuestNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "Room name must be between "+strconv.Itoa(MinNameLength)+" and "+strconv.Itoa(MaxNameLength)+" characters")
	}
	if len(payload.GuestName) < MinGuestNameLength || len(payload.GuestName) > MaxGuestNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "Guest name must be between "+strconv.Itoa(MinGuestNameLength)+" and "+strconv.Itoa(MaxGuestNameLength)+" characters)")
	}

	if payload.Password != "" && (len(payload.Password) < MinPasswordLength || len(payload.Password) > MaxPasswordLength) {
		return fiber.NewError(fiber.StatusBadRequest, "Password must be 4–12 characters")
	}

	// check room exists
	exists, err := h.repo.RoomExistsToday(payload.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "DB error: "+err.Error())
	}
	if exists {
		return fiber.NewError(fiber.StatusConflict, "Room already exists today")
	}

	// hash password
	var passwordHash *string
	needPass := false
	if payload.Password != "" {
		hash, err := auth.HashPassword(payload.Password)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Hash error: "+err.Error())
		}
		passwordHash = &hash
		needPass = true
	}

	// create room
	roomID, err := h.repo.CreateRoom(payload.Name, passwordHash, needPass)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Create room failed: "+err.Error())
	}
	err = h.repo.AddGuest(roomID, payload.GuestName, true)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Add guest failed: "+err.Error())
	}

	// create token
	claims := map[string]any{
		"room_name":   payload.Name,
		"guest_name":  payload.GuestName,
		"issued_date": time.Now().Format("2006-01-02"),
	}
	token, err := auth.CreateToken(h.JWTSecret, claims, 24)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Token error: "+err.Error())
	}

	// Return token
	return c.JSON(fiber.Map{
		"message": "Room created",
		"token":   token,
	})
}
