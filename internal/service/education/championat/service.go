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

func (s *DummyChampionatService) Create(championat Championat) error {
	allChampionats = append(allChampionats, championat)
	return nil
}

func (s *DummyChampionatService) Update(championatId uint64, championat Championat) error {
	allChampionats[championatId] = championat
	return nil
}

func (s *DummyChampionatService) Remove(championatId uint64) error {
	allChampionats = append(allChampionats[:championatId], allChampionats[championatId+1:]...)
	return nil
}

func (ch *Championat) String() string {
	return "Title: " + ch.Title
}
