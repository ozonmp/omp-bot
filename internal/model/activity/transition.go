package activity

import "fmt"

type Transition struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}

func (t Transition) String() string {
	return fmt.Sprintf("id: %v\nname: %v\nfrom: %v\nto: %v\n", t.Id, t.Name, t.From, t.To)
}
