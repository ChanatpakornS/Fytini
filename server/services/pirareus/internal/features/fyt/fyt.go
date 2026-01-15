package fyt

import (
	"Fytini/pirareus/config"
	"Fytini/pirareus/internal/client"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	client *client.Client
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		client: client.NewClient(cfg.Services.Fyt.URL),
		config: cfg,
	}
}

func (s *Server) Mount(r fiber.Router) {
	fytGroup := r.Group("/fyt")
	fytGroup.Post("/url/shorten", s.ShortenURL)
}

type CreateUrlRequest struct {
	Url            string `json:"url"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	CustomAlias    string `json:"custom_alias"`
}

func (s *Server) ShortenURL(c fiber.Ctx) error {
	// Parse request body
	var req CreateUrlRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Validate required fields
	if req.Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "url is required",
		})
	}

	if req.CustomAlias == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "custom_alias is required",
		})
	}

	// Proxy request to Fyt service
	body, statusCode, err := s.client.Post("/url/shorten", req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to communicate with shortening service",
		})
	}

	// If the response is OK, forward it
	if statusCode == fiber.StatusOK {
		return c.SendStatus(fiber.StatusOK)
	}

	// For error responses, try to parse and forward the error message
	if len(body) > 0 {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			return c.Status(statusCode).JSON(errorResponse)
		}
	}

	// Fallback: just forward the status code
	return c.SendStatus(statusCode)
}
