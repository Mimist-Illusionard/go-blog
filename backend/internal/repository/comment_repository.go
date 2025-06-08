package repository

import (
	"go-blog/backend/internal/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(post *models.Comment) error
	Edit(id uint, post *models.Comment) error
	Delete(id uint) error

	GetByID(id uint) (*models.Comment, error)
	GetAllComments(post *models.Post) (*[]models.Comment, error)
}

type CommentGORMRepository struct {
	DB *gorm.DB
}

func (r *CommentGORMRepository) Create(comment *models.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentGORMRepository) Edit(id uint, post *models.Comment) error {
	return r.DB.Model(&models.Comment{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Text": post.Text,
	}).Error
}

func (r *CommentGORMRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Comment{}, id).Error
}

func (r *CommentGORMRepository) GetByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB.First(&comment, id).Error

	return &comment, err
}

func (r *CommentGORMRepository) GetAllComments(post *models.Post) (*[]models.Comment, error) {
	var comments []models.Comment
	err := r.DB.Where("post_id = ?", post.ID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
