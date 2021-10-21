package recommendation

import "fmt"

type Production struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int64  `json:"rating"`
}

func (p Production) String() string {
	return fmt.Sprintf("%v %v %v %v", p.Id, p.Title, p.Description, p.Rating)
}
