package group

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Group, error) {
	var data []Group
	err := s.repo.FindAll(&data)
	return data, err
}