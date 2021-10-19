package platform

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"sort"
)

type DummyPlatformService struct{}

func (s *DummyPlatformService) Describe(platformID uint64) (*education.Platform, error) {
	platform, ok := allPlatforms[platformID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("platform with ID %d not found", platformID))
	}

	return &platform, nil
}

func (s *DummyPlatformService) List(cursor uint64, limit uint64) ([]education.Platform, error) {
	count := s.count()
	if cursor > count {
		return []education.Platform{}, nil
	}

	v := make([]education.Platform, 0, len(allPlatforms))

	for _, value := range allPlatforms {
		v = append(v, value)
	}

	sort.SliceStable(v, func(i, j int) bool {
		return v[i].ID < v[j].ID
	})

	high := cursor + limit
	if high > count {
		high = count
	}

	return v[cursor:high], nil
}

func (s *DummyPlatformService) Create(platform education.Platform) (uint64, error) {
	nextID := s.nextID()
	platform.ID = nextID

	allPlatforms[nextID] = platform

	return nextID, nil
}

func (s *DummyPlatformService) Update(platformID uint64, platform education.Platform) error {
	if _, ok := allPlatforms[platformID]; !ok {
		return errors.New(fmt.Sprintf("platform with ID %d is not exists", platformID))
	}

	platform.ID = platformID
	allPlatforms[platformID] = platform

	return nil
}

func (s *DummyPlatformService) Remove(platformID uint64) (bool, error) {
	if _, ok := allPlatforms[platformID]; !ok {
		return false, errors.New(fmt.Sprintf("platform with ID %d is not exists", platformID))
	}

	delete(allPlatforms, platformID)

	return true, nil
}

func (s *DummyPlatformService) nextID() uint64 {
	platformSequence++

	return platformSequence
}

func (s *DummyPlatformService) count() uint64 {
	return uint64(len(allPlatforms))
}

func NewDummyPlatformService() *DummyPlatformService {
	return &DummyPlatformService{}
}
