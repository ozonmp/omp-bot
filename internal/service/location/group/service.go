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

func (s *LocationGroupService) List(start, limit int) []location.Group {
	size := len(s.allEntities)
	switch {
	case start < 0:
		start = 0
	case start > size:
		return []location.Group{}
	}
	limit += start
	switch {
	case limit < 0:
		limit = 0
	case limit > size:
		limit = size
	}
	return s.allEntities[start:limit]
}

func (s *LocationGroupService) Get(idx int) (*location.Group, error) {
	if idx < 0 || idx >= len(s.allEntities) {
		return nil, fmt.Errorf("idx %d not found", idx)
	}
	return &s.allEntities[idx], nil
}

func (s *LocationGroupService) Delete(idx int) (err error) {
	if idx < 0 || idx >= len(s.allEntities) {
		err = fmt.Errorf("idx %d not found", idx)
		return
	}
	s.allEntities = append(s.allEntities[:idx], s.allEntities[idx+1:]...)
	return
}

func (s LocationGroupService) Size() int {
	return len(s.allEntities)
}