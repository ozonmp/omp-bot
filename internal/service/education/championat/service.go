package championat

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Championat {
	return allEntities
}

func (s *Service) Get(idx int) (*Championat, error) {
	return &allEntities[idx], nil
}

func (s *Service) New(title string) error {
	newEntity(title)
	return nil
}

func (s *Service) Delete(id int) error {
	deleteEntity(id)
	return nil
}

func (s *Service) Edit(id int, title string) error {
	editEntity(id, title)
	return nil
}
