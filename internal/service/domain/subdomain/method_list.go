package subdomain

import (
	"github.com/ozonmp/omp-bot/internal/model/domain"
)

func (s *DummySubdomainService) List(cursor uint64, limit uint64) ([]domain.Subdomain, error) {
	if cursor >= uint64(len(s.storage)) {
		return []domain.Subdomain{}, nil
	}
	var from uint64
	var until uint64 = cursor + limit - 1
	if cursor == 0 {
		from = 0
	} else {
		from = cursor - 1
	}

	if until > uint64(len(s.storage)) {
		until = uint64(len(s.storage))
	}
	if from > until {
		return []domain.Subdomain{}, ErrListBoundsInvalid
	}

	res := make([]domain.Subdomain, 0, limit)
	for _, v := range s.storage[from:until] {
		res = append(res, *v)
	}
	return res, nil
}
