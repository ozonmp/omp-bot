package notification

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Notification {
	return allEntities
}

func (s *Service) Get(idx int) (*Notification, error) {
	return &allEntities[idx], nil
}
