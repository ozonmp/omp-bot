package ground

import (
	"errors"
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

type GroundService interface {
	Describe(groundID uint64) (*autotransport.Ground, error)
	List(cursor uint64, limit uint64) ([]autotransport.Ground, error)
	Create(ground autotransport.Ground) (uint64, error)
	Update(groundID uint64, ground autotransport.Ground) error
	Remove(groundID uint64) (bool, error)
}

type DummyGroundService struct{}

func NewDummyGroundService() *DummyGroundService {
	return &DummyGroundService{}
}

func (s *DummyGroundService) Describe(groundID uint64) (*autotransport.Ground, error) {
	if err := s.checkSliceIndex(groundID); err != nil {
		return nil, err
	}
	allGrounds := *autotransport.AllGrounds()
	return &allGrounds[groundID], nil
}

func (s *DummyGroundService) List(cursor uint64, limit uint64) ([]autotransport.Ground, error) {
	if err := s.checkSliceIndex(cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}
	min := min(cursor+limit, uint64(len(*autotransport.AllGrounds())))
	return (*autotransport.AllGrounds())[cursor:min], nil
}

func (s *DummyGroundService) Create(ground autotransport.Ground) (uint64, error) {
	allGrounds := autotransport.AllGrounds()
	*allGrounds = append(*allGrounds, ground)
	length := len(*allGrounds)
	return uint64(length) - 1, nil
}

func (s *DummyGroundService) Update(groundID uint64, ground autotransport.Ground) error {
	if err := s.checkSliceIndex(groundID); err != nil {
		return err
	}
	allGrounds := *autotransport.AllGrounds()
	allGrounds[groundID] = ground
	return nil
}

func (s *DummyGroundService) Remove(groundID uint64) (bool, error) {
	if err := s.checkSliceIndex(groundID); err != nil {
		return false, err
	}

	allGrounds := autotransport.AllGrounds()
	*allGrounds = append((*allGrounds)[:groundID], (*allGrounds)[groundID+1:]...)
	return true, nil
}

func (s *DummyGroundService) checkSliceIndex(groundID uint64) error {
	allGrounds := autotransport.AllGrounds()
	length := len(*allGrounds)
	if groundID >= uint64(length) {
		return errors.New(
			fmt.Sprintf("index %d is out of range 0 - %d", groundID, length),
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
