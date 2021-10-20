package group

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Create(group product.Group) (uint64, error) {
	group.Owner = strings.Trim(group.Owner, " ")

	if group.Owner == "" {
		return 0, fmt.Errorf("group owner is empty")
	}
	group.Items = strings.Trim(group.Items, " ")

	if group.Items == "" {
		return 0, fmt.Errorf("group items is empty")
	}
	var newID uint64 = 1

	if len(c.groups) != 0 {
		newID = c.groups[len(c.groups)-1].ID + 1
	}

	c.groups = append(c.groups, product.Group{
		ID:    newID,
		Owner: group.Owner,
		Items: group.Items,
	})

	return newID, nil
}
