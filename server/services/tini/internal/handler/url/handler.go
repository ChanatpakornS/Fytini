package url

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validate *validator.Validate
	db       *gorm.DB
}

func NewHandler(validate *validator.Validate, db *gorm.DB) *Handler {
	return &Handler{
		validate: validate,
		db:       db,
	}
}

func (h *Handler) Mount(r fiber.Router) {
	urlPrefix := r.Group("/url")
	urlPrefix.Get("/redirect", h.GetShortenURL)
}
