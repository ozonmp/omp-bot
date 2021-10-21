package storage

import (
	"fmt"
	"sort"
	"sync"
)

type MemoryStorage struct {
	currentId uint64
	cache     []*RentEntity
	entities  map[uint64]*RentEntity
	mx        sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{cache: make([]*RentEntity, 0), entities: make(map[uint64]*RentEntity)}
}

func (s *MemoryStorage) Count() uint64 {
	defer s.mx.Unlock()
	s.mx.Lock()
	s.updateTempSlice()
	return uint64(len(s.cache))
}

func (s *MemoryStorage) Read(rentID uint64) (*RentEntity, error) {
	s.mx.RLock()
	model, ok := s.entities[rentID]
	s.mx.RUnlock()
	if !ok {
		return nil, fmt.Errorf("model not found (ID: %v)", rentID)
	}
	return model, nil
}

func (s *MemoryStorage) ReadAll() []*RentEntity {
	defer s.mx.Unlock()
	s.mx.Lock()
	s.updateTempSlice()

	return s.cache
}

func (s *MemoryStorage) ReadPage(cursor uint64, limit uint64) ([]*RentEntity, error) {
	defer s.mx.Unlock()
	s.mx.Lock()
	length := uint64(len(s.entities))
	lastIndex := cursor + limit

	if cursor < 0 || lastIndex > length || limit == 0 {
		return nil, nil
	}

	s.updateTempSlice()
	res := append([]*RentEntity(nil), s.cache[cursor:cursor+limit]...)
	return res, nil
}

func (s *MemoryStorage) Create(entity RentEntity) (uint64, error) {
	defer s.mx.Unlock()
	s.mx.Lock()
	s.updateTempSlice()
	s.currentId++
	entity.ID = s.currentId
	s.cache = append(s.cache, &entity)
	s.entities[s.currentId] = &entity

	return s.currentId, nil
}

func (s *MemoryStorage) Update(entity RentEntity) error {
	defer s.mx.Unlock()
	var err error = nil
	s.mx.Lock()
	s.updateTempSlice()
	if _, ok := s.entities[entity.ID]; ok {
		current := s.entities[entity.ID]
		current.ObjectType = entity.ObjectType
		current.ObjectInfo = entity.ObjectInfo
		current.RenterId = entity.RenterId
		current.Price = entity.Price
	} else {
		err = fmt.Errorf("entity not found (ID: %v)", entity.ID)
	}
	return err
}

func (s *MemoryStorage) Delete(rentID uint64) (bool, error) {
	defer s.mx.Unlock()
	s.mx.Lock()
	var err error = nil
	if _, ok := s.entities[rentID]; ok {
		delete(s.entities, rentID)
	} else {
		err = fmt.Errorf("entity not found (ID: %v)", rentID)
	}
	s.clearTempSlice()
	return err == nil, err
}

func (s *MemoryStorage) updateTempSlice() {
	if len(s.cache) > 0 {
		return
	}

	for _, v := range s.entities {
		s.cache = append(s.cache, v)
	}

	sort.Slice(s.cache, func(i, j int) bool {
		return s.cache[i].ID < s.cache[j].ID
	})
}

func (s *MemoryStorage) clearTempSlice() {
	s.cache = s.cache[:0]
}