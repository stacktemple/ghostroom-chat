package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/health"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/messages"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/rooms"
	"github.com/stacktemple/realtime-chat/server/handlers/socket/chat"
	"github.com/stacktemple/realtime-chat/server/repository"
)

type Dependencies struct {
	AppName   string
	DB        *sqlx.DB
	JWTSecret string
	ChatHub   *chat.Hub
}

func RegisterRoutes(app *fiber.App, deps Dependencies) {
	api := app.Group("/api")

	// REST routes

	healthGroup := api.Group("/health")
	healthHandler := health.HealthHandler{AppName: deps.AppName, DB: deps.DB}
	healthGroup.Get("/", healthHandler.Check)

	rooms.RegisterRoutes(api.Group("/rooms"), &rooms.RoomHandler{
		JWTSecret: deps.JWTSecret,
		Repo:      repository.NewRoomRepository(deps.DB),
		MsgRepo:   repository.NewMessageRepository(deps.DB),
	})

	messages.RegisterRoutes(api.Group("/messages"), &messages.MessageHandler{
		JWTSecret: deps.JWTSecret,
		Repo:      repository.NewMessageRepository(deps.DB),
	})

	app.Use("/ws/chat", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/chat/:room", websocket.New(chat.NewHandler(deps.DB, deps.JWTSecret, deps.ChatHub)))
}
