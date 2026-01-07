package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChanatpakornS/Fytini/fyt/config"
	"github.com/ChanatpakornS/Fytini/fyt/internal/domain/url"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
			fmt.Println("failed to start server", err)
			stop()
		}
	}()

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.ShutdownWithContext(shutdownCtx); err != nil {
			fmt.Println("failed to shutdown server", err)
		}
		fmt.Println("gracefully shutdown server")
	}()

	<-ctx.Done()
}
