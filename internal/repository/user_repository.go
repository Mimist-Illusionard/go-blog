package repository

import (
	"go-blog/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByLogin(login string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
}

type UserGORMRepository struct {
	DB *gorm.DB
}

func (r *UserGORMRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserGORMRepository) GetByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("login = ?", login).First(&user).Error
	return &user, err
}

func (r *UserGORMRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.Find(&user, id).Error
	return &user, err
}
