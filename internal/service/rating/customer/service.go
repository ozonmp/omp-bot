package customer

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/general_errors"
)

type DummyService struct{}

func NewDummyService() *DummyService {
	return &DummyService{}
}

func (s *DummyService) List(cursor uint64, limit uint64) ([]Customer, error) {
	if err := s.validateIndex(cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}
	min := int(limit)
	if min > len(allEntities) {
		min = len(allEntities)
	}
	return allEntities[cursor:min], nil
}

func (s *DummyService) Describe(idx uint64) (*Customer, error) {
	if err := s.validateIndex(idx); err != nil {
		return nil, err
	}
	return &allEntities[idx], nil
}

func (s *DummyService) Remove(idx uint64) (bool, error) {
	if err := s.validateIndex(idx); err != nil {
		return false, err
	}

	if int(idx) == len(allEntities)-1 {
		allEntities = allEntities[:idx-1]
		return true, nil
	}

	copy(allEntities[idx:], allEntities[idx+1:])
	allEntities = allEntities[:len(allEntities)-1]
	return true, nil
}

func (s *DummyService) Create(customer Customer) (uint64, error) {
	allEntities = append(allEntities, customer)
	return uint64(len(allEntities) - 1), nil
}

func (s *DummyService) Update(idx uint64, customer Customer) error {
	if err := s.validateIndex(idx); err != nil {
		return err
	}
	allEntities[idx] = customer
	return nil
}

func (s *DummyService) validateIndex(idx uint64) error {
	if int(idx) >= len(allEntities) {
		return general_errors.NewUserError(fmt.Sprintf("index %d is out of range 0 - %d", idx, len(allEntities)))
	}
	return nil
}
