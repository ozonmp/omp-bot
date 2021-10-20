package group

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/product"
)

type GroupService interface {
	Describe(groupID uint64) (*product.group, error)
	List() ([]product.Group, error)
	Create(product.Group) (uint64, error)
	Update(groupID uint64, group product.Group) error
	Remove(groupID uint64) (bool, error)
}

type DummyGroupService struct {
	groups []product.Group
}

func NewDummyGroupService() *DummyGroupService {
	return &DummyGroupService{}
}
