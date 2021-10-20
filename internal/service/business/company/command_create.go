package company

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *DummyCompanyService) Create(business.Company) (uint64, error) {
	return 0, fmt.Errorf("not implemented")
}
