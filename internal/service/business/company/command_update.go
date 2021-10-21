package company

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) Update(companyID uint64, company business.Company) error {
	if companyID < 0 || int(companyID) >= len(business.AllEntities) {
		return fmt.Errorf("Invalid index %d, max elements of companies - %d.", companyID, len(business.AllEntities)-1)
	}

	business.AllEntities[companyID] = company

	return nil
}
