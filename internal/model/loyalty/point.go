package loyalty

import "fmt"

type Point struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}


func (o *Point) String() string {
	return fmt.Sprintf("id:%d\nPoint name:%s\nDescription:%s", (o.Id +1), o.Name, o.Description)
}