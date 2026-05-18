package main

import (
	"log/slog"
	"os"

	"github.com/Hacklabs-app/merch-backend/internal/handler"
	"github.com/Hacklabs-app/merch-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	app := fiber.New(fiber.Config{
		AppName:               "merch API",
		DisableStartupMessage: true,
	})

	app.Use(middleware.RequestLogger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the merch API",
			"docs":    "Please use /api/v1 endpoints",
		})
	})

	v1 := app.Group("/api/v1")
	v1.Get("/health", handler.HealthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("merch API starting", slog.String("port", port))

	err := app.Listen(":" + port)
	if err != nil {
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
