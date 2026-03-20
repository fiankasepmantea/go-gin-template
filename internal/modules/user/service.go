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

// GET ALL
func (s *Service) GetAll() ([]User, error) {
	var users []User
	err := s.repo.FindAll(&users)
	return users, err
}

// CREATE
func (s *Service) Create(user *User) error {
	user.UserID = NewUserID()
	return s.repo.Create(user)
}

// UPDATE
func (s *Service) Update(user *User) error {
	return s.repo.Update(user)
}

// DELETE
func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

// LOGIN
func (s *Service) Login(username, password string) (*LoginResponse, error) {

	var user User
	err := s.repo.FindByUsername(username, &user)
	if err != nil {
		return nil, errors.New("user not found")
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