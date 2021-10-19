package insurance

import "fmt"

type Life struct {
	Id     uint64
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

func (life *Life) String() string {
	return fmt.Sprintf("Field1: %s, Field2: %s", life.Field1, life.Field2)
}
