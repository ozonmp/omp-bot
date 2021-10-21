package order

import (
	"fmt"
	"sort"
	"time"

	"github.com/ozonmp/omp-bot/internal/model/buy"
)

type DummyOrderService struct {
	curId      uint64
	orders     map[uint64]buy.Order
	listCached bool
	cache      []buy.Order
}

func NewDummyOrderService() *DummyOrderService {
	s := &DummyOrderService{}
	s.orders = make(map[uint64]buy.Order)
	return s
}

func (s *DummyOrderService) Describe(orderID uint64) (*buy.Order, error) {
	order, ok := s.orders[orderID]
	if !ok {
		return &buy.Order{}, fmt.Errorf("no order with id %v found", orderID)
	}

	return &order, nil
}

func (s *DummyOrderService) List(cursor uint64, limit uint64) ([]buy.Order, error) {
	if cursor > uint64(len(s.orders)) {
		return nil, fmt.Errorf("index %v is out of range", cursor)
	}

	cache := s.GetCache()
	l := min(limit+cursor, uint64(len(cache)))
	return cache[cursor:l], nil
}

func (s *DummyOrderService) Create(order buy.Order) (uint64, error) {
	id := s.curId
	s.curId++
	s.orders[id] = buy.Order{
		Id:        id,
		UserId:    order.UserId,
		Date:      time.Now(),
		AddressId: order.AddressId,
		StateId:   order.StateId,
		Paid:      order.Paid,
	}

	s.listCached = false
	return id, nil
}

func (s *DummyOrderService) Update(orderID uint64, order buy.Order) error {
	_, ok := s.orders[orderID]
	if !ok {
		return fmt.Errorf("no order with id %v found", orderID)
	}

	s.orders[orderID] = buy.Order{
		Id:        orderID,
		UserId:    order.UserId,
		Date:      time.Now(),
		AddressId: order.AddressId,
		StateId:   order.StateId,
		Paid:      order.Paid,
	}

	s.listCached = false
	return nil
}

func (s *DummyOrderService) Remove(orderID uint64) (bool, error) {
	_, ok := s.orders[orderID]
	if !ok {
		return false, fmt.Errorf("no order with id %v found", orderID)
	}

	delete(s.orders, orderID)

	s.listCached = false
	return true, nil
}

func (s *DummyOrderService) GetCache() []buy.Order {
	if s.listCached {
		return s.cache
	}

	len := len(s.orders)

	keys := make([]uint64, 0, len)
	for k := range s.orders {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	s.cache = make([]buy.Order, 0, len)
	for _, k := range keys {
		s.cache = append(s.cache, s.orders[k])
	}

	s.listCached = true
	return s.cache
}

func min(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
