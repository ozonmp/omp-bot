package service

import "fmt"

type Verification struct {
	ID uint64
	Name string
}

func (s *Verification) String() string {
	return fmt.Sprintf("Verification id: %d, Name: %s", s.ID, s.Name)
}
