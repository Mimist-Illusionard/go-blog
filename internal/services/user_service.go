package services

import (
	"fmt"
	"go-blog/internal/models"
	"go-blog/internal/repository"
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

func (s *UserService) CreateUser(user *models.User) error {
	dbUser, _ := s.repo.GetByLogin(user.Login)

	if dbUser.Login != "" {
		return fmt.Errorf("User with this login already registered")
	}

	return s.repo.Create(user)
}

func (s *UserService) Login(login, password string) error {
	user, err := s.repo.GetByLogin(login)
	if err != nil {
		return fmt.Errorf("Cannot find user %s", login)
	}

	if user.Password != password {
		return fmt.Errorf("Incorrect password")
	}

	return nil
}
