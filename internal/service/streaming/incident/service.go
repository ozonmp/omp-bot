package incident

import (
	"errors"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []*Incident {
	result := make([]*Incident, 0, len(allEntities))
	for _, entity := range allEntities {
		result = append(result, entity)
	}
	return result
}

func (s *Service) Get(idx int) (*Incident, error) {
	if entity, ok := allEntities[idx]; ok {
		return entity, nil
	}
	return nil, errors.New("not found")
}

func (s *Service) Delete(idx int) error {
	if _, ok := allEntities[idx]; ok {
		delete(allEntities, idx)
		return nil
	}
	return errors.New("not found")
}

func (s *Service) New(entity Incident) error {
	if _, ok := allEntities[entity.Id]; !ok {
		allEntities[entity.Id] = &entity
		return nil
	}
	return errors.New("incident with such id already exists")
}

func (s *Service) Edit(entity Incident) error {
	if _, ok := allEntities[entity.Id]; ok {
		allEntities[entity.Id] = &entity
		return nil
	}
	return errors.New("incident with such id not exists")
}
