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
	exchanges := dummy.accessor.Entities()
	exchangesLength := uint64(len(exchanges))
	if cursor >= exchangesLength {
		return nil, errors.New(fmt.Sprintf("Cursor has overran the last index of the list.\n"+
			"Last index = %d\nCursor position = %d\n", exchangesLength-1, cursor))
	}
	var last uint64
	if (cursor + limit) >= exchangesLength {
		last = exchangesLength
	} else {
		last = cursor + limit
	}
	return exchanges[cursor:last], nil
}

func (dummy *DummyExchangeService) Create(exchange exchange.Exchange) (uint64, error) {
	if _, isExistsWithSameId := dummy.accessor.Get(exchange.Id); isExistsWithSameId {
		errorMessage := fmt.Sprintf("Unable to create Exchange with the same id %v.\n", exchange.Id)
		return 0, errors.New(errorMessage)
	}
	return dummy.accessor.Add(exchange), nil
}

func (dummy *DummyExchangeService) Update(exchangeID uint64, exchange exchange.Exchange) error {
	if dummy.accessor.Replace(exchangeID, exchange) {
		return nil
	}
	errorMessage := fmt.Sprintf("Unable to update Exchange with id %v.\n"+
		"Reason: Exchange not found.", exchangeID)
	return errors.New(errorMessage)
}

func (dummy *DummyExchangeService) Remove(exchangeID uint64) (bool, error) {
	if dummy.accessor.Remove(exchangeID) == false {
		errorMessage := fmt.Sprintf("Unable to remove Exchange with id %v.\n"+
			"Reason: Exchange not found.", exchangeID)
		return false, errors.New(errorMessage)
	}
	return true, nil
}
