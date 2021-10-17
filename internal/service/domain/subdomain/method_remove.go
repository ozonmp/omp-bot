package subdomain

func (s *DummySubdomainService) Remove(subdomain_id uint64) (bool, error) {
	if _, ok := s.mapper[subdomain_id]; !ok {
		return false, ErrSubdomainNotFound
	} else {
		delete(s.mapper, subdomain_id)
		s.removeFromStorage(subdomain_id)
		return true, nil
	}
}

func (s *DummySubdomainService) removeFromStorage(subdomain_id uint64) error {
	idx := -1
	for i, v := range s.storage {
		if v.ID == subdomain_id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return ErrSubdomainNotFound
	}
	s.storage = append(s.storage[:idx], s.storage[idx+1:]...)
	return nil
}
