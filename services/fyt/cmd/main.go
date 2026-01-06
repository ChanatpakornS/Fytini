package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func main() {
	app := fiber.New()

	// middleware
	app.Use(cors.New())
	app.All("/healthz", healthcheck.New())

	app.Listen(":3000")
}
