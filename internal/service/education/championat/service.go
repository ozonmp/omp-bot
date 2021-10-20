package championat

type DummyChampionatService struct{}

func NewDummyChampionatService() *DummyChampionatService {
	return &DummyChampionatService{}
}

func (s *DummyChampionatService) Describe(championatId uint64) (*Championat, error) {
	return &allChampionats[championatId], nil
}

func (s *DummyChampionatService) List(cursor uint64, limit uint64) []Championat {
	var championatsLen = uint64(len(allChampionats))
	if cursor >= championatsLen {
		cursor = championatsLen
	}
	if (cursor + limit) >= championatsLen {
		limit = championatsLen - cursor
	}
	return allChampionats[cursor : cursor+limit]
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
