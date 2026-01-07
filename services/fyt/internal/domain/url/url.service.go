package url

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateNewShortURL(c fiber.Ctx) {
	fmt.Println("Creating new short URL")
}
