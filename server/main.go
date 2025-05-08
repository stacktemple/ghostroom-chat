package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stacktemple/realtime-chat/server/config"
	"github.com/stacktemple/realtime-chat/server/handlers"
)

func main() {

	config.Init()
	db := config.ConnectDB()

	println(db)

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

	handlers.RegisterRoutes(app, handlers.Dependencies{
		AppName: "StackTemple", DB: db, JWTSecret: config.Cfg.JWTSecret,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, StackTemple!")
	})

	println(config.Cfg.DatabaseURL)
	println(config.Cfg.JWTSecret)

	log.Fatal(app.Listen(":" + config.Cfg.Port))
}
