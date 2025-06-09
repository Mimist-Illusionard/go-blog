package services

import (
	"go-blog/internal/models"
	"go-blog/internal/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *CommentService {
	return &CommentService{repo: r}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.repo.Create(comment)
}

func (s *CommentService) EditComment(id uint, comment *models.Comment) error {
	return s.repo.Edit(id, comment)
}

func (s *CommentService) DeleteComment(id uint) error {
	return s.repo.Delete(id)
}

func (s *CommentService) GetCommentByID(id uint) (*models.Comment, error) {
	return s.repo.GetByID(id)
}

func (s *CommentService) GetAllComments(postID uint) (*[]models.Comment, error) {
	post := &models.Post{ID: postID}
	return s.repo.GetAllComments(post)
}
