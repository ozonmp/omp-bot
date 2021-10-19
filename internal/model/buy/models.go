package buy

import "fmt"

type Order struct {
	Id       uint64
	Title    string
	Quantity uint64
}

func (o Order) String() string {
	return fmt.Sprintf(`Order{id: %v, title: "%v", quantity: %v}`, o.Id, o.Title, o.Quantity)
}
