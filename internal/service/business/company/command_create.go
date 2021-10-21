package company

import (
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) Create(company business.Company) (uint64, error) {
	business.AllEntities = append(business.AllEntities, company)

	return uint64(len(business.AllEntities) - 1), nil
}
