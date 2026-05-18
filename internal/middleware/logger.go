package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		
		c.Set("X-Request-ID", requestID)

		reqLogger := slog.With(
			slog.String("request_id", requestID),
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.String("ip", c.IP()),
		)

		c.Locals("logger", reqLogger)

		reqLogger.Info("Incoming request")

		err := c.Next()

		if err != nil {
			reqLogger.Error("Request failed", 
				slog.Int("status", c.Response().StatusCode()),
				slog.String("error", err.Error()),
			)
			return err
		}

		reqLogger.Info("Request completed", 
			slog.Int("status", c.Response().StatusCode()),
		)

		return nil
	}
}
