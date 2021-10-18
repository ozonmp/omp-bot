package subscription

import "fmt"

type SingleSubscription struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (s SingleSubscription) String() string {
	return fmt.Sprintf("SingleSubscription{ID: %d, Name: \"%s\"}", s.ID, s.Name)
}
