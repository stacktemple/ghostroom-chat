package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/health"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/rooms"
)

type Dependencies struct {
	AppName   string
	DB        *sqlx.DB
	JWTSecret string
}

func RegisterRoutes(app *fiber.App, deps Dependencies) {
	api := app.Group("/api")

	// REST routes

	healthGroup := api.Group("/health")
	healthHandler := health.HealthHandler{AppName: deps.AppName, DB: deps.DB}
	healthGroup.Get("/", healthHandler.Check)

	rooms.RegisterRoutes(api.Group("/rooms"), &rooms.RoomHandler{
		DB:        deps.DB,
		JWTSecret: deps.JWTSecret,
	})

}
