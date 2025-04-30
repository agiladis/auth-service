package repository

import (
	"github.com/agiladis/auth-service/internal/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	db.AutoMigrate(&model.User{})

	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *AuthRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
