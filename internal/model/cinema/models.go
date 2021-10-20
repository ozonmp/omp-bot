package cinema

import "fmt"

type Seat struct {
	Row    uint64
	Number uint64
	Price  uint64
}

func (r *Seat) String() string {
	return fmt.Sprintf(`Seat (Row: %d; Number: %d) price: %d`, r.Row, r.Number, r.Price)
}
