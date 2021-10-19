package delivery

import (
	"fmt"
	"time"
)

type Status string

const (
	Processed  Status = "обработан"
	Completion        = "собирается"
	InWay             = "доставляется"
	Delivered         = "доставлен"
)

type Common struct {
	id       uint64
	Order    uint64
	Status   Status
	Receiver uint64
	Courier  uint64
	Location uint64
	Time     time.Time
	Notes    string
}

func (entity Common) String() string {
	return fmt.Sprintf(
		"id: %d, Order: %d, Status: %s, Receiver: %d, Courier: %d, Location: %d, Time: %s, Notes: %s",
		entity.id,
		entity.Order,
		entity.Status,
		entity.Receiver,
		entity.Courier,
		entity.Location,
		entity.Time.Format(time.RFC3339),
		entity.Notes,
	)
}

func (entity *Common) SetId(id uint64) {
	entity.id = id
}

func (entity Common) Id() uint64 {
	return entity.id
}

var CommonStorage = []Common{
	{id: 1, Order: 9, Status: Status(InWay), Receiver: 34, Courier: 111, Location: 1, Time: time.Now(), Notes: ""},
	{id: 2, Order: 8, Status: Status(Completion), Receiver: 22, Courier: 222, Location: 6, Time: time.Now(), Notes: ""},
	{id: 3, Order: 7, Status: Status(Processed), Receiver: 78, Courier: 333, Location: 3, Time: time.Now(), Notes: ""},
	{id: 4, Order: 6, Status: Status(Delivered), Receiver: 87, Courier: 111, Location: 4, Time: time.Now(), Notes: ""},
}

var Sequence = CommonStorage[len(CommonStorage)-1].id
