package serial

import (
	"errors"
	"fmt"
)

func (s *Serial) String() string {
	return fmt.Sprintf("id=%d title=%s year=%d", s.ID, s.Title, s.Year)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Serial {
	return allEntities
}

func (s *Service) Get(idx int) (*Serial, error) {
	if idx < 0 {
		return nil, errors.New("id out of range")
	}
	for i, v := range allEntities {
		if v.ID == idx {
			return &allEntities[i], nil
		}
	}
	return nil, errors.New("item not found")
}

func (s *Service) New(new Serial) error {
	for _, v := range allEntities {
		if v.ID == new.ID {
			return errors.New("item already exists")
		}
	}
	allEntities = append(allEntities, new)
	return nil
}

func (s *Service) Edit(idx int, new Serial) error {
	if new.ID == 0 && new.ID != idx {
		for _, v := range allEntities {
			if v.ID == new.ID {
				return errors.New("item already exists")
			}
		}
	}

	for i, v := range allEntities {
		if v.ID == idx {
			if new.ID != 0 {
				allEntities[i].ID = new.ID
			}
			if new.Title != "" {
				allEntities[i].Title = new.Title
			}
			if new.Year != 0 {
				allEntities[i].Year = new.Year
			}
			return nil
		}
	}
	return errors.New("item not found")
}

func (s *Service) Delete(idx int) error {
	if idx < 0 {
		return errors.New("id out of range")
	}
	for i, v := range allEntities {
		if v.ID == idx {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
