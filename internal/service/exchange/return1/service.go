package return1

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

type Return1Service interface {
	Describe(return1ID uint64) (*exchange.Return1, error)
	List(cursor uint64, limit uint64) ([]exchange.Return1, error)
	Create(exchange.Return1) (uint64, error)
	Update(return1ID uint64, return1 exchange.Return1) error
	Remove(return1ID uint64) (bool, error)
}

type DummyReturn1Service struct {
	savedReturns []exchange.Return1
}

func NewDummyReturn1Service() *DummyReturn1Service {
	s := DummyReturn1Service{}
	s.Create(exchange.Return1{Name: "zero"})
	s.Create(exchange.Return1{Name: "one"})
	s.Create(exchange.Return1{Name: "two"})
	s.Create(exchange.Return1{Name: "three"})
	s.Create(exchange.Return1{Name: "four"})
	s.Create(exchange.Return1{Name: "five"})
	s.Create(exchange.Return1{Name: "six"})

	return &s
}

var wrongIDErr = errors.New("wrong id")

func (r *DummyReturn1Service) Describe(return1ID uint64) (*exchange.Return1, error) {
	if uint64(len(r.savedReturns)) <= return1ID {
		return nil, wrongIDErr
	}

	return &r.savedReturns[return1ID], nil
}

var LastPageExceededErr = errors.New("last page exceeded")

func (r *DummyReturn1Service) List(cursor uint64, limit uint64) ([]exchange.Return1, error) {
	if len(r.savedReturns) == 0 {
		return nil, nil
	}

	var low, high uint64
	low = cursor
	high = cursor + limit

	if high > uint64(len(r.savedReturns)) {
		high = uint64(len(r.savedReturns))
	}

	if low >= uint64(len(r.savedReturns)) {
		return nil, LastPageExceededErr
	}

	return r.savedReturns[low:high], nil
}

func (r *DummyReturn1Service) Create(ret exchange.Return1) (uint64, error) {
	ret.ID = uint64(len(r.savedReturns))
	r.savedReturns = append(r.savedReturns, ret)

	return ret.ID, nil
}

func (r *DummyReturn1Service) Update(return1ID uint64, return1 exchange.Return1) error {
	if uint64(len(r.savedReturns)) <= return1ID {
		return wrongIDErr
	}

	r.savedReturns[return1ID] = return1
	r.savedReturns[return1ID].ID = return1ID

	return nil
}

func (r *DummyReturn1Service) Remove(return1ID uint64) (bool, error) {
	if uint64(len(r.savedReturns)) <= return1ID {
		return false, wrongIDErr
	}

	if len(r.savedReturns) == int(return1ID-1) {
		r.savedReturns = r.savedReturns[:len(r.savedReturns)-1]
		return true, nil
	}

	r.savedReturns = append(r.savedReturns[:return1ID], r.savedReturns[return1ID+1:]...)

	for i := 0; i < len(r.savedReturns); i++ {
		r.savedReturns[i].ID = uint64(i)
	}

	return true, nil
}
