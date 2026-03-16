package user

import (
	"github.com/fiankasepman/go-gin-template/internal/auth"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]User, error) {
	var users []User
	err := s.repo.FindAll(&users)
	return users, err
}

func (s *Service) Login(username, password string) (*User, error) {

	var user User

	err := s.repo.FindByUsername(username, &user)
	if err != nil {
		return nil, err
	}

	if !auth.CheckPassword(password, user.Password) {
		return nil, err
	}

	return &user, nil
}