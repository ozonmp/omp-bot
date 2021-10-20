package group

import (
	"fmt"
	"strings"

	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Update(group_id uint64, group product.Group) error {
	group.Owner = strings.Trim(group.Owner, " ")

	if group.Owner == "" {
		return fmt.Errorf("Group's owner is empty")
	}

	c.groups[group_id] = group

	return nil
}
