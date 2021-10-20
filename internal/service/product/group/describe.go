package group

import (
	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *DummyGroupService) Describe(group_id uint64) (*product.Group, error) {
	group := c.groups[group_id]

	return &group, nil
}
