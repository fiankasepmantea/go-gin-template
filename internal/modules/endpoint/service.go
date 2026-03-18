package endpoint

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CheckAccess(path, method string) (*Endpoint, error) {

	var ep Endpoint
	err := s.repo.FindByPath(path, method, &ep)

	return &ep, err
}