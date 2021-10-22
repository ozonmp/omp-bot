package storage

import "github.com/ozonmp/omp-bot/internal/model/estate"

type RentEntity struct {
	ID uint64
	estate.Rent
}

func NewRentEntity(id uint64, rent estate.Rent) RentEntity {
	return RentEntity{ID: id, Rent: rent}
}

func (r RentEntity) ToModel() estate.Rent {
	rent := r.Rent
	rent.ID = r.ID
	return rent
}