package point

import (
	"fmt"
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
		{
			Id:          2,
			Name:        "Two",
			Description: "Point two",
		},
		{
			Id:          3,
			Name:        "three",
			Description: "Point tree",
		},
		{
			Id:          4,
			Name:        "four",
			Description: "Point four",
		},
		{
			Id:          5,
			Name:        "five",
			Description: "Point 5",
		},
		{
			Id:          6,
			Name:        "six",
			Description: "Point 6",
		},
		{
			Id:          7,
			Name:        "seven",
			Description: "Point 7",
		},
		{
			Id:          8,
			Name:        "eight",
			Description: "Point 8",
		},
		{
			Id:          9,
			Name:        "eight",
			Description: "Point 9",
		},
	}}
}

func (s *DummyPointService) List(cursor uint64, limit uint64) ([]loyalty.Point, error) {
	// когда сущеностей осталось меньше, чем лимит на выдачу, но их надо показать
	if uint64(len(s.allEntities)) > cursor && uint64(len(s.allEntities)) < cursor+limit {
		return s.allEntities[cursor:], nil
	}

	if uint64(len(s.allEntities)) < cursor || uint64(len(s.allEntities)) < cursor+limit {
		return nil, fmt.Errorf("the requested page is out of range")
	}

	return s.allEntities[cursor : cursor+limit], nil
}

func (s *DummyPointService) Get(pointId uint64) (*loyalty.Point, error) {
	for _, entity := range s.allEntities {
		if entity.Id == pointId {
			return &entity, nil
		}
	}

	return nil, fmt.Errorf("entity with id %d not found", pointId)
}