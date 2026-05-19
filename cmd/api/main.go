package main

import (
	"log/slog"
	"os"

	"github.com/Hacklabs-app/merch-backend/internal/config"
	"github.com/Hacklabs-app/merch-backend/internal/handler"
	"github.com/Hacklabs-app/merch-backend/internal/middleware"
	"github.com/Hacklabs-app/merch-backend/internal/repository/postgres"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db, err := postgres.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Error("Database connection failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()
	logger.Info("Successfully connected to the database")

	if cfg.Environment == "development" {
		err = postgres.RunMigrations(db, "migrations")
		if err != nil {
			logger.Error("Database migration failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
		logger.Info("Database migrations applied successfully")
	}

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

	v1.Get("/crash", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusInternalServerError, "simulated server crash for alerting test")
	})

	app.Use(func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Route not found")
	})

	logger.Info("merch API starting", slog.String("port", cfg.Port))

	err = app.Listen(":" + cfg.Port)
	if err != nil {
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
