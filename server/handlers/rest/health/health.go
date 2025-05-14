package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type HealthHandler struct {
	AppName string
	DB      *sqlx.DB
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	var ok bool
	err := h.DB.Get(&ok, "SELECT true")
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status":  "unhealthy",
			"service": h.AppName,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "healthy",
		"service": h.AppName,
	})
}
