package singlesubscription

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

func (s *DummySingleSubscriptionService) List(cursor uint64, limit uint64) ([]subscription.SingleSubscription, error) {
	if cursor >= uint64(len(s.storage)) {
		fmt.Println(1)
		return []subscription.SingleSubscription{}, nil
	}
	var from uint64
	var until uint64 = cursor + limit - 1
	if cursor == 0 {
		from = 0
	} else {
		from = cursor - 1
	}

	if until > uint64(len(s.storage)) {
		until = uint64(len(s.storage))
	}
	if from > until {
		fmt.Println(2)
		return []subscription.SingleSubscription{}, ErrListBoundsInvalid
	}

	res := make([]subscription.SingleSubscription, 0, limit)
	for i := from; i < until && i < uint64(len(s.storage)); i++ {
		v := s.storage[i]
		if v.IsDeleted {
			if until+1 < uint64(len(s.storage)) {
				until++
			}
			continue
		}
		res = append(res, *v)
	}
	return res, nil
}
