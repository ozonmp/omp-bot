package internship

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Internship {
	return allEntities
}

func (s *Service) Get(idx int) (*Internship, error) {
	return &allEntities[idx], nil
}
