package subdomain

import "github.com/ozonmp/omp-bot/internal/model/domain"

func (s *DummySubdomainService) Create(subdomain domain.Subdomain) (uint64, error) {
	s.serial++
	subdomain.ID = s.serial
	s.storage = append(s.storage, &subdomain)
	s.mapper[subdomain.ID] = &subdomain
	return subdomain.ID, nil
}
