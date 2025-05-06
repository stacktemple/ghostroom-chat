package rooms

import (
	"github.com/gofiber/fiber/v2"
)

func (h *RoomHandler) ListTodayRooms(c *fiber.Ctx) error {
	rooms, err := h.repo.GetTodayRooms()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "DB error: "+err.Error())
	}
	return c.JSON(rooms)
}
