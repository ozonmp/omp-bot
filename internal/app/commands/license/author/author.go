package author

import "fmt"

type Author struct {
	id        uint64 `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}

func (p Author) String() string {
	return fmt.Sprintf("%v %v %v", p.id, p.firstName, p.lastName)
}

var tempAuthors = []Author{
	{id: 1, firstName: "Anastasia", lastName: "Ivanova"},
	{id: 2, firstName: "Anastasia", lastName: "Petrova"},
	{id: 3, firstName: "Eugenia", lastName: "Smirnova"},
}
