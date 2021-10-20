package warehouse

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

type EstateWarehouseService interface {
	Help() error
	Describe(idx uint64) (*Warehouse, error)
	List(offset uint64, limit uint64) ([]Warehouse, error)
	Delete(idx uint64) (bool, error)
	New(warehouse Warehouse) (uint64, error)
	Edit(id uint64, warehouse Warehouse) error
	Default() error
}

type DummyEstateService struct{}

func NewDummyEstateService() *DummyEstateService {
	return &DummyEstateService{}
}

func (s *DummyEstateService) Help() error {
	return nil
}

func (s *DummyEstateService) validateId(id uint64) error {
	if id < 1 {
		return errors.New("Warehouse id have to be >= 1")
	}

	if uint64(id) > uint64(len(allWarehouses)) {
		return fmt.Errorf("There is no warehouse with id = %v", id)
	}

	return nil
}

func (s *DummyEstateService) Describe(id uint64) (*Warehouse, error) {
	log.Printf("Estate Service: getting warehouse with id %v", id)
	if err := s.validateId(id); err != nil {
		return nil, err
	}
	return &allWarehouses[id-1], nil
}

func (s *DummyEstateService) List(offset uint64, limit uint64) ([]Warehouse, error) {
	var totalLen = uint64(len(allWarehouses))
	/// limit not used now

	limit = 0
	if offset >= totalLen {
		return []Warehouse{}, nil
	}

	res := make([]Warehouse, 0, totalLen)
	for _, wrh := range allWarehouses {
		res = append(res, wrh)
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	return res, nil
}

func (s *DummyEstateService) Delete(id uint64) (bool, error) {
	if err := s.validateId(id); err != nil {
		return false, err
	}

	index := id - 1
	allWarehouses = append(allWarehouses[:index], allWarehouses[index+1:]...)
	return true, nil
}

func (s *DummyEstateService) New(warehouse Warehouse) (uint64, error) {
	allWarehouses = append(allWarehouses, warehouse)
	return uint64(len(allWarehouses)), nil
}

func (s *DummyEstateService) Edit(id uint64, warehouse Warehouse) error {
	if err := s.validateId(id); err != nil {
		return err
	}

	allWarehouses[id-1] = warehouse
	return nil
}

func (s *DummyEstateService) Default() error {
	return nil
}
