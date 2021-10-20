package group

import (
	"github.com/ozonmp/omp-bot/internal/model/product"
)

func (c *DummyGroupService) List() ([]product.Group, error) {
	return c.groups, nil
}
