package service

import "fmt"

type ServiceService interface {
	Get(service_id uint64) (*Service, error)
	Describe(service_id uint64) (*Service, error)
	List(cursor uint64, limit uint64) ([]Service, error)
	Create(Service) (uint64, error)
	Update(service_id uint64, service Service) error
	Remove(service_id uint64) (bool, error)
}

type DummyServiceService struct {
	data   map[uint64]*Service
	lastId uint64
}

func NewDummyServiceService(allEntities []Service) *DummyServiceService {
	data := make(map[uint64]*Service, len(allEntities))
	var lastId uint64 = 0
	for _, entity := range allEntities {
		lastId++
		data[lastId] = NewService(lastId, entity.Title, entity.Description)
	}
	return &DummyServiceService{data: data, lastId: lastId}
}

func (s *DummyServiceService) Get(service_id uint64) (*Service, error) {
	if service, exists := s.data[service_id]; exists {
		return service, nil
	} else {
		return nil, fmt.Errorf("Service with id=%d not found", service_id)
	}
}

func (s *DummyServiceService) Describe(service_id uint64) (*Service, error) {
	if service, exists := s.data[service_id]; exists {
		return service, nil
	} else {
		return nil, fmt.Errorf("Service with id=%d not found", service_id)
	}
}

func (s *DummyServiceService) List(cursor uint64, limit uint64) ([]Service, error) {
	var result = make([]Service, 0, limit)
	for ind, val := range s.data {
		if ind >= limit {
			break
		}
		result = append(result, *val)
	}
	return result, nil
}

func (s *DummyServiceService) Create(service Service) (uint64, error) {
	s.lastId++
	storedService := NewService(s.lastId, service.Title, service.Description)
	s.data[storedService.Id] = storedService
	return storedService.Id, nil
}

func (s *DummyServiceService) Update(service_id uint64, service Service) error {
	if _, ok := s.data[service_id]; ok {
		service.Id = service_id
		s.data[service.Id] = &service
		return nil
	}
	return fmt.Errorf("service with id=%d not found", service_id)
}

func (s *DummyServiceService) Remove(service_id uint64) (bool, error) {
	if _, ok := s.data[service_id]; ok {
		delete(s.data, service_id)
		return true, nil
	}
	return false, fmt.Errorf("service with id=%d not found", service_id)
}
