package ticket

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"log"
)

type Service struct{}

func (s *Service) validateId(ticket_id uint64) error {
	if ticket_id < 1 {
		return errors.New("Ticket id can not be less that 1")
	}

	if ticket_id > uint64(len(allTickets)) {
		return fmt.Errorf("There is no such ticket %v", ticket_id)
	}

	return nil
}

func (s *Service) validateRequiredFields(ticket travel.Ticket) error {
	if ticket.User == nil {
		return errors.New("Ticket must have 'User' property.")
	}

	if ticket.Schedule == nil {
		return errors.New("Ticket must have 'Schedule' property.")
	}

	return nil
}

func (s *Service) Describe(ticket_id uint64) (*travel.Ticket, error) {
	log.Printf("Travel.TicketService: getting ticket with id %v", ticket_id)

	if err := s.validateId(ticket_id); err != nil {
		return nil, err
	}

	return &allTickets[ticket_id-1], nil
}

func (s *Service) List(cursor uint64, limit uint64) []travel.Ticket {
	log.Printf("Travel.TicketService: listing tickets in range from %v to %v", cursor, cursor+limit)

	count := s.Count()
	if cursor > count {
		return []travel.Ticket{}
	}

	maxIndex := cursor + limit
	if maxIndex > count {
		maxIndex = count
	}

	return allTickets[cursor:maxIndex]
}

func (s *Service) Create(new_ticket travel.Ticket) (uint64, error) {
	log.Printf("Travel.TicketService: creating ticket %#v", new_ticket)

	if err := s.validateRequiredFields(new_ticket); err != nil {
		return 0, err
	}

	allTickets = append(allTickets, new_ticket)

	return uint64(len(allTickets)), nil
}

func (s *Service) Update(ticket_id uint64, ticket travel.Ticket) error {
	log.Printf("Travel.TicketService: updating ticket %v with value: %#v", ticket_id, ticket)

	if err := s.validateId(ticket_id); err != nil {
		return err
	}
	if err := s.validateRequiredFields(ticket); err != nil {
		return nil
	}

	allTickets[ticket_id-1] = ticket

	return nil
}

func (s *Service) Remove(ticket_id uint64) (bool, error) {
	log.Printf("Travel.TicketService: deleting ticket with id %v", ticket_id)

	if err := s.validateId(ticket_id); err != nil {
		return false, err
	}

	index := ticket_id - 1

	allTickets = append(allTickets[:index], allTickets[index+1:]...)

	return true, nil
}

func (s *Service) Count() uint64 {
	return uint64(len(allTickets))
}
