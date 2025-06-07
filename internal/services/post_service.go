package services

import (
	"go-blog/internal/models"
	"go-blog/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) *PostService {
	return &PostService{repo: r}
}

func (s *PostService) CreatePost(userId uint, post *models.Post) error {
	post.UserID = userId
	return s.repo.Create(post)
}

func (s *PostService) EditPost(id uint, post *models.Post) error {
	return s.repo.Edit(id, post)
}

func (s *PostService) DeletePost(id uint) error {
	return s.repo.Delete(id)
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.GetPostWithComments(id)
}

func (s *PostService) GetAllPosts() (*[]models.Post, error) {
	return s.repo.GetAllPosts()
}
