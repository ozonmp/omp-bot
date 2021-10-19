package company

type DummyCompanyService struct{}

func NewDummyCompanyService() *DummyCompanyService {
	return &DummyCompanyService{}
}

func (s *DummyCompanyService) Get(idx int) (*Company, error) {
	return &allEntities[idx], nil
}
