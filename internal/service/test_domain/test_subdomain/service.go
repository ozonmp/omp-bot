package testsubdomain

import (
	"errors"
	"fmt"
)

type TestSubdomainService struct {
	allEntities []TestSubdomain
}

var ErrEntityNotExists = errors.New("Entity Not Exists")

func (s TestSubdomainService) Describe(test_subdomainID uint64) (*TestSubdomain, error) {
	if len(s.allEntities) <= int(test_subdomainID) {
		return nil, ErrEntityNotExists
	}
	return &s.allEntities[test_subdomainID], nil
}

func (s TestSubdomainService) List(cursor uint64, limit uint64) ([]TestSubdomain, error) {
	curLen := uint64(len(s.allEntities))
	from := cursor
	if from+1 > curLen {
		return nil, ErrEntityNotExists
	}

	to := cursor + limit

	if to > curLen {
		to = curLen
	}

	res := make([]TestSubdomain, to-from)

	copy(res, s.allEntities[from:to])
	return res, nil
}

func (s *TestSubdomainService) Remove(verificationID uint64) (bool, error) {
	if s.allEntities == nil {
		return false, ErrEntityNotExists
	}
	if verificationID+1 > uint64(len(s.allEntities)) {
		return false, ErrEntityNotExists
	}
	s.allEntities = append(s.allEntities[:verificationID], s.allEntities[verificationID+1:]...)
	return true, nil
}

func (s *TestSubdomainService) Create(d TestSubdomain) (uint64, error) {
	return 0, fmt.Errorf("not implemented")
}

func (s *TestSubdomainService) Update(verificationID uint64, d TestSubdomain) error {
	return fmt.Errorf("not implemented")
}
