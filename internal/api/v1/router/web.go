package router

import (
	"github.com/agiladis/auth-service/internal/api/v1/controller"
	"github.com/agiladis/auth-service/internal/repository"
	"github.com/agiladis/auth-service/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupWebRoutes(router fiber.Router, db *gorm.DB) {
	repo := repository.NewAuthRepository(db)
	service := service.NewAuthService(repo)
	controller := controller.NewAuthController(service)

	web := router.Group("/web")
	web.Post("/register", controller.Register)
	web.Post("/login", controller.Login)
}
