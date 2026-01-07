package main

import (
	"github.com/ChanatpakornS/Fytini/fyt/internal/domain/url"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func main() {
	app := fiber.New()

	// middleware
	app.Use(cors.New())
	app.All("/healthz", healthcheck.New())

	// Initialize handler
	urlHandler := url.NewHandler()

	// Mount endpoint
	api := app.Group("/api")
	v1 := api.Group("/v1")
	urlHandler.Mount(v1)

	app.Listen(":3000")
}
