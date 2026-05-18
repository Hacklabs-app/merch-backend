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

		// If there's a Fiber error (like 404 Route Not Found), handle its status code correctly
		status := c.Response().StatusCode()
		if e, ok := err.(*fiber.Error); ok {
			status = e.Code
		}

		if status >= 500 {
			msg := "Server Error"
			if err != nil {
				msg = err.Error()
			}
			
			reqLogger.Error(msg, 
				slog.Int("status", status),
			)
			return err
		}

		// 404s and 400s are client errors, they shouldn't trigger PagerDuty alerts, so we log them as Warnings or Info
		if status >= 400 {
			reqLogger.Warn("Client Error", 
				slog.Int("status", status),
				slog.String("error", err.Error()),
			)
			return err
		}

		reqLogger.Info("Request completed", 
			slog.Int("status", status),
		)

		return nil
	}
}
