package singleSubscription

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

type SingleSubscriptionService interface {
	Describe(singleSubscriptionID uint64) (*subscription.SingleSubscription, error)
	List(cursor uint64, limit uint64) ([]subscription.SingleSubscription, error)
	Create(subscription.SingleSubscription) (uint64, error)
	Update(singleSubscriptionID uint64, singleSubscription subscription.SingleSubscription) error
	Remove(singleSubscriptionID uint64) (bool, error)
}

type DummySingleSubscriptionService struct {
	storage []*subscription.SingleSubscription
	mapper  map[uint64]*subscription.SingleSubscription
	serial  uint64
}

var ErrSingleSubsriptionNotFound = errors.New("single subsription not found")
var ErrListBoundsInvalid = errors.New("list bound invalid")

func NewDummySingleSubscriptionService() *DummySingleSubscriptionService {
	s := &DummySingleSubscriptionService{
		storage: make([]*subscription.SingleSubscription, 0),
		mapper:  make(map[uint64]*subscription.SingleSubscription),
		serial:  0,
	}
	dataFill(s)
	return s
}
