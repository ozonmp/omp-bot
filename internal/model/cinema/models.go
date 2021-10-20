package cinema

import "fmt"

type Seat struct {
	ID     uint64
	Row    uint64
	Number uint64
	Price  uint64
}

func (r *Seat) String() string {
	return fmt.Sprintf(`Seat #%d (Row: %d; Number: %d) price: %d`, r.ID, r.Row, r.Number, r.Price)
}

var InitialSeats = []Seat{
	{
		ID:     1,
		Row:    1,
		Number: 1,
		Price:  100,
	}, {
		ID:     1,
		Row:    1,
		Number: 2,
		Price:  100,
	}, {
		ID:     2,
		Row:    2,
		Number: 1,
		Price:  150,
	}, {
		ID:     3,
		Row:    2,
		Number: 2,
		Price:  150,
	}, {
		ID:     4,
		Row:    3,
		Number: 3,
		Price:  200,
	},
}
