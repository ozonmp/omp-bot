package employee

type Employee struct {
	Id    int
	Title string
}

func (employee Employee) String() string {
	return employee.Title
}
