package singleSubscription

import "github.com/ozonmp/omp-bot/internal/model/subscription"

func (s *DummySingleSubscriptionService) Describe(singleSubscriptionID uint64) (*subscription.SingleSubscription, error) {
	if v, ok := s.mapper[singleSubscriptionID]; !ok {
		return nil, ErrSingleSubsriptionNotFound
	} else {
		return v, nil
	}
}
