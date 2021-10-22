package rent

import (
	"github.com/ozonmp/omp-bot/internal/model/estate"
)

type RentService interface {
	Describe(rentID uint64) (*estate.Rent, error)
	List(cursor uint64, limit uint64) ([]estate.Rent, error)
	Create(estate.Rent) (uint64, error)
	Update(rentID uint64, rent estate.Rent) error
	Remove(rentID uint64) (bool, error)
}


