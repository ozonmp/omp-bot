package production

import (
	"errors"
	"fmt"
	"sort"
	"sync/atomic"

	"github.com/ozonmp/omp-bot/internal/model/recommendation"
)

var ErrProductionNotFound = errors.New("production not found")

type RecommendationProductionService struct {
	productions []recommendation.Production
	sequence    uint64
}

func NewRecommendationProductionService() *RecommendationProductionService {
	return &RecommendationProductionService{}
}

func (d *RecommendationProductionService) List(offset int, limit int) ([]recommendation.Production, error) {
	start, end := offset, offset
	ln := len(d.productions)

	if start > ln {
		return nil, nil
	}

	if start+limit > ln {
		end = ln
	} else {
		end = start + limit
	}

	return d.productions[start:end], nil
}

func (d *RecommendationProductionService) Create(production recommendation.Production) error {
	production.Id = atomic.AddUint64(&d.sequence, 1)

	d.productions = append(d.productions, production)

	return nil
}

func (d *RecommendationProductionService) Read(productionId uint64) (*recommendation.Production, error) {
	index, err := d.find(productionId)
	if err != nil {
		return nil, err
	}

	return &d.productions[index], nil
}

func (d *RecommendationProductionService) Update(production recommendation.Production) error {
	index, err := d.find(production.Id)
	if err != nil {
		return err
	}

	d.productions[index].Title = production.Title
	d.productions[index].Description = production.Description
	d.productions[index].Rating = production.Rating

	return nil
}

func (d *RecommendationProductionService) Delete(productionId uint64) error {
	index, err := d.find(productionId)
	if err != nil {
		return err
	}

	d.productions = append(d.productions[:index], d.productions[index+1:]...)

	return nil
}

func (d *RecommendationProductionService) find(productionId uint64) (int, error) {
	ln := len(d.productions)

	index := sort.Search(ln, func(i int) bool {
		return d.productions[i].Id >= productionId
	})

	if index < ln && d.productions[index].Id == productionId {
		return index, nil
	} else {
		return index, fmt.Errorf("%w id %v", ErrProductionNotFound, productionId)
	}
}
