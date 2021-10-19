package verification

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/service"
	"sort"
)

type SubdomainService interface {
	Describe(subdomainID uint64) (*service.Verification, error)
	List(cursor uint64, limit uint64) ([]service.Verification, error)
	Create(service.Verification) (uint64, error)
	Update(subdomainID uint64, subdomain service.Verification) error
	Remove(subdomainID uint64) (bool, error)
}

type DummyVerificationService struct {}

func NewDummyVerificationService() *DummyVerificationService {
	return &DummyVerificationService{}
}

func (s *DummyVerificationService) Describe(itemID uint64) (*service.Verification, error) {
	item, ok := data[itemID]

	if !ok {
		errorMsg := errors.New(fmt.Sprintf("Search item by ID: %d Error ", itemID))
		return nil, errorMsg
	}

	return &item, nil
}

func (s *DummyVerificationService) List(position uint64, limit uint64) ([]service.Verification, error) {
	withdrawnItems := make([]service.Verification, 0)

	for idx, value := range data {
		if idx > position && idx <= position + limit {
			withdrawnItems = append(withdrawnItems, value)
		}
	}

	sort.SliceStable(withdrawnItems, func(i, j int) bool {
		return withdrawnItems[i].ID < withdrawnItems[j].ID
	})

	return withdrawnItems, nil
}


func (s *DummyVerificationService) Create(name string) service.Verification {
	item := service.Verification{ID: uint64(len(data)), Name: name}
	data[item.ID] = item
	return item
}

func (s *DummyVerificationService) Update(itemID uint64, name string) (service.Verification, error) {
	item := service.Verification{ID: itemID, Name: name}

	if _, existenceChecker := data[item.ID]; !existenceChecker {
		err := errors.New(fmt.Sprintf("Item ID: %d is not found", itemID))
		return item, err
	}

	data[item.ID] = item
	return item, nil
}

func (s *DummyVerificationService) Remove(itemID uint64) (bool, error) {

	if _, item := data[itemID]; !item {
		err := errors.New(fmt.Sprintf("Item ID: %d is not found", itemID))
		return false, err
	}

	delete(data, itemID)
	return true, nil
}

func (s *DummyVerificationService) GetDataLen() uint64 {
	return uint64(len(data))
}
