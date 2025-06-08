package repository

import (
	"go-blog/backend/internal/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	Edit(id uint, post *models.Post) error
	Delete(id uint) error

	GetPostByID(id uint) (*models.Post, error)
	GetPostWithComments(postId uint) (*models.Post, error)
	GetAllPosts() (*[]models.Post, error)
}

type PostGORMRepository struct {
	DB *gorm.DB
}

func (r *PostGORMRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostGORMRepository) Edit(id uint, post *models.Post) error {
	return r.DB.Model(&models.Post{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Text": post.Text,
	}).Error
}

func (r *PostGORMRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Post{}, id).Error
}

func (r *PostGORMRepository) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.DB.First(&post, id).Error

	return &post, err
}

func (r *PostGORMRepository) GetPostWithComments(postId uint) (*models.Post, error) {
	var post models.Post
	err := r.DB.
		Preload("Comments").
		Preload("Comments.User").
		First(&post, postId).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostGORMRepository) GetAllPosts() (*[]models.Post, error) {
	var posts []models.Post

	result := r.DB.
		Preload("User").
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	return &posts, nil
}
