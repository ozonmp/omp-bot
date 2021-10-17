package business

import "fmt"

type Office struct {
	Id          uint64
	Name        string
	Description string
}

func (o *Office) String() string {
	return fmt.Sprintf("id:%d, Office name:%s Description:%s", o.Id, o.Name, o.Description)
}
