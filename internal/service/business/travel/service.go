package travel

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Travel {
	return allEntities
}

func (s *Service) Get(idx int) (*Travel, error) {
	return &allEntities[idx], nil
}
