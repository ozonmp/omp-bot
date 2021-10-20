package group

import "fmt"

type Group struct {
	ID    uint64
	Owner string
	Items string
}

func (c *Group) String() string {
	return fmt.Sprintf(`ID: %d, Owner: %s, Items: %s`, c.ID, c.Owner, c.Items)
}
