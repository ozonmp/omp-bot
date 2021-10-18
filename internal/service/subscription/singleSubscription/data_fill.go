package singleSubscription

import "github.com/ozonmp/omp-bot/internal/model/subscription"

var (
	sampleData = []subscription.SingleSubscription{
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

func dataFill(s SingleSubscriptionService) {
	for _, v := range sampleData {
		s.Create(v)
	}
}
