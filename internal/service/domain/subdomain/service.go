package subdomain

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/domain"
)

type SubdomainService interface {
	Describe(subdomain_id uint64) (*domain.Subdomain, error)
	List(cursor uint64, limit uint64) ([]domain.Subdomain, error)
	Create(domain.Subdomain) (uint64, error)
	Update(subdomain_id uint64, subdomain domain.Subdomain) error
	Remove(subdomain_id uint64) (bool, error)
}

type DummySubdomainService struct {
	storage []*domain.Subdomain
	mapper  map[uint64]*domain.Subdomain
	serial  uint64
}

var ErrSubdomainNotFound = errors.New("subdomain not found")
var ErrListBoundsInvalid = errors.New("list bound invalid")

func NewDummySubdomainService() *DummySubdomainService {
	s := &DummySubdomainService{
		storage: make([]*domain.Subdomain, 0),
		mapper:  make(map[uint64]*domain.Subdomain),
		serial:  0,
	}
	dataFill(s)
	return s
}
