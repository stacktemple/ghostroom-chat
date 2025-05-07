package rooms

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *RoomHandler) CheckToken(c *fiber.Ctx) error {
	roomName := c.Locals(ctxRoomName).(string)
	guestName := c.Locals(ctxGuestName).(string)
	issuedDate := c.Locals(ctxIssuedDate).(string)

	// Check date is today
	loc, _ := time.LoadLocation("Asia/Bangkok")
	today := time.Now().In(loc).Format("2006-01-02")
	if issuedDate != today {
		return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
	}

	// Check room exists
	room, err := h.Repo.GetRoomByNameToday(roomName)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Room not found")
	}

	// Check guest exists
	exists, err := h.Repo.GuestExistsToday(room.ID, guestName)
	if err != nil || !exists {
		return fiber.NewError(fiber.StatusUnauthorized, "Guest not found in room")
	}

	return c.JSON(fiber.Map{
		"message":     "Token valid",
		"room_name":   roomName,
		"guest_name":  guestName,
		"issued_date": issuedDate,
	})
}
