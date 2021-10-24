package estate

import (
	"fmt"
	"strings"
)

type Apartment struct {
	ID    uint64
	Title string
	Price int64
}

func (a Apartment) String() string {
	return fmt.Sprintf("(%d) %s. Price: %d", a.ID, a.Title, a.Price)
}

func StringFromApartments(apartments []Apartment) string {
	if len(apartments) == 0 {
		return ""
	}
	builder := strings.Builder{}

	// ignore errors, since WriteByte and WriteString always return nil error
	builder.WriteString(apartments[0].String())
	for i := 1; i < len(apartments); i++ {
		builder.WriteByte('\n')
		builder.WriteString(apartments[i].String())
	}
	return builder.String()
}
