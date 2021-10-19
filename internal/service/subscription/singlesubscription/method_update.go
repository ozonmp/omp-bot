package singlesubscription

import "github.com/ozonmp/omp-bot/internal/model/subscription"

func (s *DummySingleSubscriptionService) Update(
	singleSubscriptionID uint64,
	singleSubscription subscription.SingleSubscription,
) error {
	if v, ok := s.mapper[singleSubscriptionID]; !ok {
		return ErrSingleSubsriptionNotFound
	} else {
		v.ID = singleSubscription.ID
		v.UserID = singleSubscription.UserID
		v.ServiceID = singleSubscription.ServiceID
		v.IsDeleted = singleSubscription.IsDeleted
		v.CreatedAt = singleSubscription.CreatedAt
		v.DeletedAt = singleSubscription.DeletedAt
		v.ExpireAt = singleSubscription.ExpireAt
		return nil
	}
}
