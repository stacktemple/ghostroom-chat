package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/config"
)

func main() {

	config.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, StackTemple!")
	})

	println(config.Cfg.DatabaseURL)
	println(config.Cfg.JWTSecret)

	log.Fatal(app.Listen(":" + config.Cfg.Port))
}
