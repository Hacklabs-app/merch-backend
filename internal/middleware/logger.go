package middleware

import (
	"log/slog"
	"time"

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

		clientIP := c.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.IP()
		}

		reqLogger := slog.With(
			slog.String("request_id", requestID),
			slog.String("method", c.Method()),
			slog.String("path", c.Path()),
			slog.String("ip", clientIP),
		)

		c.Locals("logger", reqLogger)

		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

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
				slog.String("duration", duration.String()),
			)
			return err
		}

		if status >= 400 {
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}
			reqLogger.Warn("Client Error", 
				slog.Int("status", status),
				slog.String("error", errStr),
				slog.String("duration", duration.String()),
			)
			return err
		}

		reqLogger.Info("Request completed", 
			slog.Int("status", status),
			slog.String("duration", duration.String()),
		)

		return nil
	}
}
