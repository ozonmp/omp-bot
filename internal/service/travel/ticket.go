package travel

import (
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"github.com/ozonmp/omp-bot/internal/service/travel/ticket"
)

type TicketService interface {
	Describe(ticket_id uint64) (*travel.Ticket, error)
	List(cursor uint64, limit uint64) []travel.Ticket
	Create(new_ticket travel.Ticket) (uint64, error)
	Update(ticket_id uint64, ticket travel.Ticket) error
	Remove(ticket_id uint64) (bool, error)
	Count() uint64
}

func NewTravelTicketService() TicketService {
	return &ticket.Service{}
}
