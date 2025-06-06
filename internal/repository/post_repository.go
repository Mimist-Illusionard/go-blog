package repository

import (
	"go-blog/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	Edit(id uint, post *models.Post) error
	Delete(id uint) error

	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() (*[]models.Post, error)
}

type PostGORMRepository struct {
	DB *gorm.DB
}

func (r *PostGORMRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostGORMRepository) Edit(id uint, post *models.Post) error {
	return r.DB.Find(post, id).Updates(post).Error
}

func (r *PostGORMRepository) Delete(id uint) error {

	post, err := r.GetPostByID(id)

	if err != nil {
		return err
	}

	return r.DB.Delete(post).Error
}

func (r *PostGORMRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.DB.First(&post, id).Error

	return &post, err
}

func (r *PostGORMRepository) GetAllPosts() (*[]models.Post, error) {
	var posts []models.Post
	result := r.DB.Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &posts, nil
}
