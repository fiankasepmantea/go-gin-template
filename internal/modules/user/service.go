package user

import (
	"errors"
	"time"

	"github.com/fiankasepman/go-gin-template/internal/auth"
	"github.com/fiankasepman/go-gin-template/internal/base"
	usertoken "github.com/fiankasepman/go-gin-template/internal/modules/user_token"
	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
)

type Service struct {
	repo *Repository
	tokenRepo  *usertoken.Repository
}

func NewService(repo *Repository, tokenRepo *usertoken.Repository) *Service {
	return &Service{
		repo:      repo,
		tokenRepo: tokenRepo,
	}
}

type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Service) GetAll(limit, offset int) (interface{}, error) {
	var data []User

	err := s.repo.Paginate(limit, offset, &data)
	if err != nil {
		return nil, err
	}

	total, err := s.repo.Count()
	if err != nil {
		return nil, err
	}

	page := 1
	if limit > 0 {
		page = (offset / limit) + 1
	}

	return base.BuildPagination(page, limit, total, data), nil
}

func (s *Service) Create(user *User) error {
	user.UserID = NewUserID()
	return s.repo.Create(user)
}

func (s *Service) Update(user *User) error {
	return s.repo.Update(user)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
func (s *Service) Login(username, password, device, ua, ip string) (*LoginResponse, error) {

	var user User

	err := s.repo.FindByUsername(username, &user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !auth.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	accessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken := idgen.NewRefreshToken()

	token := usertoken.UserToken{
		ID:           idgen.NewUserTokenID(),
		UserID:       user.UserID,
		RefreshToken: refreshToken,
		Device:       &device,
		UserAgent:    &ua,
		IPAddress:    &ip,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
	}

	err = s.tokenRepo.Create(&token)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:         &user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *Service) Refresh(refreshToken string) (string, error) {

	var token usertoken.UserToken

	err := s.tokenRepo.FindByToken(refreshToken, &token)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {
		return "", errors.New("refresh token expired")
	}

	return auth.GenerateToken(token.UserID)
}
func (s *Service) Logout(refreshToken string) error {
	return s.tokenRepo.DeleteByToken(refreshToken)
}

func (s *Service) LogoutAll(userID string) error {
	return s.tokenRepo.DeleteByUser(userID)
}