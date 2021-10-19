package exchange

import (
	"errors"
	"fmt"
)

type ModelAccessor struct{}

func NewModelAccessor() *ModelAccessor {
	return &ModelAccessor{}
}

func (ma *ModelAccessor) Add(entity Exchange) uint64 {
	var newId uint64
	if len(entities) == 0 {
		entity.Id = 1
		newId = 1
	}
	if entity.Id == 0 {
		entity.Id = entities[len(entities) - 1].Id + 1
		newId = entity.Id
	}
	entities = append(entities, entity)
	return newId
}

func (ma *ModelAccessor) Get(id uint64) (*Exchange, bool) {
	wantedSliceLength := int(id) + 1
	if len(entities) < wantedSliceLength {
		return &Exchange{}, false
	}
	entity := entities[id]
	return &entity, true
}

func (ma *ModelAccessor) Remove(id uint64) (bool, error) {
	for key, item := range entities {
		if item.Id == id {
			entities = append(entities[:key], entities[key+1:]...)
			return true, nil
		}
	}
	errorMessage := fmt.Sprintf("Unable to remove Exchange with id %v.\n"+
		                        "Reason: Exchange not found.", id)
	err := errors.New(errorMessage)
	return false, err
}

func (ma *ModelAccessor) Entities() []Exchange {
	return entities
}