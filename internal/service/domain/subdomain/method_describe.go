package subdomain

import "github.com/ozonmp/omp-bot/internal/model/domain"

func (s *DummySubdomainService) Describe(subdomain_id uint64) (*domain.Subdomain, error) {
	if v, ok := s.mapper[subdomain_id]; !ok {
		return nil, ErrSubdomainNotFound
	} else {
		return v, nil
	}
}
