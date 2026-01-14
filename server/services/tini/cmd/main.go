package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Fytini/tini/config"
	"Fytini/tini/internal/handler/url"
	"Fytini/tini/internal/validator"

	"Fytini/pkg/postgres"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Connect to database and migration
	db, err := postgres.Open()
	if err != nil {
		panic(err)
	}
	if err := postgres.Migrate(db); err != nil {
		panic(err)
	}

	validator := validator.New()

	app := fiber.New()

	// middleware
	app.Use(cors.New())
	app.All("/healthz", healthcheck.New())

	// Initialize handler
	urlHandler := url.NewHandler(validator, db)

	// Mount endpoint
	urlHandler.Mount(app)

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
