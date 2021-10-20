package visit

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/activity"
)

type VisitService interface {
	Describe(visitID uint64) (*activity.Visit, error)
	List(cursor uint64, limit uint64) ([]activity.Visit, error)
	Create(activity.Visit) (uint64, error)
	Update(visitID uint64, visit activity.Visit) error
	Remove(visitID uint64) (bool, error)
	GetCount() int
}

func NewDummyVisitService() *DummyVisitService {
	return &DummyVisitService{}
}

type DummyVisitService struct{}

func (s *DummyVisitService) Describe(visitID uint64) (*activity.Visit, error) {
	return activity.Visits.FindById(visitID)
}

func (s *DummyVisitService) List(cursor uint64, limit uint64) ([]activity.Visit, error) {
	visitCount := uint64(len(activity.Visits))

	if cursor > visitCount {
		return []activity.Visit{}, errors.New("cursor out of visits range")
	}

	max := cursor + limit
	if max > visitCount {
		max = visitCount
	}

	return activity.Visits[cursor:max], nil
}

func (s *DummyVisitService) Create(Visit activity.Visit) (uint64, error) {
	if _, err := activity.Visits.FindById(Visit.Id); err == nil {
		return 0, errors.New("visit id must be unique")
	}

	activity.Visits = append(activity.Visits, Visit)

	return Visit.Id, nil
}

func (s *DummyVisitService) Update(visitId uint64, visit activity.Visit) error {
	foundVisit, err := activity.Visits.FindById(visitId)
	if err != nil {
		return err
	}

	foundVisit.Title = visit.Title

	return nil
}

func (s *DummyVisitService) Remove(visitId uint64) (bool, error) {
	id, err := activity.Visits.FindIndexById(visitId)
	if err != nil {
		return false, err
	}

	activity.Visits = append(activity.Visits[:id], activity.Visits[id+1:]...)

	return true, nil
}

func (s *DummyVisitService) GetCount() int {
	return len(activity.Visits)
}
