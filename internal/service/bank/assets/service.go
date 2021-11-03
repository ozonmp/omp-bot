package assets

import (
	"errors"
	bank "github.com/ozonmp/omp-bot/internal/model/bank/assets"
	"math"
	"time"
)

type AssetsService interface {
	Describe(ID uint64) (*bank.Asset, error)
	List(cursor uint64, limit uint64) ([]bank.Asset, error)
	Create(userID uint64, price float64) (uint64, error)
	Update(userID uint64, price float64) error
	Remove(ID uint64) (bool, error)
	Count() uint64
}

type DummyAssetsService struct{
	CurrentID uint64
}

func NewDummyAssetsService() *DummyAssetsService {
	return &DummyAssetsService{
		CurrentID: 1,
	}
}

func (s *DummyAssetsService) isEmpty() bool {
	return len(bank.AllEntities) == 0
}

func (s *DummyAssetsService) List(cursor int64) []bank.Asset {
	leftIndex := (cursor-1)*int64(bank.PageSize)
	rightIndex := cursor*int64(bank.PageSize)
	len := int64(len(bank.AllEntities))
	if len < rightIndex {
		return bank.AllEntities[leftIndex:len]
	}
	return bank.AllEntities[leftIndex:rightIndex]
}

func (s *DummyAssetsService) Describe(searchId uint64) (*bank.Asset, error) {
	for i := 0; i < len(bank.AllEntities); i++ {
		if searchId == bank.AllEntities[i].ID {
			return &bank.AllEntities[i], nil
		}
	}
	return nil, errors.New("ID not found")
}

func (s *DummyAssetsService) Remove(removeId uint64) error {
	for i := 0; i < len(bank.AllEntities); i++ {
		if removeId == bank.AllEntities[i].ID {
			bank.AllEntities = append(bank.AllEntities[:i], bank.AllEntities[i+1:]...)
			return nil
		}
	}
	return errors.New("ID not found")
}

func (s *DummyAssetsService) Create(userId uint64, price float64) uint64 {
	newID := s.CurrentID
	s.CurrentID++

	bank.AllEntities = append(
		bank.AllEntities,
		bank.Asset{
			ID:       newID,
			CreatedAt: time.Now(),
			User: userId,
			PriceWhenCreated: price,
			CurrentPrice: price,
		},
	)

	return newID
}

func (s *DummyAssetsService) Update(userId uint64, price float64) error {
	asset, err := s.Describe(userId)
	if err != nil {
		return err
	}

	asset.CurrentPrice = price
	return nil
}

func (s *DummyAssetsService) Count() int {
	return len(bank.AllEntities)
}

func (s *DummyAssetsService) PageCount() int64 {
	return int64(math.Ceil(float64(s.Count()) / bank.PageSize))
}

