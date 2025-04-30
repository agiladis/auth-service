package service

import (
	"github.com/agiladis/auth-service/internal/model"
	"github.com/agiladis/auth-service/internal/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(r *repository.AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Register(c *fiber.Ctx, req model.RegisterRequest) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repo.Create(&user)
}

func (s *AuthService) Login(c *fiber.Ctx, req model.LoginRequest) error {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{"token": "mocked-jwt-token"})
}
