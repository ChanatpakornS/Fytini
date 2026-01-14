package url

import (
	"Fytini/db"

	"Fytini/tini/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) GetShortenURL(c fiber.Ctx) error {
	// pre-function
	req := dto.GetShortenUrlRequest{}
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
	var originURL db.Url
	result := h.db.Where("custom_alias = ?", req.CustomAlias).First(&originURL)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// fiber.StatusFound = 302 (Analytics friendly)
	// fiber.StatusMovedPermanently = 301 (Cache friendly)
	return c.Redirect().To(originURL.Url)

}
