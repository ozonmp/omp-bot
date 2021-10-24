package purchase

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type DummyPurchaseService struct {
	savedPurchases []cinema.Purchase
}

func NewDummyPurchaseService() *DummyPurchaseService {
	s := DummyPurchaseService{}
	s.Create(cinema.Purchase{Name: "zero"})
	s.Create(cinema.Purchase{Name: "one"})
	s.Create(cinema.Purchase{Name: "two"})
	s.Create(cinema.Purchase{Name: "three"})
	s.Create(cinema.Purchase{Name: "four"})
	s.Create(cinema.Purchase{Name: "five"})
	s.Create(cinema.Purchase{Name: "six"})

	return &s
}

var wrongIDErr = errors.New("wrong id")

func (r *DummyPurchaseService) Describe(purchaseID uint64) (*cinema.Purchase, error) {
	if uint64(len(r.savedPurchases)) <= purchaseID {
		return nil, wrongIDErr
	}

	return &r.savedPurchases[purchaseID], nil
}

var LastPageExceededErr = errors.New("last page exceeded")

func (r *DummyPurchaseService) List(cursor uint64, limit uint64) ([]cinema.Purchase, error) {
	if len(r.savedPurchases) == 0 {
		return nil, nil
	}

	var low, high uint64
	low = cursor
	high = cursor + limit

	if high > uint64(len(r.savedPurchases)) {
		high = uint64(len(r.savedPurchases))
	}

	if low >= uint64(len(r.savedPurchases)) {
		return nil, LastPageExceededErr
	}

	return r.savedPurchases[low:high], nil
}

func (r *DummyPurchaseService) Create(p cinema.Purchase) (uint64, error) {
	p.ID = uint64(len(r.savedPurchases))
	r.savedPurchases = append(r.savedPurchases, p)

	return p.ID, nil
}

func (r *DummyPurchaseService) Update(purchaseID uint64, purchase cinema.Purchase) error {
	if uint64(len(r.savedPurchases)) <= purchaseID {
		return wrongIDErr
	}

	r.savedPurchases[purchaseID] = purchase
	r.savedPurchases[purchaseID].ID = purchaseID

	return nil
}

func (r *DummyPurchaseService) Remove(purchaseID uint64) (bool, error) {
	if uint64(len(r.savedPurchases)) <= purchaseID {
		return false, wrongIDErr
	}

	r.savedPurchases = append(r.savedPurchases[:purchaseID], r.savedPurchases[purchaseID+1:]...)

	for i := 0; i < len(r.savedPurchases); i++ {
		r.savedPurchases[i].ID = uint64(i)
	}

	return true, nil
}
