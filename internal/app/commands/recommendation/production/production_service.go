package production

import "github.com/ozonmp/omp-bot/internal/model/recommendation"

type ProductionService interface {
	List(offset int, limit int) ([]recommendation.Production, error)
	Create(production recommendation.Production) error
	Read(productionId uint64) (*recommendation.Production, error)
	Update(production recommendation.Production) error
	Delete(productionId uint64) error
}
