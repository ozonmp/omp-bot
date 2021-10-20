package estate

import (
	"fmt"
	"sort"
	"sync"
)

type ApartmentStorage interface {
	Describe(apartmentID uint64) (*Apartment, error)
	List(cursor uint64, limit uint64) ([]Apartment, error)
	Create(Apartment) (uint64, error)
	Update(apartmentID uint64, apartment Apartment) error
	Remove(apartmentID uint64) (bool, error)
}

type InMemoryApartmentStorage struct {
	mem    map[uint64]Apartment
	mu     sync.RWMutex
	nextID uint64
}

func NewInMemoryApartmentStorage(items []Apartment) *InMemoryApartmentStorage {
	mem := make(map[uint64]Apartment)
	for i, item := range items {
		item.ID = uint64(i)
		mem[item.ID] = item
	}
	return &InMemoryApartmentStorage{
		mem:    mem,
		nextID: uint64(len(items)),
	}
}

func NewEmptyInMemoryApartmentStorage() *InMemoryApartmentStorage {
	return NewInMemoryApartmentStorage(nil)
}

func (s *InMemoryApartmentStorage) Describe(apartmentID uint64) (apartment *Apartment, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tmp, ok := s.mem[apartmentID]
	if !ok {
		err = fmt.Errorf("no element with id=%d", apartmentID)
		return
	}
	apartment = &tmp
	return
}

func (s *InMemoryApartmentStorage) List(cursor uint64, limit uint64) (apartments []Apartment, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.mem) == 0 {
		return
	}
	all := make([]Apartment, len(s.mem))
	i := 0
	for _, item := range s.mem {
		all[i] = item
		i++
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].ID < all[j].ID
	})
	start := uint64(sort.Search(len(all), func(i int) bool {
		return all[i].ID >= cursor
	}))
	if start == uint64(len(all)) {
		return
	}
	var end uint64
	if start+limit > uint64(len(all)) {
		end = uint64(len(all))
	} else {
		end = start + limit
	}

	// make a copy so that the big array under "all" slice doesn't escape this function
	apartments = make([]Apartment, end-start)
	copy(apartments, all[start:end])
	return
}

func (s *InMemoryApartmentStorage) Create(apartment Apartment) (id uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id = s.nextID
	s.nextID++
	apartment.ID = id
	s.mem[id] = apartment
	return
}

func (s *InMemoryApartmentStorage) Update(apartmentID uint64, apartment Apartment) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.mem[apartmentID]
	if !ok {
		err = fmt.Errorf("no element with id=%d", apartmentID)
		return
	}
	apartment.ID = apartmentID
	s.mem[apartmentID] = apartment
	return
}

func (s *InMemoryApartmentStorage) Remove(apartmentID uint64) (ok bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok = s.mem[apartmentID]
	if ok {
		delete(s.mem, apartmentID)
	}
	return
}
