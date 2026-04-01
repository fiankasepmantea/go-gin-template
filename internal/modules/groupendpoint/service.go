package groupendpoint

import (
	"github.com/fiankasepman/go-gin-template/internal/cache"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}
func (s *Service) Assign(groupID, endpointID string) error {

	data := GroupEndpoint{
		ID:         NewGroupEndpointID(),
		GroupID:    groupID,
		EndpointID: endpointID,
	}

	err := s.repo.Create(&data)
	if err != nil {
		return err
	}

	s.clearCache(groupID)

	return nil
}

func (s *Service) Remove(groupID, endpointID string) error {

	err := s.repo.DeleteByGroupAndEndpoint(groupID, endpointID)
	if err != nil {
		return err
	}

	// 🔥 CLEAR CACHE GROUP
	s.clearCache(groupID)

	return nil
}

func (s *Service) GetByGroup(groupID string) ([]GroupEndpoint, error) {

	var data []GroupEndpoint
	err := s.repo.FindByGroup(groupID, &data)
	return data, err
}

func (s *Service) clearCache(groupID string) {

	iter := cache.RDB.Scan(cache.Ctx, 0, "rbac:"+groupID+":*", 0).Iterator()

	for iter.Next(cache.Ctx) {
		cache.RDB.Del(cache.Ctx, iter.Val())
	}
}