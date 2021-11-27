package ticket

import (
	"context"
	"errors"
	"log"

	"github.com/ozonmp/omp-bot/internal/model/travel"
	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
)

// Service handles tickets
type Service struct {
	client trv_ticket_api.TravelTicketApiServiceClient
	ctx    context.Context
}

// NewTravelTicketService creates new service to handle tickets
func NewTravelTicketService(
	ctx context.Context,
	client trv_ticket_api.TravelTicketApiServiceClient) *Service {
	return &Service{
		client: client,
		ctx:    ctx,
	}
}

func (s *Service) validateId(ticket_id uint64) error {
	if ticket_id < 1 {
		return errors.New("ticket id can not be less that 1")
	}

	return nil
}

func (s *Service) validateRequiredFields(ticket travel.Ticket) error {
	if ticket.User == nil {
		return errors.New("ticket must have 'User' property")
	}

	if ticket.User.ID == 0 {
		return errors.New("user ID must be greater than 0")
	}

	if ticket.Schedule == nil {
		return errors.New("ticket must have 'Schedule' property")
	}

	if ticket.Schedule.ID == 0 {
		return errors.New("schedule ID must be greater than 0")
	}

	return nil
}

// Describe a ticket
func (s *Service) Describe(ticket_id uint64) (*travel.Ticket, error) {
	log.Printf("Travel.TicketService: getting ticket with id %v", ticket_id)

	if err := s.validateId(ticket_id); err != nil {
		return nil, err
	}

	res, err := s.client.DescribeTicketV1(s.ctx, &trv_ticket_api.DescribeTicketV1Request{
		TicketId: ticket_id,
	})

	if err != nil {
		return nil, err
	}

	ticket := &travel.Ticket{
		Seat:     res.GetData().GetSeat(),
		Comments: res.GetData().GetComment(),
	}

	if user := res.GetData().GetUser(); user != nil {
		ticket.User = &travel.User{
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}
	}

	if schedule := res.GetData().GetSchedule(); schedule != nil {
		ticket.Schedule = &travel.Schedule{
			Destination: schedule.Destination,
			Departure:   schedule.GetDeparture().AsTime(),
			Arrival:     schedule.GetArrival().AsTime(),
		}
	}

	return ticket, nil
}

// List some tickets
func (s *Service) List(cursor uint64, limit uint64) []travel.Ticket {
	log.Printf("Travel.TicketService: listing tickets in range from %v to %v", cursor, cursor+limit)

	res, err := s.client.ListTicketsV1(s.ctx, &trv_ticket_api.ListTicketsV1Request{
		Limit:  limit,
		Offset: cursor,
	})

	if err != nil {
		return []travel.Ticket{}
	}

	tickets := make([]travel.Ticket, 0, len(res.GetItems()))
	for _, t := range res.GetItems() {
		ticket := travel.Ticket{
			Seat:     t.GetSeat(),
			Comments: t.GetComment(),
		}

		if user := t.GetUser(); user != nil {
			ticket.User = &travel.User{
				FirstName: user.GetFirstName(),
				LastName:  user.GetLastName(),
			}
		}

		if schedule := t.GetSchedule(); schedule != nil {
			ticket.Schedule = &travel.Schedule{
				Destination: schedule.Destination,
				Departure:   schedule.GetDeparture().AsTime(),
				Arrival:     schedule.GetArrival().AsTime(),
			}
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}

// Create a ticket
func (s *Service) Create(new_ticket travel.Ticket) (uint64, error) {
	log.Printf("Travel.TicketService: creating ticket %#v", new_ticket)

	if err := s.validateRequiredFields(new_ticket); err != nil {
		return 0, err
	}

	res, err := s.client.CreateTicketV1(s.ctx, &trv_ticket_api.CreateTicketV1Request{
		UserId:     new_ticket.User.ID,
		Seat:       new_ticket.Seat,
		ScheduleId: new_ticket.Schedule.ID,
		Comment:    new_ticket.Comments,
	})

	if err != nil {
		return 0, err
	}

	return res.GetData().TicketId, nil
}

// Update a ticket
func (s *Service) Update(ticket_id uint64, ticket travel.Ticket) error {
	log.Printf("Travel.TicketService: updating ticket %v with value: %#v", ticket_id, ticket)

	if err := s.validateId(ticket_id); err != nil {
		return err
	}
	if err := s.validateRequiredFields(ticket); err != nil {
		return nil
	}

	_, err := s.client.UpdateTicketV1(s.ctx, &trv_ticket_api.UpdateTicketV1Request{
		TicketId:   ticket_id,
		Seat:       ticket.Seat,
		ScheduleId: ticket.Schedule.ID,
		Comment:    ticket.Comments,
	})

	return err
}

// Remove a ticket
func (s *Service) Remove(ticket_id uint64) (bool, error) {
	log.Printf("Travel.TicketService: deleting ticket with id %v", ticket_id)

	if err := s.validateId(ticket_id); err != nil {
		return false, err
	}

	_, err := s.client.RemoveTicketV1(s.ctx, &trv_ticket_api.RemoveTicketV1Request{TicketId: ticket_id})

	if err != nil {
		return false, err
	}

	return true, nil
}
