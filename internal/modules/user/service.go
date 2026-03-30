package user

import (
	"errors"

	"github.com/fiankasepman/go-gin-template/internal/auth"
	"github.com/fiankasepman/go-gin-template/internal/base"
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
func (s *Service) Login(username, password string) (*LoginResponse, error) {

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

	err = s.repo.UpdatesWhere(
		map[string]interface{}{"user_id": user.UserID},
		map[string]interface{}{"token": refreshToken},
		&user,
	)
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

	var user User

	err := s.repo.FindOneWhere(
		map[string]interface{}{"token": refreshToken},
		&user,
	)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	newAccessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (s *Service) Logout(userID string) error {

	var user User

	err := s.repo.FindOneWhere(
		map[string]interface{}{"user_id": userID},
		&user,
	)
	if err != nil {
		return err
	}

	return s.repo.UpdatesWhere(
		map[string]interface{}{"user_id": userID},
		map[string]interface{}{"token": nil},
		&user,
	)
}