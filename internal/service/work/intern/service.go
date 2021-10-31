package intern

import (
	"fmt"
	"github.com/google/uuid"
)

type Service interface {
	Describe(internID uint64) (*Intern, error)
	List(cursor uint64, limit uint64) ([]Intern, error)
	Create(Intern) (uint64, error)
	Update(internID uint64, intern Intern) error
	Remove(internID uint64) (bool, error)
}

type InternService struct{}

func NewService() *InternService {
	return &InternService{}
}

func (s *InternService) List() []Intern {
	return allInterns
}

func (s *InternService) Describe(internID uint64) (*Intern, error) {
	_, intern, ok := getIntern(internID)
	if ok {
		return &intern, nil
	}
	return nil, fmt.Errorf("cannot find intern with id %d", internID)
}

func (s *InternService) Create(intern Intern) (uint64, error) {
	newUUID, err := uuid.NewUUID()
	intern.UniqueKey = newUUID
	intern.InternshipID = getNextInternshipId()
	allInterns = append(allInterns, intern)
	return intern.InternshipID, err
}

func (s *InternService) Update(internID uint64, intern Intern) error {
	idx, _, removed := getIntern(internID)
	if !removed {
		return fmt.Errorf("cannot find intern with id %d", internID)

	}

	allInterns[idx].Name = intern.Name
	return nil
}

func (s *InternService) Remove(internID uint64) (bool, error) {
	idx, _, ok := getIntern(internID)
	if ok {
		allInterns[idx] = allInterns[len(allInterns)-1]
		allInterns = allInterns[:len(allInterns)-1]
		return true, nil
	}
	return false, nil
}

func getIntern(internID uint64) (int, Intern, bool) {
	var res Intern
	for idx, intern := range allInterns {
		if intern.InternshipID == internID {
			res = intern
			return idx, res, true
		}
	}
	return -1, res, false
}
