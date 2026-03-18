package user

import (
	"errors"

	"github.com/fiankasepman/go-gin-template/internal/auth"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

func (s *Service) GetAll() ([]User, error) {
	var users []User
	err := s.repo.FindAll(&users)
	return users, err
}

func (s *Service) Login(username, password string) (*LoginResponse, error) {

	var user User

	err := s.repo.FindByUsername(username, &user)
	if err != nil {
		return nil, err
	}

	if !auth.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:  &user,
		Token: token,
	}, nil
}