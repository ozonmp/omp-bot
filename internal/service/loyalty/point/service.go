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
			Id:          0,
			Name:        "One",
			Description: "Point one",
		},
		{
			Id:          1,
			Name:        "Two",
			Description: "Point two",
		},
		{
			Id:          2,
			Name:        "Three",
			Description: "Point tree",
		},
		{
			Id:          3,
			Name:        "Four",
			Description: "Point four",
		},
		{
			Id:          4,
			Name:        "Five",
			Description: "Point 5",
		},
		{
			Id:          5,
			Name:        "Six",
			Description: "Point 6",
		},
		{
			Id:          6,
			Name:        "Seven",
			Description: "Point 7",
		},
		{
			Id:          7,
			Name:        "Eight",
			Description: "Point 8",
		},
		{
			Id:          8,
			Name:        "Nine",
			Description: "Point 9",
		},
	}}
}

func (s *DummyPointService) Size() int {
	return len(s.allEntities)
}

func (s *DummyPointService) List(cursor int, limit int) ([]loyalty.Point, error) {
	if len(s.allEntities) == 0 {
		return nil, errors.New("List is emty")
	}else if len(s.allEntities) <= cursor || cursor < 0 {
		return nil, errors.New("Cursor out of list bound")
	}

	if len(s.allEntities) > cursor && len(s.allEntities) < cursor+limit {
		return s.allEntities[cursor:], nil
	}else{
		return s.allEntities[cursor : cursor+limit], nil
	}

}

func (s *DummyPointService) Get(pointId uint64) (*loyalty.Point, error) {
	for _, entity := range s.allEntities {
		if entity.Id == pointId {
			return &entity, nil
		}
	}

	return nil, fmt.Errorf("Can't find entity with id %d", (pointId + 1))
}

func (s *DummyPointService) Delete(pointId uint64) (bool, error) {
	for idx, entity := range s.allEntities {
		if entity.Id == pointId {
			s.allEntities = append(s.allEntities[:idx], s.allEntities[idx+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("Can't find entity with id %d", (pointId + 1))
}

func (s *DummyPointService) Create(p loyalty.Point) (uint64, error) {
	if len(p.Name) == 0 {
		return 0, errors.New("field 'Name' is required")
	}

	if len(p.Description) == 0 {
		return 0, errors.New("field 'Description' is required")
	}

	p.Id = uint64(len(s.allEntities))

	s.allEntities = append(s.allEntities, p)

	return uint64(len(s.allEntities) - 1), nil
}

func (s *DummyPointService) Edit(pointId uint64, point loyalty.Point) error {

	index := -1

	for idx, entity := range s.allEntities {
		if entity.Id == pointId {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Can't find entity with id %d", (pointId + 1))
	}

	if len(point.Name) > 0 {
		s.allEntities[index].Name = point.Name
	} else {
		return errors.New("field 'Name' is required")
	}

	if len(point.Description) > 0 {
		s.allEntities[index].Description = point.Description
	} else {
		return errors.New("field 'Description' is required")
	}

	return nil
}