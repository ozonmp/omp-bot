package assets

import (
	"fmt"
	"math/rand"
	"time"
)

var AllEntities []Asset

var PageSize float64 = 10

type Asset struct {
	ID               uint64  `json:"id"`
	CreatedAt        time.Time   `json:"created_at"`
	User             uint64  `json:"user"`
	PriceWhenCreated float64 `json:"price_when_created"`
	CurrentPrice     float64 `json:"current_price"`
}

func (a *Asset) String() string {
	return fmt.Sprintf(
		"ID: %d; ID пользователя: %d; Создано: %s; Начальная цена: %.2f руб.; Текущая цена: %.2f руб.",
		a.ID,
		a.User,
		fmt.Sprintf(a.CreatedAt.Format("2006-01-02 15:04:05")),
		a.PriceWhenCreated,
		a.CurrentPrice,
	)
}

func init() {
	for i := uint64(0); i < uint64(25); i++ {
		rnd := rand.Float64()*float64(i)
		AllEntities = append(
			AllEntities,
			Asset{
				ID:               i,
				CreatedAt:        time.Now(),
				User:             1000 + i,
				PriceWhenCreated: rnd,
				CurrentPrice:     rnd,
			},
		)
	}
}