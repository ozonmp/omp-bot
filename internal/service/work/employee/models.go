package employee

var allEntities = []Employee{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Employee struct {
	Title string
}

func (employee Employee) String() string {
	return employee.Title
}
