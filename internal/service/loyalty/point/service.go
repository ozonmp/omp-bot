package point

import (
	"fmt"
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
)

type PointService interface {
	Describe(pointId uint64) (*loyalty.Point, error)
	List(cursor uint64, limit uint64) ([]loyalty.Point, error)
	Create(loyalty.Point) (uint64, error)
	Update(pointId uint64, point loyalty.Point) error
	Remove(pointId uint64) (bool, error)
}

type DummyPointService struct {
	allEntities []loyalty.Point
}

func NewDummyPointService() *DummyPointService {
	return &DummyPointService{allEntities: []loyalty.Point{
		{
			Id:          1,
			Name:        "One",
			Description: "Point one",
		},
		/*
		{
			Id:          2,
			Name:        "Two",
			Description: "Point two",
		},
		{
			Id:          3,
			Name:        "Three",
			Description: "Point tree",
		},
		{
			Id:          4,
			Name:        "Four",
			Description: "Point four",
		},
		{
			Id:          5,
			Name:        "Five",
			Description: "Point 5",
		},
		{
			Id:          6,
			Name:        "Six",
			Description: "Point 6",
		},
		{
			Id:          7,
			Name:        "Seven",
			Description: "Point 7",
		},
		{
			Id:          8,
			Name:        "Eight",
			Description: "Point 8",
		},
		{
			Id:          9,
			Name:        "Nine",
			Description: "Point 9",
		},
		*/
	}}
}

func (s *DummyPointService) Size() int {
	return len(s.allEntities)
}

func (s *DummyPointService) List() ([]loyalty.Point, error) {
	return s.allEntities[:], nil
}

func (s *DummyPointService) Get(pointId uint64) (*loyalty.Point, error) {
	for _, entity := range s.allEntities {
		if entity.Id == pointId {
			return &entity, nil
		}
	}

	return nil, fmt.Errorf("Can't find entity with id %d", pointId)
}

func (s *DummyPointService) Delete(pointId uint64) (bool, error) {
	for key, entity := range s.allEntities {
		if entity.Id == pointId {
			s.allEntities = append(s.allEntities[:key], s.allEntities[key+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("Can't find entity with id %d", pointId)
}

func (s *DummyPointService) Create(p loyalty.Point) (uint64, error) {
	if len(p.Name) == 0 {
		return 0, errors.New("field 'Name' is required")
	}

	if len(p.Description) == 0 {
		return 0, errors.New("field 'Description' is required")
	}

	s.allEntities = append(s.allEntities, p)

	return uint64(len(s.allEntities) - 1), nil
}

func (s *DummyPointService) Edit(pointId uint64, point loyalty.Point) error {
	if uint64(len(s.allEntities)-1) < pointId {
		return fmt.Errorf("Can't find entity with id %d", pointId)
	}

	if len(point.Name) > 0 {
		s.allEntities[pointId - 1].Name = point.Name
		errors.New("field 'Name' is required")
	}

	if len(point.Description) > 0 {
		s.allEntities[pointId -1].Description = point.Description
	}

	return nil
}