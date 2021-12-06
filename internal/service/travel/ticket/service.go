package ticket

import (
	"context"
	"errors"
	"log"

	trv_ticket_facade "github.com/ozonmp/trv-ticket-facade/pkg/trv-ticket-facade"

	"github.com/ozonmp/omp-bot/internal/model/travel"
	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
)

// Service handles tickets
type Service struct {
	apiClient    trv_ticket_api.TravelTicketApiServiceClient
	facadeClient trv_ticket_facade.TravelTicketFacadeServiceClient
	ctx          context.Context
}

// NewTravelTicketService creates new service to handle tickets
func NewTravelTicketService(
	ctx context.Context,
	apiClient trv_ticket_api.TravelTicketApiServiceClient,
	facadeClient trv_ticket_facade.TravelTicketFacadeServiceClient) *Service {
	return &Service{
		apiClient:    apiClient,
		facadeClient: facadeClient,
		ctx:          ctx,
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

func (s *Service) mapToTicketFromApi(res *trv_ticket_api.Ticket) *travel.Ticket {
	if res == nil {
		return nil
	}

	ticket := &travel.Ticket{
		ID:       res.GetTicketId(),
		Seat:     res.GetSeat(),
		Comments: res.GetComment(),
	}

	if user := res.GetUser(); user != nil {
		ticket.User = &travel.User{
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}
	}

	if schedule := res.GetSchedule(); schedule != nil {
		ticket.Schedule = &travel.Schedule{
			Destination: schedule.Destination,
			Departure:   schedule.GetDeparture().AsTime(),
			Arrival:     schedule.GetArrival().AsTime(),
		}
	}

	return ticket
}

func (s *Service) mapToTicketFromFacade(res *trv_ticket_facade.Ticket) *travel.Ticket {
	if res == nil {
		return nil
	}

	ticket := &travel.Ticket{
		ID:       res.GetTicketId(),
		Seat:     res.GetSeat(),
		Comments: res.GetComment(),
	}

	if user := res.GetUser(); user != nil {
		ticket.User = &travel.User{
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
		}
	}

	if schedule := res.GetSchedule(); schedule != nil {
		ticket.Schedule = &travel.Schedule{
			Destination: schedule.Destination,
			Departure:   schedule.GetDeparture().AsTime(),
			Arrival:     schedule.GetArrival().AsTime(),
		}
	}

	return ticket
}

// Describe a ticket
func (s *Service) Describe(ticket_id uint64) (*travel.Ticket, error) {
	log.Printf("Travel.TicketService: getting ticket with id %v", ticket_id)

	if err := s.validateId(ticket_id); err != nil {
		return nil, err
	}

	res, err := s.apiClient.DescribeTicketV1(s.ctx, &trv_ticket_api.DescribeTicketV1Request{
		TicketId: ticket_id,
	})

	if err != nil {
		return nil, err
	}

	return s.mapToTicketFromApi(res.GetData()), nil
}

// List some tickets
func (s *Service) List(cursor uint64, limit uint64) ([]travel.Ticket, uint64) {
	log.Printf("Travel.TicketService: listing tickets in range from %v to %v", cursor, cursor+limit)

	res, err := s.facadeClient.ListTicketsV1(s.ctx, &trv_ticket_facade.ListTicketsV1Request{
		Limit:  limit,
		Offset: cursor,
	})

	if err != nil {
		return []travel.Ticket{}, 0
	}

	tickets := make([]travel.Ticket, 0, len(res.GetItems()))
	for _, t := range res.GetItems() {
		ticket := s.mapToTicketFromFacade(t)
		if ticket == nil {
			log.Printf("Travel.TicketService: failed to map ticket from %v. Skipping", t)
			continue
		}

		tickets = append(tickets, *ticket)
	}

	return tickets, res.GetTotal()
}

// Create a ticket
func (s *Service) Create(new_ticket travel.Ticket) (uint64, error) {
	log.Printf("Travel.TicketService: creating ticket %#v", new_ticket)

	if err := s.validateRequiredFields(new_ticket); err != nil {
		return 0, err
	}

	res, err := s.apiClient.CreateTicketV1(s.ctx, &trv_ticket_api.CreateTicketV1Request{
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

	_, err := s.apiClient.UpdateTicketV1(s.ctx, &trv_ticket_api.UpdateTicketV1Request{
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

	_, err := s.apiClient.RemoveTicketV1(s.ctx, &trv_ticket_api.RemoveTicketV1Request{TicketId: ticket_id})

	if err != nil {
		return false, err
	}

	return true, nil
}
