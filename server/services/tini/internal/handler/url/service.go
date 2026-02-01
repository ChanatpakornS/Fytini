package url

import (
	"Fytini/db"
	"time"

	"Fytini/tini/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetShortenURL(c fiber.Ctx) error {
	// pre-function
	req := dto.GetShortenUrlRequest{}
	if err := c.Bind().JSON(&req); err != nil {
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
	var originURL db.Url
	result := h.db.Where("custom_alias = ?", req.CustomAlias).First(&originURL)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	if originURL.ExpirationDate != nil {
		if time.Now().After(*originURL.ExpirationDate) {
			status := fiber.StatusNotFound
			if originURL.ExpirationDate.Before(time.Now()) {
				status = fiber.StatusGone
			}
			return c.Status(status).JSON(fiber.Map{
				"error": "URL expired",
			})
		}
	}

	// Return the URL in JSON format instead of redirecting
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": originURL.Url,
	})

}
