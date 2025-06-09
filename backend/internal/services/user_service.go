package services

import (
	"fmt"
	"go-blog/backend/internal/models"
	"go-blog/backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserSevice(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	dbUser, _ := s.repo.GetByLogin(user.Login)

	if dbUser.Login != "" {
		return nil, fmt.Errorf("User with this login already registered")
	}

	createdUser, err := s.repo.Create(user)

	if err != nil {
		return nil, fmt.Errorf("Error creating user")
	}

	return createdUser, nil
}

func (s *UserService) Login(login, password string) (*models.User, error) {
	user, err := s.repo.GetByLogin(login)
	if err != nil {
		return nil, fmt.Errorf("Cannot find user %s", login)
	}

	if user.Password != password {
		return nil, fmt.Errorf("Incorrect password")
	}

	return user, nil
}
