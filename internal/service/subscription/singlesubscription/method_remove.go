package singlesubscription

import "time"

func (s *DummySingleSubscriptionService) Remove(singleSubscriptionID uint64) (bool, error) {
	if v, ok := s.mapper[singleSubscriptionID]; !ok {
		return false, ErrSingleSubsriptionNotFound
	} else {
		v.IsDeleted = true
		v.DeletedAt = time.Now()
		return true, nil
	}
}
