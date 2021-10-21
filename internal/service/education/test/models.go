package test

import "fmt"

var allEntities = []Test{
	{id: 1, name: "one", description: "test one", min_score: 5},
	{id: 2, name: "two", description: "test two", min_score: 5},
	{id: 3, name: "three", description: "test three", min_score: 5},
	{id: 4, name: "four", description: "test four", min_score: 10},
	{id: 5, name: "five", description: "test five", min_score: 10},
}

type Test struct {
	id          int
	name        string
	description string
	min_score   int
}

func (d *Test) String() string {
	return fmt.Sprintf("Test:\n\tId:\t%d\n\tName:\t%s\n\tDescription:\t%s\n\tMin score:\t%d\n",
		d.id, d.name, d.description, d.min_score)
}
