package tini

import (
	"Fytini/pirareus/config"
	"Fytini/pirareus/internal/client"

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
	tiniGroup.Get("/url/redirect", s.RedirectURL)
}

func (s *Server) RedirectURL(c fiber.Ctx) error {
	// Get query parameters
	customAlias := c.Query("custom_alias")
	if customAlias == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "custom_alias is required",
		})
	}

	// Prepare query params
	queryParams := map[string]string{
		"custom_alias": customAlias,
	}

	// Proxy redirect request to Tini service
	location, statusCode, err := s.client.ProxyRedirect("/url/redirect", queryParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to communicate with redirect service",
		})
	}

	// Handle different status codes
	if statusCode == fiber.StatusNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	if statusCode >= 300 && statusCode < 400 && location != "" {
		// Proxy the redirect to the client
		return c.Redirect().Status(statusCode).To(location)
	}

	// Forward the response as-is for other status codes
	return c.SendStatus(statusCode)
}
