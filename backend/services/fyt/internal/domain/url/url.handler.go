package url

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Mount(r fiber.Router) {
	urlPrefix := r.Group("/url")
	urlPrefix.Post("/shorten", h.CreateNewShortURL)
}
