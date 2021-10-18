package business

import "fmt"

type Workplace struct {
	ID uint64
	Title string
}

func (w *Workplace) String() string {
	return fmt.Sprintf("Workplace: ID - %d, Title - %s ", w.ID, w.Title)
}
