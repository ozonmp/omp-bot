package test

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Test {
	return allEntities
}

func (s *Service) Get(idx int) (*Test, error) {
	return &allEntities[idx], nil
}

func (s *Service) New(title string, description string, min_score int) error {
	var test Test
	test.id = allEntities[len(allEntities)-1].id + 1
	test.name = title
	test.description = description
	test.min_score = min_score

	allEntities = append(allEntities, test)
	return nil
}

func (s *Service) Edit(idx int, title string, description string, min_score int) error {
	for i := 0; i < len(allEntities); i++ {
		if allEntities[i].id == idx {
			allEntities[i].name = title
			allEntities[i].description = description
			allEntities[i].min_score = min_score
			return nil
		}
	}
	return fmt.Errorf("no element for edit")
}

func (s *Service) Delete(idx int) (bool, error) {
	for i := 0; i < len(allEntities); i++ {
		if allEntities[i].id == idx {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
		}
	}

	return true, nil
}
