package group

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/location"
)

type LocationGroupService struct {
	allEntities []location.Group
}

func NewLocationGroupService() *LocationGroupService {
	return &LocationGroupService{
		allEntities: []location.Group{
			{Id: 1, CountOfLocation: 2, Type: "building"},
			{Id: 2, CountOfLocation: 3, Type: "building"},
			{Id: 3, CountOfLocation: 3, Type: "room"},
			{Id: 4, CountOfLocation: 1, Type: "warehouse"},
			{Id: 5, CountOfLocation: 1, Type: "warehouse"},
		},
	}
}

func (s *LocationGroupService) List() *[]location.Group {
	return &s.allEntities
}

func (s *LocationGroupService) Get(idx int) (*location.Group, error) {
	if idx >= len(s.allEntities) {
		return nil, fmt.Errorf("idx %d not found", idx)
	}
	return &s.allEntities[idx], nil
}

func (s *LocationGroupService) Delete(idx int) (err error) {
	if idx >= len(s.allEntities) {
		err = fmt.Errorf("idx %d not found", idx)
		return
	}
	s.allEntities = append(s.allEntities[:idx], s.allEntities[idx+1:]...)
	return
}