package autotransport

import (
	"github.com/ozonmp/omp-bot/internal/model/autotransport"
	"github.com/ozonmp/omp-bot/internal/service/autotransport/ground"
)

type GroundService interface {
	Describe(groundID uint64) (*autotransport.Ground, error)
	List(cursor uint64, limit uint64) ([]autotransport.Ground, error)
	Create(ground autotransport.Ground) (uint64, error)
	Update(groundID uint64, ground autotransport.Ground) error
	Remove(groundID uint64) (bool, error)

	Count() uint64
}

func NewAutotransportGroundService() GroundService {
	return &ground.DummyGroundService{}
}
