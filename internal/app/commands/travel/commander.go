package travel

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
	trv_ticket_facade "github.com/ozonmp/trv-ticket-facade/pkg/trv-ticket-facade"

	"google.golang.org/grpc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/travel/ticket"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TravelCommander struct {
	bot             Sender
	ticketCommander Commander
}

func NewTravelCommander(
	bot Sender,
) *TravelCommander {
	apiAddress := os.Getenv("TRV_TICKET_API_ADDRESS")
	facadeAddress := os.Getenv("TRV_TICKET_FACADE_ADDRESS")

	travelTicketApiConn, err := grpc.Dial(apiAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	travelTicketFacadeConn, err := grpc.Dial(facadeAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		travelTicketApiConn.Close()
		travelTicketFacadeConn.Close()
	}()

	travelTicketApiClient := trv_ticket_api.NewTravelTicketApiServiceClient(travelTicketApiConn)
	travelTicketFacadeClient := trv_ticket_facade.NewTravelTicketFacadeServiceClient(travelTicketFacadeConn)

	return &TravelCommander{
		bot: bot,
		// subdomainCommander
		ticketCommander: ticket.NewTravelTicketCommander(
			context.Background(),
			travelTicketApiClient,
			travelTicketFacadeClient,
			bot,
		),
	}
}

func (c *TravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "ticket":
		c.ticketCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("TravelCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *TravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "ticket":
		c.ticketCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("TravelCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
