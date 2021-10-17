package service

import (
	"errors"
	"fmt"
	"time"
)

func NewService() *Service {
	return &Service{}
}

var ErrIndexOutOfRange = errors.New("index out of range")
var NotImplementedError = errors.New("not implemented")

func (s *Service) List() []Service {
	return allEntities
}

func (s *Service) Get(idx int) (*Service, error) {
	if idx >= len(allEntities) {
		return nil, ErrIndexOutOfRange
	}
	return &allEntities[idx], nil
}

func (s *Service) Create(id int, from time.Time, to time.Time, contract string, deadline time.Time, duration time.Duration) (Service, error) {
	return Service{}, NotImplementedError
}

func (s Service) Update(from time.Time, to time.Time, contract string, deadline time.Time, duration time.Duration) (Service, error) {
	return s, NotImplementedError
}

func (s *Service) Remove(idx int) (bool, error) {
	if idx >= len(allEntities) {
		return false, ErrIndexOutOfRange
	}

	copy(allEntities[idx:], allEntities[idx+1:])
	allEntities = allEntities[:len(allEntities)-1]

	return true, nil
}

func (s Service) String() string {
	return fmt.Sprintf(
		"Id: %d\n"+
			"From: %v\n"+
			"To: %v\n"+
			"Contract: %v\n"+
			"Deadline: %v\n"+
			"Duration: %d",
		s.Id,
		s.From.String(),
		s.To.String(),
		s.Contract,
		s.Deadline.String(),
		s.Duration,
	)
}
