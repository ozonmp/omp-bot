package announcement

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Announcement {
	return allEntities
}

func (s *Service) Get(idx int) (*Announcement, error) {
	return &allEntities[idx], nil
}