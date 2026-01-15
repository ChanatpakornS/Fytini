package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Fytini/pirareus/config"
	"Fytini/pirareus/internal/features/fyt"
	"Fytini/pirareus/internal/features/tini"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	// Health check
	app.All("/healthz", healthcheck.New())

	// API versioning
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Initialize service gateways
	fytServer := fyt.NewServer(cfg)
	tiniServer := tini.NewServer(cfg)

	// Mount service routes
	fytServer.Mount(v1)
	tiniServer.Mount(v1)

	// Root endpoint
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": cfg.App.Name,
			"version": "1.0.0",
			"status":  "running",
		})
	})

	// Start server
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
			stop()
		}
	}()

	fmt.Printf("%s gateway started on port %d\n", cfg.App.Name, cfg.App.Port)
	fmt.Printf("Proxying to:\n")
	fmt.Printf("  - Fyt service: %s\n", cfg.Services.Fyt.URL)
	fmt.Printf("  - Tini service: %s\n", cfg.Services.Tini.URL)

	// Graceful shutdown
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			fmt.Printf("Failed to shutdown server: %v\n", err)
		}
		fmt.Println("Server gracefully stopped")
	}()

	<-ctx.Done()
}
