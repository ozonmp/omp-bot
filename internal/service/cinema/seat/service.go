package seat

import (
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type SeatService interface {
	Describe(seatID uint64) (*cinema.Seat, error)
	List(cursor uint64, limit uint64) ([]cinema.Seat, error)
	Create(cinema.Seat) (uint64, error)
	Update(seatID uint64, seat cinema.Seat) error
	Remove(seatID uint64) (bool, error)
}

type DummySeatService struct {
	seats  []cinema.Seat
	nextID uint64
}

func NewDummySeatService() *DummySeatService {
	return &DummySeatService{
		seats:  cinema.InitialSeats,
		nextID: uint64(len(cinema.InitialSeats) + 1),
	}
}
