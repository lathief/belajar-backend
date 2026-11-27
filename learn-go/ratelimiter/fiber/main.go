package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	limit := limiter.New(limiter.Config{
		Max:        3,
		Expiration: 10 * time.Second,
		// Semua klien berbagi kuota untuk endpoint ini.
		KeyGenerator: func(c *fiber.Ctx) string { return "ping" },
	})
	app.Get("/ping", limit, func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	log.Fatal(app.Listen("127.0.0.1:8081"))
}
