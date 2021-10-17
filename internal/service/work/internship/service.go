package internship

import (
	"errors"
	"log"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Internship {
	return allEntities
}

func (s *Service) ShortString(p Internship) string {
	var result string = ""
	result += "ID: " + strconv.Itoa(p.Id) + ":"
	if &p.Description != nil {
		result += " Description: " + p.Description
	}
	return result
}

func (s *Service) FullString(p Internship) string {
	var result string = ""
	result += "ID: " + strconv.Itoa(p.Id) + ":"
	result += " Team - " + strconv.Itoa(p.Team_id) + ";"
	if &p.Description != nil {
		result += " Description: " + p.Description + ";"
	}
	if &p.Period != nil {
		result += " Period: " + p.Period + ";"
	}
	if &p.Compensation != nil {
		if p.Compensation {
			result += " compensation: yes."
		} else {
			result += " compensation: no."
		}
	}
	return result
}

func (s *Service) Get(idx int) (*Internship, error) {
	var resId int = -1
	for i := 0; i < len(allEntities); i++ {
		if idx == allEntities[i].Id {
			resId = i
			break
		}
	}
	if resId < 0 {
		err := errors.New("id not found")
		log.Printf("intership.Service.Get: id not found - %v", err)
		return nil, err
	}
	return &allEntities[resId], nil
}

func (s *Service) Delete(idx int) bool {
	var result bool = false
	for i := 0; i < len(allEntities); i++ {
		if idx == allEntities[i].Id {
			result = true
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			break
		}
	}
	return result
}
