package user

import (
	"errors"

	"github.com/fiankasepman/go-gin-template/internal/auth"
	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
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

	// access token
	accessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken := idgen.NewRefreshToken()

	// simpan ke DB
	user.Token = &refreshToken
	s.repo.Update(&user)

	return &LoginResponse{
		User:         &user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Refresh(refreshToken string) (string, error) {

	var user User

	err := s.repo.DB.Where("token = ?", refreshToken).First(&user).Error
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// generate access token baru
	newAccessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *Service) Logout(userID string) error {

	var user User

	err := s.repo.DB.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return err
	}

	user.Token = nil

	return s.repo.Update(&user)
}