package repository

import "go-blog/internal/models"

type CommentRepository interface {
	Create(post *models.Comment) error
	Edit(id uint, post *models.Comment) error
	Delete(post *models.Comment) error

	AddComment(post *models.Post, comment *models.Comment) error

	GetByID(id uint) (*models.Comment, error)
	GetAllComments(post *models.Post) (*[]models.Comment, error)
}
