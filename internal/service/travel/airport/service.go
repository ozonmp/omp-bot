package airport

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Airport {
	return allAirports
}

func (s *Service) Get(idx int) (*Airport, error) {
	return &allAirports[idx], nil
}

func (s *Service) Delete(idx int) (*Airport, error) {
	// TODO delete
	return &allAirports[idx], nil
}
func (s *Service) Create(name string, location string) (Airport, error) {
	airport := Airport{name, location}
	// TODO addition
	return airport, nil
}

func (s *Service) Edit(idx int, name string, location string) (*Airport, error) {
	a := &allAirports[idx]
	fmt.Printf("Editing %s to %s, %s", a, name, location)
	return a, nil
}
