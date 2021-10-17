package business

import "fmt"

type Office struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (o *Office) String() string {
	return fmt.Sprintf("Office name:%s Description:%s", o.Name, o.Description)
}
