package customer

import "github.com/ozonmp/omp-bot/internal/service/rating/customer"

type DummyService interface {
	List(cursor uint64, limit uint64) ([]customer.Customer, error)
	Count() int
	Describe(idx uint64) (*customer.Customer, error)
	Remove(idx uint64) (bool, error)
	Create(customer customer.Customer) (uint64, error)
	Update(idx uint64, customer customer.Customer) error
}
