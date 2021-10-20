package apartment

import "github.com/ozonmp/omp-bot/internal/model/estate"

type ApartmentService interface {
	Describe(apartmentID uint64) (*estate.Apartment, error)
	List(cursor uint64, limit uint64) ([]estate.Apartment, error)
	Create(estate.Apartment) (uint64, error)
	Update(apartmentID uint64, apartment estate.Apartment) error
	Remove(apartmentID uint64) (bool, error)
}

type ApartmentStorage interface {
	Describe(apartmentID uint64) (*estate.Apartment, error)
	List(cursor uint64, limit uint64) ([]estate.Apartment, error)
	Create(estate.Apartment) (uint64, error)
	Update(apartmentID uint64, apartment estate.Apartment) error
	Remove(apartmentID uint64) (bool, error)
}

type DummyApartmentService struct {
	storage ApartmentStorage
}

func NewDummyApartmentService() *DummyApartmentService {
	startingItems := []estate.Apartment{
		{Title: "My awesome apartment", Price: 132},
		{Title: "Not so awesome apartment", Price: 93},
		{Title: "Not my apartment", Price: 208},
		{Title: "Not my but awesome apartment", Price: 248},
		{Title: "Ruined apartment", Price: 18},
		{Title: "My friend's apartment", Price: 180},
		{Title: "My friend's other apartment", Price: 182},
		{Title: "No one's apartment", Price: 308},
		{Title: "Apartment on sale", Price: 112},
	}
	return &DummyApartmentService{
		storage: estate.NewInMemoryApartmentStorage(startingItems),
	}
}

func (d *DummyApartmentService) Describe(apartmentID uint64) (*estate.Apartment, error) {
	return d.storage.Describe(apartmentID)
}

func (d *DummyApartmentService) List(cursor uint64, limit uint64) ([]estate.Apartment, error) {
	return d.storage.List(cursor, limit)
}

func (d *DummyApartmentService) Create(apartment estate.Apartment) (uint64, error) {
	return d.storage.Create(apartment)
}

func (d *DummyApartmentService) Update(apartmentID uint64, apartment estate.Apartment) error {
	return d.storage.Update(apartmentID, apartment)
}

func (d *DummyApartmentService) Remove(apartmentID uint64) (bool, error) {
	return d.storage.Remove(apartmentID)
}
