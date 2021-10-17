package office

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type OfficeService interface {
	Describe(office_id uint64) (*business.Office, error)
	List(cursor uint64, limit uint64) ([]business.Office, error)
	Create(business.Office) (uint64, error)
	Update(office_id uint64, office business.Office) error
	Remove(office_id uint64) (bool, error)
}

type DummyOfficeService struct {
	allEntities []business.Office
}

func NewDummyOfficeService() *DummyOfficeService {
	return &DummyOfficeService{allEntities: []business.Office{
		{
			Id:          1,
			Name:        "One",
			Description: "Office one",
		},
		{
			Id:          2,
			Name:        "Two",
			Description: "Office two",
		},
		{
			Id:          3,
			Name:        "three",
			Description: "Office tree",
		},
		{
			Id:          4,
			Name:        "four",
			Description: "Office four",
		},
		{
			Id:          5,
			Name:        "five",
			Description: "Office 5",
		},
		{
			Id:          6,
			Name:        "six",
			Description: "Office 6",
		},
		{
			Id:          7,
			Name:        "seven",
			Description: "Office 7",
		},
		{
			Id:          8,
			Name:        "eight",
			Description: "Office 8",
		},
		{
			Id:          9,
			Name:        "eight",
			Description: "Office 9",
		},
	}}
}

func (s *DummyOfficeService) List(cursor uint64, limit uint64) ([]business.Office, error) {
	// когда сущеностей осталось меньше, чем лимит на выдачу, но их надо показать
	if uint64(len(s.allEntities)) > cursor && uint64(len(s.allEntities)) < cursor+limit {
		return s.allEntities[cursor:], nil
	}

	if uint64(len(s.allEntities)) < cursor || uint64(len(s.allEntities)) < cursor+limit {
		return nil, fmt.Errorf("the requested page is out of range")
	}

	return s.allEntities[cursor : cursor+limit], nil
}

func (s *DummyOfficeService) Get(office_id uint64) (*business.Office, error) {
	for _, entity := range s.allEntities {
		if entity.Id == office_id {
			return &entity, nil
		}
	}

	return nil, fmt.Errorf("entity with id %d not found", office_id)
}
