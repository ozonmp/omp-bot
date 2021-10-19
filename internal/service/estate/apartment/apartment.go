package apartment

import "github.com/ozonmp/omp-bot/internal/model/estate"

type ApartmentService interface {
	Describe(apartmentID uint64) (*estate.Apartment, error)
	List(cursor uint64, limit uint64) ([]estate.Apartment, error)
	Create(estate.Apartment) (uint64, error)
	Update(apartmentID uint64, apartment estate.Apartment) error
	Remove(apartmentID uint64) (bool, error)
}

type DummyApartmentService struct {
}

func NewDummyApartmentService() *DummyApartmentService {
	return &DummyApartmentService{}
}

func (d *DummyApartmentService) Describe(apartmentID uint64) (*estate.Apartment, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DummyApartmentService) List(cursor uint64, limit uint64) ([]estate.Apartment, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DummyApartmentService) Create(apartment estate.Apartment) (uint64, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DummyApartmentService) Update(apartmentID uint64, apartment estate.Apartment) error {
	panic("not implemented") // TODO: Implement
}

func (d *DummyApartmentService) Remove(apartmentID uint64) (bool, error) {
	panic("not implemented") // TODO: Implement
}
