package recommendation

import "fmt"

type Production struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (p Production) String() string {
	return fmt.Sprintf("%v %v %v", p.Id, p.Description, p.Type)
}
