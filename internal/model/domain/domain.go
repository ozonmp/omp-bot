package domain

import "fmt"

type Subdomain struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (s Subdomain) String() string {
	return fmt.Sprintf("Subdomain{ID: %d, Name: \"%s\"}", s.ID, s.Name)
}
