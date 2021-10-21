package company

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) Describe(companyID uint64) (*business.Company, error) {
	if companyID < 0 || int(companyID) >= len(business.AllEntities) {
		return nil, fmt.Errorf("Invalid index %d, max elements of companies - %d.", companyID, len(business.AllEntities)-1)
	}

	return &business.AllEntities[companyID], nil
}
