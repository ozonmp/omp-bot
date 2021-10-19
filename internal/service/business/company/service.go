package company

import (
	"github.com/ozonmp/omp-bot/internal/model/business"
)

type CompanyService interface {
	Describe(companyID uint64) (*business.Company, error)
	List(cursor uint64, limit uint64) ([]business.Company, error)
	Create(business.Company) (uint64, error)
	Update(companyID uint64, company business.Company) error
	Remove(companyID uint64) (bool, error)
}

type DummyCompanyService struct{}

func NewDummyCompanyService() *DummyCompanyService {
	return &DummyCompanyService{}
}
