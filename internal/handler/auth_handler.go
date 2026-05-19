package handler

import (
	"log/slog"
	"strings"

	"github.com/Hacklabs-app/merch-backend/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(service domain.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.UserRegistration
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Email == "" || req.Password == "" || req.FullName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	logger := c.Locals("logger").(*slog.Logger)

	user, err := h.authService.Register(c.Context(), req.Email, req.Password, req.FullName, req.PhoneNumber)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
		}
		logger.Error("Registration failed", slog.String("error", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.UserRegistration
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	logger := c.Locals("logger").(*slog.Logger)

	token, err := h.authService.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		logger.Warn("Login attempt failed", slog.String("error", err.Error()))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	return c.JSON(fiber.Map{"token": token})
}
