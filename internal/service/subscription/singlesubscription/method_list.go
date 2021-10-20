package singlesubscription

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

func (s *DummySingleSubscriptionService) List(cursor uint64, limit uint64) ([]subscription.SingleSubscription, error) {
	var from uint64
	if cursor == 0 {
		from = 0
	} else {
		from = cursor - 1
	}
	var until uint64 = from + limit

	if from >= uint64(len(s.storage)) {
		return []subscription.SingleSubscription{}, nil
	}

	res := make([]subscription.SingleSubscription, 0, limit)
	for i := from; i < until && i < uint64(len(s.storage)); i++ {
		v := s.storage[i]
		if v.IsDeleted {
			if until+1 < uint64(len(s.storage)) {
				until++
			}
			fmt.Println("inc", until)
			continue
		}
		res = append(res, *v)
	}
	return res, nil
}
