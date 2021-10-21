package company

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) Remove(companyID uint64) (bool, error) {
	if companyID < 0 || int(companyID) >= len(business.AllEntities) {
		return false, fmt.Errorf("Invalid index %d, max elements of companies - %d.", companyID, len(business.AllEntities)-1)
	}

	business.AllEntities = append(business.AllEntities[:companyID], business.AllEntities[companyID+1:]...)

	return true, nil
}
