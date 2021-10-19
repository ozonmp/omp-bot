package employee

import "strconv"

type Employee struct {
	Id    int
	Title string
}

func (employee Employee) String() string {
	return employee.Title
}

func (employee Employee) idAsString() string {
	return strconv.Itoa(employee.Id)
}
