package ground

import (
	"errors"
	"fmt"
)

type DummyGroundService struct{}

func NewDummyGroundService() *DummyGroundService {
	return &DummyGroundService{}
}

func (s *DummyGroundService) Describe(idx uint64) (*Ground, error) {
	if err := s.checkSliceIndex(idx); err != nil {
		return nil, err
	}
	return &allGrounds[idx], nil
}

func (s *DummyGroundService) List(cursor uint64, limit uint64) ([]Ground, error) {
	if err := s.checkSliceIndex(cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}
	min := min(cursor+limit, uint64(len(allGrounds)))
	return allGrounds[cursor:min], nil
}

func (s *DummyGroundService) Create(ground Ground) (uint64, error) {
	allGrounds = append(allGrounds, ground)
	return uint64(len(allGrounds)) - 1, nil
}

func (s *DummyGroundService) Update(idx uint64, ground Ground) error {
	if err := s.checkSliceIndex(idx); err != nil {
		return err
	}
	allGrounds[idx] = ground
	return nil
}

func (s *DummyGroundService) Remove(idx uint64) (bool, error) {
	if err := s.checkSliceIndex(idx); err != nil {
		return false, err
	}

	allGrounds = append(allGrounds[:idx], allGrounds[idx+1:]...)
	return true, nil
}

func (s *DummyGroundService) checkSliceIndex(idx uint64) error {
	if idx >= uint64(len(allGrounds)) {
		return errors.New(
			fmt.Sprintf("index %d is out of range 0 - %d", idx, len(allGrounds)),
		)
	}
	return nil
}

func min(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}
