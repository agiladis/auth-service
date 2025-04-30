package controller

import (
	"github.com/agiladis/auth-service/internal/model"
	"github.com/agiladis/auth-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func (h *AuthController) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	return h.service.Register(c, req)
}

func (h *AuthController) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	return h.service.Login(c, req)
}
