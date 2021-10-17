package subdomain

import "github.com/ozonmp/omp-bot/internal/model/domain"

func (s *DummySubdomainService) Update(subdomain_id uint64, subdomain domain.Subdomain) error {
	if v, ok := s.mapper[subdomain_id]; !ok {
		return ErrSubdomainNotFound
	} else {
		v.ID = subdomain.ID
		v.Name = subdomain.Name
		return nil
	}
}
