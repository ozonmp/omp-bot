package ground

import (
	"errors"
	"fmt"
)

type DummyGroundService struct{}

func NewDummyGroundService() *DummyGroundService {
	return &DummyGroundService{}
}

func (s *DummyGroundService) Describe(idx int64) (*Ground, error) {
	if err := s.checkSliceIndex(idx); err != nil {
		return nil, err
	}
	return &allGrounds[idx], nil
}

func (s *DummyGroundService) List(cursor int64, limit int64) ([]Ground, error) {
	if err := s.checkSliceIndex(cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}
	min := min(cursor+limit, int64(len(allGrounds)))
	return allGrounds[cursor:min], nil
}

func (s *DummyGroundService) Create(ground Ground) (int64, error) {
	allGrounds = append(allGrounds, ground)
	return int64(len(allGrounds)) - 1, nil
}

func (s *DummyGroundService) Update(idx int64, ground Ground) error {
	if err := s.checkSliceIndex(idx); err != nil {
		return err
	}
	allGrounds[idx] = ground
	return nil
}

func (s *DummyGroundService) Remove(idx int64) (bool, error) {
	if err := s.checkSliceIndex(idx); err != nil {
		return false, err
	}

	allGrounds = append(allGrounds[:idx], allGrounds[idx+1:]...)
	return true, nil
}

func (s *DummyGroundService) checkSliceIndex(idx int64) error {
	if idx >= int64(len(allGrounds)) {
		return errors.New(
			fmt.Sprintf("index %d is out of range 0 - %d", idx, len(allGrounds)),
		)
	}
	return nil
}

func min(a int64, b int64) int64 {
	if a > b {
		return b
	}
	return a
}
