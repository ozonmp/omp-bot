package click

import (
	"errors"
	"fmt"
	model "github.com/ozonmp/omp-bot/internal/model/actvity"
)

type ClickService interface {
	Describe(subdomainId uint64) (*model.Click, error)
	List(cursor uint64, limit uint64) []*model.Click
	Create(m model.Click) (uint64, error)
	Update(subdomainId uint64, subdomain model.Click) error
	Remove(subdomainId uint64) (bool, error)
}

type ActivityClickService struct {
	clicks []*model.Click
}

func NewActivityClickService() *ActivityClickService {
	return &ActivityClickService{
		clicks: []*model.Click{{Title: "my1"}, {Title: "my2"}, {Title: "my3"}}, // TODO: remove example data
	}
}

func (s *ActivityClickService) Describe(subdomainId uint64) (*model.Click, error) {
	if subdomainId > uint64(len(s.clicks)) {
		return nil, errors.New(s.getOutOfRangeErrorString("can't get item", subdomainId))
	}

	return s.clicks[subdomainId], nil
}

func (s *ActivityClickService) List(cursor uint64, limit uint64) []*model.Click {
	if cursor == 0 && limit == 0 {
		return s.clicks
	}

	l := uint64(len(s.clicks))

	if cursor+limit >= l {
		return s.clicks[cursor:]
	}

	return s.clicks[cursor : limit+cursor]
}

func (s *ActivityClickService) Create(m model.Click) (uint64, error) {
	s.clicks = append(s.clicks, &m)

	return uint64(len(s.clicks) - 1), nil
}

func (s *ActivityClickService) Update(subdomainId uint64, subdomain model.Click) error {
	length := uint64(len(s.clicks))

	if subdomainId >= length {
		return errors.New(s.getOutOfRangeErrorString("can't update item", subdomainId))
	}

	s.clicks[subdomainId] = &subdomain

	return nil
}

func (s *ActivityClickService) Remove(subdomainId uint64) (bool, error) {
	length := uint64(len(s.clicks))

	if subdomainId >= length {
		return false, errors.New(s.getOutOfRangeErrorString("can't remove item", subdomainId))
	}

	if subdomainId == 0 {
		s.clicks = s.clicks[0:]
	} else if subdomainId == length-1 {
		s.clicks = s.clicks[:length-1]
	} else {
		s.clicks = append(s.clicks[0:subdomainId], s.clicks[subdomainId+1:]...)
	}

	return true, nil
}

func (s *ActivityClickService) getOutOfRangeErrorString(action string, subdomainId uint64) string {
	length := len(s.clicks)

	return fmt.Sprintf("%s %d: now we have only %d items — from 0 to %d", action, subdomainId, length, length-1)
}
