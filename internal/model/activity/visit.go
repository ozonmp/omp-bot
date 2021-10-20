package activity

import (
	"errors"
	"fmt"
)

type Visit struct {
	Id    uint64
	Title string
}

func (v *Visit) String() string {
	return fmt.Sprintf("Visit - Id: %d, Title: %s.", v.Id, v.Title)
}

type VisitModel []Visit

func (v *VisitModel) FindById(visitId uint64) (*Visit, error) {
	for _, v := range *v {
		if v.Id == visitId {
			return &v, nil
		}
	}

	return &Visit{}, errors.New("visit not found")
}

func (v *VisitModel) FindIndexById(visitId uint64) (int, error) {
	for i, v := range *v {
		if v.Id == visitId {
			return i, nil
		}
	}

	return 0, errors.New("visit index not found")
}

var Visits = VisitModel{
	Visit{Id: 1, Title: "Some title for first visit"},
	Visit{Id: 2, Title: "Some title for second visit"},
	Visit{Id: 3, Title: "Some title for third visit"},
	Visit{Id: 4, Title: "Some title for fourth visit"},
	Visit{Id: 5, Title: "Some title for fifth visit"},
	Visit{Id: 6, Title: "Some title for sixth visit"},
	Visit{Id: 7, Title: "Some title for seventh visit"},
	Visit{Id: 8, Title: "Some title for eighth visit"},
	Visit{Id: 9, Title: "Some title for ninth visit"},
	Visit{Id: 10, Title: "Some title for tenth visit"},
}
