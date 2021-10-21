package operation

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/bank"
	"sort"
)

type ServiceInterface interface {
	Describe(operationID uint64) (*bank.Operation, error)
	List(cursor uint64, limit uint64) ([]bank.Operation, error)
	Create(bank.Operation) (uint64, error)
	Update(operationID uint64, operation bank.Operation) error
	Remove(operationID uint64) (bool, error)
}

type DummyService struct {
	operations map[uint64]*bank.Operation
	idCounter  uint64
	ids        []uint64
}

func NewDummyService() *DummyService {
	service:= &DummyService{
		operations: map[uint64]*bank.Operation{},
		idCounter:  1,
		ids:        []uint64{},
	}

	service.Create(bank.NewOperation("Salary", 1))
	service.Create(bank.NewOperation("Scholarship", 2))
	service.Create(bank.NewOperation("Restaurant", 3))
	service.Create(bank.NewOperation("Transfer", 4))
	service.Create(bank.NewOperation("Electricity", 5))
	service.Create(bank.NewOperation("Tax", 6))
	service.Create(bank.NewOperation("Tickets", 7))

	return service
}

func (s *DummyService) Describe(operationID uint64) (*bank.Operation, error) {
	op, ok := s.operations[operationID]
	if !ok {
		return nil, fmt.Errorf("operation with ID %v not found", operationID)
	}

	return op, nil
}

func (s *DummyService) List(cursor uint64, limit uint64) ([]bank.Operation, error) {
	if limit == 0 {
		return []bank.Operation{}, nil
	}

	var elementsFound uint64

	result := make([]bank.Operation, 0, limit)
	for i := cursor; i < uint64(len(s.ids)); i++ {
		id := s.ids[i]
		result = append(result, *s.operations[id])
		elementsFound++

		if elementsFound == limit {
			break
		}
	}

	return result, nil
}

func (s *DummyService) Create(operation bank.Operation) (uint64, error) {
	operation.ID = s.idCounter

	s.operations[s.idCounter] = &operation
	s.ids = append(s.ids, s.idCounter)
	s.idCounter++

	return operation.ID, nil
}

func (s *DummyService) Update(operationID uint64, operation bank.Operation) error {
	op, ok := s.operations[operationID]

	if !ok {
		return fmt.Errorf("operation with ID %v not found", operationID)
	}

	op.TransactionID = operation.TransactionID
	op.Status = operation.Status
	op.OperationType = operation.OperationType
	op.ProceedAt = operation.ProceedAt

	return nil
}

func (s *DummyService) Remove(operationID uint64) (bool, error) {
	_, ok := s.operations[operationID]

	if !ok {
		return false, fmt.Errorf("operation with ID %v not found", operationID)
	}

	idx := s.findIndexByID(operationID)
	delete(s.operations, operationID)
	if idx == 0 {
		s.ids = s.ids[1:]
	} else if idx == len(s.ids) {
		s.ids = s.ids[:idx]
	} else {
		s.ids = append(s.ids[:idx], s.ids[idx+1:]...)
	}

	return true, nil
}

func (s *DummyService) findIndexByID(operationID uint64) int {
	return sort.Search(len(s.ids), func(i int) bool {
		return s.ids[i] >= operationID
	})
}
