package test

import "errors"

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

func (s *Service) New(title string) error {
	var test Test
	test.Title = title
	allEntities = append(allEntities, test)
	return nil
}

func (s *Service) Edit(idx int, title string) error {
	if idx > len(allEntities) {
		return errors.New("incorrect id")
	}

	var test Test
	test.Title = title
	allEntities[idx] = test
	return nil
}

func (s *Service) Delete(idx int) (bool, error) {
	if idx > len(allEntities) {
		return false, nil
	}

	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
	return true, nil
}
