package company

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) List(cursor uint64, limit uint64) ([]business.Company, error) {
	if cursor < 0 || int(cursor) >= len(business.AllEntities) {
		return nil, fmt.Errorf("Invalid index %d, max elements of companies - %d.", cursor, len(business.AllEntities)-1)
	}

	max := int(cursor + limit)
	if max > len(business.AllEntities) {
		max = len(business.AllEntities)
	}

	return business.AllEntities[cursor:max], nil
}
