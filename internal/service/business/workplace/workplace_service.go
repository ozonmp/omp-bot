package workplace

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"sort"
)

type WorkplaceService interface {
	Describe(workplaceID uint64) (*business.Workplace, error)
	List(cursor uint64, limit uint64) ([]business.Workplace, error)
	Create(business.Workplace) (uint64, error)
	Update(workplaceID uint64, workplace business.Workplace) error
	Remove(workplaceID uint64) (bool, error)
}

type DummyWorkplaceService struct{}

func NewDummyWorkplaceService() *DummyWorkplaceService {
	return &DummyWorkplaceService{}
}

func (s *DummyWorkplaceService) Describe(workplaceID uint64) (*business.Workplace, error) {
	workplace, ok := workplaceDB[workplaceID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Workplace with ID %d is not found", workplaceID))
	}

	return &workplace, nil
}

func (s *DummyWorkplaceService) List(cursor uint64, limit uint64) ([]business.Workplace, error) {
	var retVal = make([]business.Workplace, 0, len(workplaceDB))

	for idx, value := range workplaceDB {
		if idx > cursor && idx <= cursor + limit {
			retVal = append(retVal, value)
		}
	}

	sort.SliceStable(retVal, func(i, j int) bool {
		return retVal[i].ID < retVal[j].ID
	})

	return retVal, nil
}

func (s *DummyWorkplaceService) Create(workplace business.Workplace) (uint64, error) {
	if _, ok := workplaceDB[workplace.ID]; ok {
		return 0, errors.New(fmt.Sprintf("Workplace with ID %d is exists", workplace.ID))
	}

	workplaceDB[workplace.ID] = workplace
	return workplace.ID, nil
}

func (s *DummyWorkplaceService) Update(workplaceID uint64, workplace business.Workplace) error {
	if _, ok := workplaceDB[workplaceID]; !ok {
		return errors.New(fmt.Sprintf("Workplace with ID %d is not exists", workplaceID))
	}

	workplaceDB[workplaceID] = workplace
	return nil
}

func (s *DummyWorkplaceService) Remove(workplaceID uint64) (bool, error) {
	if _, ok := workplaceDB[workplaceID]; !ok {
		return false, errors.New(fmt.Sprintf("Workplace with ID %d is not found", workplaceID))
	}

	delete(workplaceDB, workplaceID)
	return true, nil
}

func (s *DummyWorkplaceService) GetDataSize() uint64 {
	return uint64(len(workplaceDB))
}
