package rent

import "github.com/ozonmp/omp-bot/internal/model/cinema"

type RentService interface {
	Describe(rentID uint64) (*cinema.Rent, error)
	List(cursor uint64, limit uint64) ([]cinema.Rent, error)
	Create(cinema.Rent) (uint64, error)
	Update(rentID uint64, rent cinema.Rent) error
	Remove(rentID uint64) (bool, error)
}
