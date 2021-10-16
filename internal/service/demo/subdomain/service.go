package subdomain

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Entity {
	return allEntities
}

func (s *Service) Get(idx int) (*Entity, error) {
	return &allEntities[idx], nil
}
