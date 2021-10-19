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

//

// func (s *DummyCompanyService) List() []Company {
// 	return allEntities
// }

// func (s *DummyCompanyService) Get(idx int) (*Company, error) {
// 	return &allEntities[idx], nil
// }

// func (s *DummyCompanyService) Delete(idx int) (bool, error) {
// 	if idx < 0 || idx > len(allEntities) {
// 		return false, fmt.Errorf("Invalid index %d, max elements of companies - %d.", idx, len(allEntities))
// 	}

// 	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)

// 	return true, nil
// }
