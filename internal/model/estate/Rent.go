package estate

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

type RentObjectType string

func (r RentObjectType) String() string {
	return string(r)
}

func (r RentObjectType) Icon() string {
	switch r {
	case Car:
		return "🚗"
	case House:
		return "🏚"
	default:
		panic(fmt.Errorf("RentObjectType: error unexpected value %s", r))
	}
	return r.String()
}

func (r RentObjectType) Full() string {
	switch r {
	case Car:
		return "🚗 car"
	case House:
		return "🏚 house"
	default:
		panic(fmt.Errorf("RentObjectType: error unexpected value %s", r))
	}
	return r.String()
}

const (
	House RentObjectType = "house"
	Car   RentObjectType = "car"
)

type Rent struct {
	ID         uint64
	RenterId   uint64          `json:"renter_id"`
	ObjectType RentObjectType  `json:"object_type"`
	ObjectInfo string          `json:"object_info"`
	Period     time.Duration   `json:"period"`
	Price      decimal.Decimal `json:"price"`
}

func EmptyRent(objectType RentObjectType) *Rent {
	return &Rent{
		ObjectType: objectType,
		Price: decimal.NewFromInt(0),
	}
}

func (r *Rent) SetRenterId(raw string) error {
	var err error
	if r.RenterId, err = strconv.ParseUint(raw, 10, 64); err == nil {
		return nil
	}
	return fmt.Errorf("ParseRent.RentId: expected uint64, but %v", raw)
}

func (r *Rent) PatchRenterId(raw string) error {
	raw = strings.TrimSpace(raw)
	if len(raw) == 0 {
		return nil
	}
	return r.SetRenterId(raw)
}

func (r *Rent) SetObjectType(raw string) error {
	objectType := RentObjectType(strings.TrimSpace(raw))
	switch objectType {
	case House:
	case Car:
	default:
		return fmt.Errorf("ParseRent.ObjectType: expected positive integer, but %s", raw)
	}
	r.ObjectType = objectType
	return nil
}

func (r *Rent) SetObjectInfo(raw string) error {
	raw = strings.TrimSpace(raw)
	r.ObjectInfo = strings.TrimSpace(raw)
	return nil
}

func (r *Rent) SetPeriod(raw string) error {
	raw = strings.TrimSpace(raw)
	var err error
	if r.Period, err = time.ParseDuration(raw); err == nil {
		return nil
	}
	return fmt.Errorf("ParseRent.Period: expected duration (like 1m, 1d, 1h), but %s", raw)
}

func (r *Rent) SetPrice(raw string) error {
	raw = strings.TrimSpace(raw)
	var err error
	if r.Price, err = decimal.NewFromString(raw); err == nil {
		return nil
	}
	return fmt.Errorf("ParseRent.Price: expected decimal price (dot separated), but %s", raw)
}

func (r *Rent) String() string {
	return fmt.Sprintf("%+v\n", *r)
}

func (r *Rent) ToFormatRowsString() string {
	return fmt.Sprintf(`<code>
object_type: %s
object_info: %s
renter_id:   %v
period:      %v
price:       %v
</code>`, r.ObjectType, r.ObjectInfo, r.RenterId, r.Period, r.Price)
}

func (r *Rent) ToShortRowsString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", r.ObjectType))
	sb.WriteString(fmt.Sprintf("%s\n", r.ObjectInfo))
	sb.WriteString(fmt.Sprintf("%v\n", r.RenterId))
	sb.WriteString(fmt.Sprintf("%v\n", r.Period))
	sb.WriteString(fmt.Sprintf("%v\n", r.Price))
	return sb.String()
}