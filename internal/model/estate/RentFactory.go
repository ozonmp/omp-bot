package estate

import (
	"fmt"
	"strings"
)

var orderSet = [5]func(*Rent, string) error{
	(*Rent).SetObjectType,
	(*Rent).SetObjectInfo,
	(*Rent).SetRenterId,
	(*Rent).SetPeriod,
	(*Rent).SetPrice,
}

func ParseRent(inline string) (*Rent, error) {
	inline = strings.TrimSpace(inline)
	parts := strings.Split(inline, "\n")

	rent := Rent{}
	if err := rent.Init(parts); err != nil {
		return nil, err
	}

	return &rent, nil
}

func (r *Rent) Init(parts []string) error {
	expected := 5
	partsCount := len(parts)
	if partsCount != expected {
		return fmt.Errorf("Rent.Init: error not enought fields, expected: %v, but %v", expected, parts)
	}

	for i, part := range parts {
		if err := orderSet[i](r, part); err != nil {
			return err
		}
	}

	return nil
}

func (r *Rent) Patch(parts []string) error {
	patchIfNotEmpty := func(raw string, set func(r *Rent, raw string) error) error {
		raw = strings.TrimSpace(raw)
		if len(raw) > 0 {
			return set(r, raw)
		}
		return nil
	}

	expected := 1
	partsCount := len(parts)
	if partsCount <= expected {
		return fmt.Errorf("Rent.Patch: error not enought fields, expected: %v, but %v", expected, parts)
	}

	for i, part := range parts {
		if err := patchIfNotEmpty(part, orderSet[i]); err != nil {
			return err
		}
	}

	return nil
}
