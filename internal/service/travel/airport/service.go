package airport

import (
	"errors"
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Airport {
	return allAirports
}

func (s *Service) Get(idx int) (*Airport, error) {
	if idx > len(allAirports) {
		return &Airport{}, errors.New("couldn't find such airport")
	}
	return &allAirports[idx], nil
}

func (s *Service) Delete(idx int) (*Airport, error) {
	if idx > len(allAirports) {
		return &Airport{}, errors.New("couldn't find such airport")
	}
	deletedAirport := allAirports[idx]
	allAirports = append(allAirports[:idx], allAirports[idx+1:]...)
	return &deletedAirport, nil
}
func (s *Service) New(name string, location string) (Airport, error) {
	airport := Airport{name, location}
	allAirports = append(allAirports, airport)
	return airport, nil
}

func (s *Service) Edit(idx int, name string, location string) (*Airport, error) {
	a := &allAirports[idx]
	a.name = name
	a.location = location
	fmt.Printf("Editing %s to %s, %s", a, name, location)
	return a, nil
}
