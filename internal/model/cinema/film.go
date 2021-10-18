package cinema

import "fmt"

type Film struct {
	ID               uint64  `json:"id"`
	Name             string  `json:"name"`
	Rating           float64 `json:"rating"`
	ShortDescription string  `json:"shortDescription"`
}

func (f *Film) String() string {
	return fmt.Sprintf("'%s' [%.1f/10]: %s", f.Name, f.Rating, f.ShortDescription)
}
