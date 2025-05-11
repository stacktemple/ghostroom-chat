package messages

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *MessageHandler) ListMessages(c *fiber.Ctx) error {
	roomName := c.Params("roomName")
	if roomName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing room name")
	}

	limitStr := c.Query("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50
	}

	messages, err := h.Repo.GetMessages(roomName, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "DB error: "+err.Error())
	}

	return c.JSON(messages)
}
