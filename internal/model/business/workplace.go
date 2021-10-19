package business

import "fmt"

type Workplace struct {
	ID              uint64
	Title           string
	EmployeeID      uint64
	WorkplaceNumber uint32
	OfficeID        uint64
}

func (w *Workplace) String() string {
	return fmt.Sprintf("Workplace: ID - %d, Title - %s, EmployeeID - %d, WorkplaceNumber - %d, OfficeID - %d ",
		w.ID,
		w.Title,
		w.EmployeeID,
		w.WorkplaceNumber,
		w.OfficeID)
}
