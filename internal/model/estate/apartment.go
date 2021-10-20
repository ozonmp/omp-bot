package estate

import "fmt"

type Apartment struct {
	ID    uint64
	Title string
	Price int64
}

func (a Apartment) String() string {
	return fmt.Sprintf("(%d) %s. Price: %d", a.ID, a.Title, a.Price)
}
