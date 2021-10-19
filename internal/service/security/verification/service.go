package verification

import (
	"errors"
	"fmt"
)

type VerificationService struct {
	allEntities []Verification
}

func NewVerificationService(allEntities []Verification) *VerificationService {
	return &VerificationService{
		allEntities: allEntities,
	}
}

var ErrEntityNotExists = errors.New("Entity Not Exists")

func (s VerificationService) Describe(test_subdomainID uint64) (*Verification, error) {
	if len(s.allEntities) <= int(test_subdomainID) {
		return nil, ErrEntityNotExists
	}
	return &s.allEntities[test_subdomainID], nil
}

func (s VerificationService) List(cursor uint64, limit uint64) ([]Verification, error) {
	curLen := uint64(len(s.allEntities))
	from := cursor
	if from+1 > curLen {
		return nil, ErrEntityNotExists
	}

	to := cursor + limit

	if to > curLen {
		to = curLen
	}

	res := make([]Verification, to-from)

	copy(res, s.allEntities[from:to])
	return res, nil
}

func (s *VerificationService) Remove(verificationID uint64) (bool, error) {
	if s.allEntities == nil {
		return false, ErrEntityNotExists
	}
	if verificationID+1 > uint64(len(s.allEntities)) {
		return false, ErrEntityNotExists
	}
	s.allEntities = append(s.allEntities[:verificationID], s.allEntities[verificationID+1:]...)
	return true, nil
}

func (s *VerificationService) Create(d Verification) (uint64, error) {
	return 0, fmt.Errorf("not implemented")
}

func (s *VerificationService) Update(verificationID uint64, d Verification) error {
	return fmt.Errorf("not implemented")
}
