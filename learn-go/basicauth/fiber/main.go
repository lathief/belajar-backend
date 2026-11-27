package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()
	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{"demo": "belajar"},
		Realm: "latihan",
	})
	app.Get("/profil", auth, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"user": c.Locals("username")})
	})
	log.Fatal(app.Listen("127.0.0.1:8081"))
}
