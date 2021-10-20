package warehouse

import "fmt"

type Warehouse struct {
	ID           uint64 `json: "id"`
	OwnerID      uint64 `json: "OwnerID"`
	Address      string `json: "Address"`
	AreaM2       uint32 `json: "AreaM2"`
	PriceInCents uint64 `json: "PriceInCents"`
}

var allWarehouses = []Warehouse{
	{ID: 1, OwnerID: 1, Address: "Location 1", AreaM2: 1, PriceInCents: 100},
	{ID: 2, OwnerID: 1, Address: "Location 2", AreaM2: 1, PriceInCents: 50000},
	{ID: 3, OwnerID: 2, Address: "Location 3", AreaM2: 2, PriceInCents: 1},
	{ID: 4, OwnerID: 2, Address: "Location 4", AreaM2: 2, PriceInCents: 1000},
	{ID: 5, OwnerID: 3, Address: "Location 5", AreaM2: 4, PriceInCents: 100000},
	{ID: 6, OwnerID: 3, Address: "Location 6", AreaM2: 3, PriceInCents: 10},
	{ID: 7, OwnerID: 3, Address: "Location 7", AreaM2: 8, PriceInCents: 10000},
}

func (t *Warehouse) String() string {
	return fmt.Sprintf(
		` id: '%d'\n OwnerID: '%d'\n Address: '%v'\n AreaM2: '%d'\n PriceInCents: '%d'`,
		t.ID,
		t.OwnerID,
		t.Address,
		t.AreaM2,
		t.PriceInCents,
	)
}
