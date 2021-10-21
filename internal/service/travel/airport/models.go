package airport

import "fmt"

type Airport struct {
	name     string
	location string
}

var allAirports = []Airport{
	{"JFK", "New York"},
	{"LAX", "Los Angeles"},
	{"OVG", "Novosibirsk"},
	{"DMD", "Moscow"},
	{"PLK", "Saint-Petersburg"},
	{"HLW", "Hello World"},
	{"ADA", "dummy location"},
}

func (a Airport) String() string {
	return fmt.Sprintf("Airport %s [%s]", a.name, a.location)
}
