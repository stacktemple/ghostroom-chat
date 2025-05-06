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
		return fiber.NewError(fiber.StatusInternalServerError, "Database not responding")
	}

	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": h.AppName,
	})
}
