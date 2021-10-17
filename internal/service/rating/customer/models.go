package customer

import "fmt"

var allEntities = []Customer{
	{FirstName: "one", SecondName: "one"},
	{FirstName: "two", SecondName: "two"},
	{FirstName: "three", SecondName: "three"},
	{FirstName: "four", SecondName: "four"},
	{FirstName: "five", SecondName: "five"},
}

type Customer struct {
	FirstName  string
	SecondName string
}

func (c Customer) String() string {
	return fmt.Sprintf("Firstname: %s SecondName: %s", c.FirstName, c.SecondName)
}

func NewCustomer(firstName string, secondName string) Customer {
	return Customer{
		FirstName:  firstName,
		SecondName: secondName,
	}
}
