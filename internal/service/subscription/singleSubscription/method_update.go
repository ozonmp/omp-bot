package singleSubscription

import "github.com/ozonmp/omp-bot/internal/model/subscription"

func (s *DummySingleSubscriptionService) Update(
	singleSubscriptionID uint64,
	singleSubscription subscription.SingleSubscription,
) error {
	if v, ok := s.mapper[singleSubscriptionID]; !ok {
		return ErrSingleSubsriptionNotFound
	} else {
		v.ID = singleSubscription.ID
		v.Name = singleSubscription.Name
		return nil
	}
}
