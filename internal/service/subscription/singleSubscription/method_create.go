package singleSubscription

import "github.com/ozonmp/omp-bot/internal/model/subscription"

func (s *DummySingleSubscriptionService) Create(singleSubscription subscription.SingleSubscription) (uint64, error) {
	s.serial++
	singleSubscription.ID = s.serial
	s.storage = append(s.storage, &singleSubscription)
	s.mapper[singleSubscription.ID] = &singleSubscription
	return singleSubscription.ID, nil
}
