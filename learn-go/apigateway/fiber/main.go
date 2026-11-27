package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/lathief/test-go/apigateway/internal/gateway"
)

func main() {
	app := fiber.New()
	app.Get("/api/products", adaptor.HTTPHandler(gateway.Products()))
	log.Fatal(app.Listen("127.0.0.1:8081"))
}
