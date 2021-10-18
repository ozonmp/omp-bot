package exchange

import "github.com/ozonmp/omp-bot/internal/model/exchange"

type ExchangeService interface {
	Describe(exchangeID uint64) (*exchange.Exchange, error)
	List(cursor uint64, limit uint64) ([]exchange.Exchange, error)
	Create(exchange exchange.Exchange) (uint64, error)
	Update(exchangeID uint64, exchange exchange.Exchange) error
	Remove(exchangeID uint64) (bool, error)
}

type DummyExchangeService struct {}

func NewDummyExchangeService() *DummyExchangeService {
	return &DummyExchangeService{}
}

func (dummy DummyExchangeService) Describe(exchangeID uint64)  (*exchange.Exchange, error) {
	exchangeID = 0
	return &exchange.Exchange{}, nil
}

func (dummy DummyExchangeService) List(cursor uint64, limit uint64) ([]exchange.Exchange, error) {

	return make([]exchange.Exchange, 0), nil
}

func (dummy DummyExchangeService) Create(exchange exchange.Exchange) (uint64, error) {

	return 0, nil
}

func (dummy DummyExchangeService) Update(exchangeID uint64, exchange exchange.Exchange) error {

	return nil
}

func (dummy DummyExchangeService) Remove(exchangeID uint64) (bool, error) {

	return true, nil
}
