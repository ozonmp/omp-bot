package seat

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

func (s *DummySeatService) Describe(seatID uint64) (*cinema.Seat, error) {
	for i := range s.seats {
		if s.seats[i].ID == seatID {
			return &s.seats[i], nil
		}
	}

	return nil, &SeatNotFoundError{seatID}
}

func (s *DummySeatService) List(cursor uint64, limit uint64) ([]cinema.Seat, error) {
	length := uint64(len(s.seats))

	if length == 0 {
		return nil, ErrEmptyList
	}

	if cursor >= length {
		return nil, fmt.Errorf("cursor %d is out of range [0, %d]", cursor, length-1)
	}

	var err error = nil
	end := cursor + limit
	if end >= length || limit == 0 {
		end = length
		err = ErrEndOfList
	}

	return s.seats[cursor:end], err
}

func (s *DummySeatService) Create(seat cinema.Seat) (uint64, error) {
	seat.ID = s.nextID
	s.nextID++
	s.seats = append(s.seats, seat)

	return seat.ID, nil
}

func (s *DummySeatService) Update(seatID uint64, seat cinema.Seat) error {
	for i := range s.seats {
		if s.seats[i].ID == seatID {
			s.seats[i] = seat
			s.seats[i].ID = seatID
			return nil
		}
	}

	return &SeatNotFoundError{seatID}
}

func (s *DummySeatService) Remove(seatID uint64) (bool, error) {
	for i := range s.seats {
		if s.seats[i].ID == seatID {
			s.seats = append(s.seats[:i], s.seats[i+1:]...)
			return true, nil
		}
	}

	return false, &SeatNotFoundError{seatID}
}

var (
	ErrEmptyList = errors.New("empty seat list")
	ErrEndOfList = errors.New("no more seat in list")
)

type SeatNotFoundError struct {
	ID uint64
}

func (e *SeatNotFoundError) Error() string {
	return fmt.Sprintf("seat with ID %d not found", e.ID)
}
