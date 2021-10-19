package exchange

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

type ExchangeService interface {
	Describe(exchangeID uint64) (*exchange.Exchange, error)
	List(cursor uint64, limit uint64) ([]exchange.Exchange, error)
	Create(exchange exchange.Exchange) (uint64, error)
	Update(exchangeID uint64, exchange exchange.Exchange) error
	Remove(exchangeID uint64) (bool, error)
}

type DummyExchangeService struct {
	accessor *exchange.ModelAccessor
}

func NewDummyExchangeService() *DummyExchangeService {
	return &DummyExchangeService{
		accessor: exchange.NewModelAccessor(),
	}
}

func (dummy *DummyExchangeService) Describe(exchangeID uint64) (*exchange.Exchange, error) {
	entity, ok := dummy.accessor.Get(exchangeID)
	if !ok {
		errorMsg := fmt.Sprintf("Cannot find exchange request by id = %d", exchangeID)
		return &exchange.Exchange{}, errors.New(errorMsg)
	}
	return entity, nil
}

func (dummy *DummyExchangeService) List(cursor uint64, limit uint64) ([]exchange.Exchange, error) {
	exchages := dummy.accessor.Entities()
	return exchages, nil
}

func (dummy *DummyExchangeService) Create(exchange exchange.Exchange) (uint64, error) {
	return dummy.accessor.Add(exchange), nil
}

func (dummy *DummyExchangeService) Update(exchangeID uint64, exchange exchange.Exchange) error {

	return nil
}

func (dummy *DummyExchangeService) Remove(exchangeID uint64) (bool, error) {
	return dummy.accessor.Remove(exchangeID)
}
