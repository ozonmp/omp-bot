package office

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type OfficeService interface {
	Describe(officeId uint64) (*business.Office, error)
	List(cursor uint64, limit uint64) ([]business.Office, error)
	Create(business.Office) (uint64, error)
	Update(officeId uint64, office business.Office) error
	Remove(officeId uint64) (bool, error)
}

type DummyOfficeService struct {
	allEntities []business.Office
}

func NewDummyOfficeService() *DummyOfficeService {
	return &DummyOfficeService{allEntities: []business.Office{
		{
			Name:        "One",
			Description: "Office one",
		},
		{
			Name:        "Two",
			Description: "Office two",
		},
		{
			Name:        "three",
			Description: "Office tree",
		},
		{
			Name:        "four",
			Description: "Office four",
		},
		{
			Name:        "five",
			Description: "Office 5",
		},
		{
			Name:        "six",
			Description: "Office 6",
		},
		{
			Name:        "seven",
			Description: "Office 7",
		},
		{
			Name:        "eight",
			Description: "Office 8",
		},
		{
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

func (s *DummyOfficeService) Get(officeId uint64) (*business.Office, error) {
	if uint64(len(s.allEntities)) < officeId {
		return nil, fmt.Errorf("entity with id %d not found", officeId)
	}

	return &s.allEntities[officeId], nil

}

func (s *DummyOfficeService) Delete(officeId uint64) (bool, error) {
	if uint64(len(s.allEntities)) < officeId {
		return false, fmt.Errorf("entity with id %d not found", officeId)
	}

	s.allEntities = append(s.allEntities[:officeId], s.allEntities[officeId+1:]...)
	return true, nil
}

//func (s *DummyOfficeService) Create(business.Office) (uint64, error) {
//	s.allEntities = append(s.allEntities[:key], s.allEntities[key+1:]...)
//
//	return false, fmt.Errorf("entity with id %d not found", office_id)
//}
