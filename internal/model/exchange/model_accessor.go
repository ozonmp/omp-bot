package exchange

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
	if len(entities) > 0 {
		for _, entity := range entities {
			if entity.Id == id {
				return &entity, true
			}
		}
	}
	return nil, false
}

func (ma *ModelAccessor) Remove(id uint64) bool {
	for key, item := range entities {
		if item.Id == id {
			entities = append(entities[:key], entities[key+1:]...)
			return true
		}
	}
	return false
}

func (ma *ModelAccessor) Replace(id uint64, entity Exchange) bool {
	for key, item := range entities {
		if item.Id == id {
			entities[key] = entity
			return true
		}
		continue
	}
	return false
}

func (ma *ModelAccessor) Entities() []Exchange {
	return entities
}
