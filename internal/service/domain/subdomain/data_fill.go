package subdomain

import "github.com/ozonmp/omp-bot/internal/model/domain"

var (
	sampleData = []domain.Subdomain{
		{Name: "Avengers"},
		{Name: "Spider-Man"},
		{Name: "Iron Man"},
		{Name: "Black Panther"},
		{Name: "Deadpool"},
		{Name: "Captain America"},
		{Name: "Jessica Jones"},
		{Name: "Ant-Man"},
		{Name: "Captain Marvel"},
		{Name: "Guardians of the Galaxy"},
		{Name: "Wolverine"},
		{Name: "Luke Cage"},
	}
)

func dataFill(s SubdomainService) {
	for _, v := range sampleData {
		s.Create(v)
	}
}
