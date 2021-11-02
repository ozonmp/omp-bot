package assets

import "fmt"

/*import "github.com/ozonmp/omp-bot/internal/model/bnk"

type AssetsService interface {
	Describe(assetsID uint64) (*bnk.Assets, error)
	List(cursor uint64, limit uint64) ([]bnk.Assets, error)
	Create(bnk.Assets) (uint64, error)
	Update(assetsID uint64, assets bnk.Assets) error
	Remove(assetsID uint64) (bool, error)
}*/

type DummyAssetsService struct {}

func NewDummyAssetsService() *DummyAssetsService {
	return &DummyAssetsService{}
}

type Service struct{
	CurrentPage uint64
	CurrentID uint64
}

func NewService() *Service {
	return &Service{
		CurrentPage: 0,
		CurrentID: 1,
	}
}

func (s *Service) isEmpty() bool {
	return len(allEntities) == 0
}

func (s *Service) List() map[uint64]Asset {
	return allEntities
}

func (s *Service) Get(idx uint64) (*Asset, error) {
	res, ok := allEntities[idx]
	if !ok {
		return nil, fmt.Errorf("no such index")
	}

	return &res, nil
}

func (s *Service) Remove(idx uint64) error {
	if !s.check(idx) {
		return fmt.Errorf("no such index")
	}

	delete(allEntities, idx)
	return nil
}

func (s *Service) New(name string, money float64) (uint64, error) {
	newID := s.CurrentID
	s.CurrentID++

	allEntities[newID] = Asset{
		ID:       newID,
		Username: name,
		Balance:  money,
	}

	if s.check(newID) {
		return newID, nil
	}
	return 0, fmt.Errorf("сreating error")
}

func (s *Service) Edit(ID uint64, name string, money float64) error {
	allEntities[ID] = Asset{
		ID:       ID,
		Username: name,
		Balance:  money,
	}

	if s.check(ID) {
		return nil
	}
	return fmt.Errorf("editing error")
}

func (s *Service) check(idx uint64) bool {
	_, ok := allEntities[idx]
	return ok
}


