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
type DeviceResponse struct {
	ID        string     `json:"id"`
	Device    *string    `json:"device"`
	UserAgent *string    `json:"user_agent"`
	IPAddress *string    `json:"ip_address"`
	ExpiresAt time.Time  `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
	Current   bool       `json:"current"`
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

	
	refreshToken := idgen.NewRefreshToken()
	tokenID := idgen.NewUserTokenID()

	token := usertoken.UserToken{
		ID:           tokenID,
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

	accessToken, err := auth.GenerateToken(user.UserID, tokenID)
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

	return auth.GenerateToken(token.UserID, token.ID)
}
func (s *Service) Logout(refreshToken string) error {
	return s.tokenRepo.DeleteByToken(refreshToken)
}

func (s *Service) LogoutAll(userID string) error {
	return s.tokenRepo.DeleteByUser(userID)
}

func (s *Service) GetDevices(userID, currentTokenID string) ([]DeviceResponse, error) {

	var tokens []usertoken.UserToken

	err := s.tokenRepo.FindByUser(userID, &tokens)
	if err != nil {
		return nil, err
	}

	var result []DeviceResponse

	for _, t := range tokens {

		result = append(result, DeviceResponse{
			ID:        t.ID,
			Device:    t.Device,
			UserAgent: t.UserAgent,
			IPAddress: t.IPAddress,
			ExpiresAt: t.ExpiresAt,
			CreatedAt: t.CreatedAt,
			Current:   t.ID == currentTokenID,
		})
	}

	return result, nil
}
func (s *Service) RevokeDevice(userID, tokenID string) error {

	// optional: validasi ownership
	var token usertoken.UserToken

	err := s.tokenRepo.FindByID(tokenID, &token)
	if err != nil {
		return errors.New("device not found")
	}

	if token.UserID != userID {
		return errors.New("unauthorized device")
	}

	return s.tokenRepo.DeleteByID(tokenID)
}