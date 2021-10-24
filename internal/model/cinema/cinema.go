package cinema

import "fmt"

type Purchase struct {
	Name string
	ID   uint64
}

func (r *Purchase) String() string {
	return fmt.Sprintf("{ name[%s]. ID[%d] }", r.Name, r.ID)
}

func ShowPurchaseInputFormat() string {
	return `[format should be { "Name" : "testName" } ]`
}
