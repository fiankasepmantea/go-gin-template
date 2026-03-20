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

func (s *Service) Create(group *Group) error {
	group.GroupID = NewGroupID()
	return s.repo.Create(group)
}

func (s *Service) Update(group *Group) error {
	return s.repo.Update(group)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}