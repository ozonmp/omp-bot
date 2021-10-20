package championat

type DummyChampionatService struct{}

func NewDummyChampionatService() *DummyChampionatService {
	return &DummyChampionatService{}
}

func (s *DummyChampionatService) Describe(championatId uint64) (*Championat, error) {
	return &allChampionats[championatId], nil
}

func (s *DummyChampionatService) List() []Championat {
	return allChampionats
}

func (s *DummyChampionatService) Create(title string) error {
	newEntity(title)
	return nil
}

func (s *DummyChampionatService) Update(championatId uint64, title string) error {
	editEntity(championatId, title)
	return nil
}

func (s *DummyChampionatService) Remove(championatId uint64) error {
	deleteEntity(championatId)
	return nil
}
