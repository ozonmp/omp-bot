package buy

import (
	"fmt"
	"time"
)

type Order struct {
	Id        uint64
	UserId    uint64
	Date      time.Time
	AddressId uint64
	StateId   uint32
	Paid      bool
}

func (o Order) String() string {
	return fmt.Sprintf(`Order{id: %v, user_id: %v, date: %v, address_id: %v, state_id: %v, paid: %v}`,
		o.Id, o.UserId, o.Date.Format(time.RFC822), o.AddressId, o.StateId, o.Paid)
}
