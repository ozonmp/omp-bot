package assets

import (
	"errors"
	bnk "github.com/ozonmp/omp-bot/internal/model/bnk/assets"
	"math"
	"time"
)

type AssetsService interface {
	Describe(ID uint64) (*bnk.Asset, error)
	List(cursor uint64, limit uint64) ([]bnk.Asset, error)
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
	return len(bnk.AllEntities) == 0
}

func (s *DummyAssetsService) List(cursor int64) []bnk.Asset {
	leftIndex := (cursor-1)*int64(bnk.PageSize)
	rightIndex := cursor*int64(bnk.PageSize)
	len := int64(len(bnk.AllEntities))
	if len < rightIndex {
		return bnk.AllEntities[leftIndex:len]
	}
	return bnk.AllEntities[leftIndex:rightIndex]
}

func (s *DummyAssetsService) Describe(searchId uint64) (*bnk.Asset, error) {
	for i := 0; i < len(bnk.AllEntities); i++ {
		if searchId == bnk.AllEntities[i].ID {
			return &bnk.AllEntities[i], nil
		}
	}
	return nil, errors.New("ID not found")
}

func (s *DummyAssetsService) Remove(removeId uint64) error {
	for i := 0; i < len(bnk.AllEntities); i++ {
		if removeId == bnk.AllEntities[i].ID {
			bnk.AllEntities = append(bnk.AllEntities[:i], bnk.AllEntities[i+1:]...)
			return nil
		}
	}
	return errors.New("ID not found")
}

func (s *DummyAssetsService) Create(userId uint64, price float64) uint64 {
	newID := s.CurrentID
	s.CurrentID++

	bnk.AllEntities = append(
		bnk.AllEntities,
		bnk.Asset{
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
	return len(bnk.AllEntities)
}

func (s *DummyAssetsService) PageCount() int64 {
	return int64(math.Ceil(float64(s.Count()) / bnk.PageSize))
}

