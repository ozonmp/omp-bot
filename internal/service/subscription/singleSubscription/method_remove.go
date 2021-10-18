package singleSubscription

func (s *DummySingleSubscriptionService) Remove(singleSubscriptionID uint64) (bool, error) {
	if _, ok := s.mapper[singleSubscriptionID]; !ok {
		return false, ErrSingleSubsriptionNotFound
	} else {
		delete(s.mapper, singleSubscriptionID)
		s.removeFromStorage(singleSubscriptionID)
		return true, nil
	}
}

func (s *DummySingleSubscriptionService) removeFromStorage(singleSubscriptionID uint64) error {
	idx := -1
	for i, v := range s.storage {
		if v.ID == singleSubscriptionID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return ErrSingleSubsriptionNotFound
	}
	s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
	return nil
}
