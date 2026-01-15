package tini

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
		client: client.NewClient(cfg.Services.Tini.URL),
		config: cfg,
	}
}

func (s *Server) Mount(r fiber.Router) {
	tiniGroup := r.Group("/tini")
	tiniGroup.Post("/url/redirect", s.RedirectURL)
}

type GetShortenUrlRequest struct {
	CustomAlias string `json:"custom_alias"`
}

type GetShortenUrlResponse struct {
	Url string `json:"url"`
}

func (s *Server) RedirectURL(c fiber.Ctx) error {
	// Parse request body
	var req GetShortenUrlRequest
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// Validate required fields
	if req.CustomAlias == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "custom_alias is required",
		})
	}

	// Proxy request to Tini service
	body, statusCode, err := s.client.Post("/url/redirect", req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to communicate with redirect service",
		})
	}

	// Handle not found
	if statusCode == fiber.StatusNotFound {
		// Try to parse error response
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			return c.Status(statusCode).JSON(errorResponse)
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	// Handle success response
	if statusCode == fiber.StatusOK {
		var urlResponse GetShortenUrlResponse
		if err := json.Unmarshal(body, &urlResponse); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to parse response from redirect service",
			})
		}

		// Return the URL
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"url": urlResponse.Url,
		})
	}

	// For other error responses, try to parse and forward the error message
	if len(body) > 0 {
		var errorResponse map[string]interface{}
		if err := json.Unmarshal(body, &errorResponse); err == nil {
			return c.Status(statusCode).JSON(errorResponse)
		}
	}

	// Fallback: just forward the status code
	return c.SendStatus(statusCode)
}
