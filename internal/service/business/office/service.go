package office

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

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
			Name:        "nine",
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

func (s *DummyOfficeService) Describe(officeId uint64) (*business.Office, error) {
	if uint64(len(s.allEntities)-1) < officeId {
		return nil, fmt.Errorf("entity with id %d not found", officeId)
	}

	return &s.allEntities[officeId], nil

}

func (s *DummyOfficeService) Remove(officeId uint64) (bool, error) {
	if uint64(len(s.allEntities)-1) < officeId {
		return false, fmt.Errorf("entity with id %d not found", officeId)
	}

	s.allEntities = append(s.allEntities[:officeId], s.allEntities[officeId+1:]...)
	return true, nil
}

func (s *DummyOfficeService) Create(o business.Office) (uint64, error) {
	if len(o.Name) == 0 {
		return 0, errors.New("field 'Name' is required")
	}

	if len(o.Description) == 0 {
		return 0, errors.New("field 'Description' is required")
	}

	s.allEntities = append(s.allEntities, o)

	return uint64(len(s.allEntities) - 1), nil
}

func (s *DummyOfficeService) Update(officeId uint64, office business.Office) error {
	if uint64(len(s.allEntities)-1) < officeId {
		return fmt.Errorf("entity with id %d not found", officeId)
	}

	if len(office.Name) > 0 {
		s.allEntities[officeId].Name = office.Name
		errors.New("field 'Name' is required")
	}

	if len(office.Description) > 0 {
		s.allEntities[officeId].Description = office.Description
	}

	return nil
}
