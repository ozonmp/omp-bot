package warehouse

type Warehouse struct {
	ID           uint64
	OwnerID      uint64
	Address      string
	AreaM2       uint32
	PriceInCents uint64
}
