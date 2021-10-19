package work

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Employee {
	return allEntities
}

func (s *Service) Get(idx int) (*Employee, error) {
	return &allEntities[idx], nil
}
