package rooms

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/auth"
)

func (h *RoomHandler) CheckToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
	}

	// Remove "Bearer " prefix if exists
	const prefix = "Bearer "
	tokenStr := strings.TrimPrefix(authHeader, prefix)

	claims, err := auth.ParseToken(h.JWTSecret, tokenStr)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	roomName, ok1 := claims["room_name"].(string)
	guestName, ok2 := claims["guest_name"].(string)
	issuedDate, ok3 := claims["issued_date"].(string)

	if !ok1 || !ok2 || !ok3 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
	}

	// Check date is today
	loc, _ := time.LoadLocation("Asia/Bangkok")
	today := time.Now().In(loc).Format("2006-01-02")
	if issuedDate != today {
		return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
	}

	// Check room exists
	room, err := h.repo.GetRoomByNameToday(roomName)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Room not found")
	}

	// Check guest exists
	exists, err := h.repo.GuestExistsToday(room.ID, guestName)
	if err != nil || !exists {
		return fiber.NewError(fiber.StatusUnauthorized, "Guest not found in room")
	}

	return c.JSON(fiber.Map{
		"message":    "Token valid",
		"room_name":  roomName,
		"guest_name": guestName,
	})
}
