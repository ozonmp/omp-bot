package click

import (
	model "github.com/ozonmp/omp-bot/internal/model/actvity"
)

type DummyClickService struct{}

func NewDummyClickService() *DummyClickService {
	return &DummyClickService{}
}

func (s *DummyClickService) Describe(subdomainId uint64) (*model.Click, error) {
	return nil, nil
}

func (s *DummyClickService) List(cursor uint64, limit uint64) []*model.Click {
	return []*model.Click{}
}

func (s *DummyClickService) Create(m model.Click) (uint64, error) {
	return 0, nil
}

func (s *DummyClickService) Update(subdomainId uint64, subdomain model.Click) error {
	return nil
}

func (s *DummyClickService) Remove(subdomainId uint64) (bool, error) {
	return false, nil
}
