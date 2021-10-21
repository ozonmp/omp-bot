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
	Visit{Id: 1, Title: "Some title for #1 visit"},
	Visit{Id: 2, Title: "Some title for #2 visit"},
	Visit{Id: 3, Title: "Some title for #3 visit"},
	Visit{Id: 4, Title: "Some title for #4 visit"},
	Visit{Id: 5, Title: "Some title for #5 visit"},
	Visit{Id: 6, Title: "Some title for #6 visit"},
	Visit{Id: 7, Title: "Some title for #7 visit"},
	Visit{Id: 8, Title: "Some title for #8 visit"},
	Visit{Id: 9, Title: "Some title for #9 visit"},
	Visit{Id: 10, Title: "Some title for #10 visit"},
	Visit{Id: 11, Title: "Some title for #11 visit"},
	Visit{Id: 12, Title: "Some title for #12 visit"},
	Visit{Id: 13, Title: "Some title for #13 visit"},
	Visit{Id: 14, Title: "Some title for #14 visit"},
	Visit{Id: 15, Title: "Some title for #15 visit"},
	Visit{Id: 16, Title: "Some title for #16 visit"},
	Visit{Id: 17, Title: "Some title for #17 visit"},
	Visit{Id: 18, Title: "Some title for #18 visit"},
	Visit{Id: 19, Title: "Some title for #19 visit"},
	Visit{Id: 20, Title: "Some title for #20 visit"},
}
