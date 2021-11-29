package travel

import (
	"github.com/ozonmp/omp-bot/internal/model/travel"
)

// TicketService provides capabilities to work with tickets
type TicketService interface {
	Describe(ticket_id uint64) (*travel.Ticket, error)
	List(cursor uint64, limit uint64) ([]travel.Ticket, uint64)
	Create(new_ticket travel.Ticket) (uint64, error)
	Update(ticket_id uint64, ticket travel.Ticket) error
	Remove(ticket_id uint64) (bool, error)
}
