package assets

import "fmt"

var allEntities = map[uint64]Asset{}

type Asset struct {
	ID uint64 `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

func (a *Asset) String() string {
	return fmt.Sprintf("ID: %d; Имя: %s; Баланс: %.2f руб.", a.ID, a.Username, a.Balance)
}

