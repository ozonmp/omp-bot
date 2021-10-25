package access

import (
	"errors"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type AccessService interface {
	List(cursor uint64, limit uint64) []Access
	Describe(accessID uint64) (*Access, error)
	Remove(accessID uint64) (bool, error)
	Update(accessID uint64, access Access) error
	Create(Access) (uint64, error)
	String(p Access) string
}

func (s *Service) List(cursor uint64, limit uint64) []Access {
	return allEntities
}

func (s *Service) Describe(accessID uint64) (*Access, error) {
	for i := range allEntities {
		if allEntities[i].ID == accessID {
			return &allEntities[i], nil
		}
	}
	return nil, errors.New("Account not found")
}

func (s *Service) Remove(accessID uint64) (bool, error) {
	for i := range allEntities {
		if allEntities[i].ID == accessID {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) Update(accessID uint64, access Access) error {
	for i := range allEntities {
		if allEntities[i].ID == accessID {
			allEntities[i] = access
			return nil
		}
	}
	return errors.New("Account not found")
}

func (s *Service) Create(acc Access) (uint64, error) {
	var newId uint64
	if len(allEntities) == 0 {
		newId = 1
	} else {
		newId = allEntities[len(allEntities)-1].ID + 1
	}
	allEntities = append(allEntities, acc)
	return newId, nil
}

func (s *Service) String(acc Access) string {
	var outstr = "ID: " + strconv.FormatUint(acc.ID, 10) +
		" Role_ID: " + strconv.FormatUint(acc.Role_ID, 10) +
		" Resource_ID: " + strconv.FormatUint(acc.Resource_ID, 10)

	return outstr
}
