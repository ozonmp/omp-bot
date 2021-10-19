package singlesubscription

import (
	"time"

	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

var (
	sampleData = []subscription.SingleSubscription{
		{
			UserID:    1,
			ServiceID: 1,
			IsDeleted: false,
			CreatedAt: time.Now(),
		},
		{
			UserID:    2,
			ServiceID: 2,
			IsDeleted: false,
			CreatedAt: time.Now(),
		},
		{
			// Will expire in 6 months
			UserID:    3,
			ServiceID: 1,
			IsDeleted: false,
			ExpireAt:  time.Now().AddDate(0, 6, 0),
			CreatedAt: time.Now(),
		},
	}
)

func dataFill(s SingleSubscriptionService) {
	for _, v := range sampleData {
		s.Create(v)
	}
}
