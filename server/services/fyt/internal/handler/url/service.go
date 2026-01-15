package url

import (
	"Fytini/db"

	"Fytini/fyt/internal/dto"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (h *Handler) CreateNewShortURL(c fiber.Ctx) error {
	// pre-function
	req := dto.CreateUrlRequest{}
	if err := c.Bind().All(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// repo
	if err := h.db.Transaction(func(tx *gorm.DB) error {
		newShortenURL := db.Url{
			Url:         req.Url,
			CustomAlias: req.CustomAlias,
		}
		if req.ExpirationDate != "" {
			newShortenURL.ExpirationDate = req.ExpirationDate
		}
		if err := tx.Create(&newShortenURL).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)

}
