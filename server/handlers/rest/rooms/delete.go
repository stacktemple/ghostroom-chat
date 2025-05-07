package rooms

import (
	"github.com/gofiber/fiber/v2"
)

func (h *RoomHandler) DeleteRoom(c *fiber.Ctx) error {
	roomNameParam := c.Params("name") // from URL
	roomNameToken := c.Locals(ctxRoomName).(string)
	guestName := c.Locals(ctxGuestName).(string)

	if roomNameParam != roomNameToken {
		return fiber.NewError(fiber.StatusUnauthorized, "Token room mismatch")
	}

	room, err := h.repo.GetRoomByNameToday(roomNameToken)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Room not found")
	}

	isOwner, err := h.repo.IsGuestOwner(room.ID, guestName)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check owner")
	}
	if !isOwner {
		return fiber.NewError(fiber.StatusForbidden, "Only owner can delete room")
	}

	if err := h.repo.DeleteRoomByID(room.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete room")
	}

	return c.JSON(fiber.Map{
		"message": "Room deleted successfully",
	})
}
