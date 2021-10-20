package exchange

import "fmt"

type Return1 struct {
	Name string
	ID   uint64
}

func (r *Return1) String() string {
	return fmt.Sprintf("{ name[%s]. ID[%d] }", r.Name, r.ID)
}
