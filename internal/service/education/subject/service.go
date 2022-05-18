package subject

import (
	"errors"
	"fmt"
)

// ...
type DummyService struct {
	indexByID map[uint64]int
	subjects  []Subject
}

func NewDummyService() *DummyService {
	return &DummyService{
		indexByID: make(map[uint64]int),
		subjects:  make([]Subject, 0),
	}
}

func (service *DummyService) Describe(subjectID uint64) (*Subject, error) {
	idx, contains := service.indexByID[subjectID]
	if !contains || idx < 0 {
		return nil, errors.New(fmt.Sprintf(
			"No subject with id %v", subjectID,
		))
	}

	return &service.subjects[subjectID], nil
}

func (service *DummyService) List(cursor uint64, limit uint64) ([]Subject, error) {
	totalCount := uint64(len(service.subjects))
	if cursor > totalCount {
		return nil, errors.New(fmt.Sprintf(
			"incorrect cursor position %v, correct cursor positions are [0...%v]",
			cursor, totalCount,
		))
	}
	right := cursor + limit
	if right > totalCount {
		right = totalCount
	}
	return service.subjects[cursor:right], nil
}

func (service *DummyService) Create(subject Subject) (uint64, error) {
	newID := uint64(len(service.subjects))
	subject.ID = newID
	service.subjects = append(service.subjects, subject)
	service.indexByID[newID] = len(service.subjects) - 1
	return newID, nil
}

func (service *DummyService) Update(subjectID uint64, subject Subject) error {
	if subjectID != subject.ID {
		return errors.New("subjectID != subject.ID")
	}

	idx, contains := service.indexByID[subjectID]
	if !contains || idx < 0 {
		return errors.New(fmt.Sprintf(
			"no subject with id %v", subjectID,
		))
	}

	service.subjects[idx] = subject
	return nil
}

func (service *DummyService) Remove(subjectID uint64) (bool, error) {
	idx, contains := service.indexByID[subjectID]
	if !contains || idx < 0 {
		return false, errors.New(fmt.Sprintf(
			"no subject with id %v", subjectID,
		))
	}

	service.subjects = append(service.subjects[:idx], service.subjects[idx+1:]...)
	delete(service.indexByID, subjectID)

	return true, nil
}

func (service *DummyService) SubjectsCount() uint64 {
	return uint64(len(service.subjects))
}
