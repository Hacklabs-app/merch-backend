package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func HealthHandler(c *fiber.Ctx) error {
	logger := c.Locals("logger").(*slog.Logger)
	
	logger.Info("Executing health check")

	return c.JSON(fiber.Map{
		"status":  "healthy",
		"service": "merch API",
		"version": "v1",
	})
}
