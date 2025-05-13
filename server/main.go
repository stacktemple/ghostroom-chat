package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stacktemple/realtime-chat/server/config"
	"github.com/stacktemple/realtime-chat/server/cronjob"
	"github.com/stacktemple/realtime-chat/server/handlers"
	"github.com/stacktemple/realtime-chat/server/handlers/socket/chat"
	"github.com/stacktemple/realtime-chat/server/repository"
)

func main() {

	config.Init()
	db := config.ConnectDB()

	cronjob.StartCleaner(db)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	app.Use(cors.New())

	messageRepo := repository.NewMessageRepository(db)
	chatHub := chat.NewHub(messageRepo)
	go chatHub.Run()

	handlers.RegisterRoutes(app, handlers.Dependencies{
		AppName: "StackTemple", DB: db, JWTSecret: config.Cfg.JWTSecret, ChatHub: chatHub,
	})

	// println(config.Cfg.DatabaseURL)
	// println(config.Cfg.JWTSecret)

	go func() {
		if err := app.Listen(":" + config.Cfg.Port); err != nil {
			log.Panicf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")

	//Gracefully shutdown fiber server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped cleanly")

}
