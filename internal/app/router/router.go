package router

import (
	"context"
	"log"

	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
	trv_ticket_facade "github.com/ozonmp/trv-ticket-facade/pkg/trv-ticket-facade"

	"github.com/ozonmp/omp-bot/internal/app/commands/travel"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/demo"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

// TelegramBotAPIWrapper is the wrapper for Telegram bot API with additional properties required for routing
type TelegramBotAPIWrapper struct {
	tgbotapi.BotAPI

	Ctx context.Context

	TravelTicketApiClient    trv_ticket_api.TravelTicketApiServiceClient
	TravelTicketFacadeClient trv_ticket_facade.TravelTicketFacadeServiceClient
}

type Router struct {
	// bot
	bot *TelegramBotAPIWrapper

	// demoCommander
	demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
	travelCommander Commander
	// loyalty
	// bank
	// subscription
	// license
	// insurance
	// payment
	// storage
	// streaming
	// business
	// work
	// service
	// exchange
	// estate
	// rating
	// security
	// cinema
	// logistic
	// product
	// education
}

func NewRouter(bot *TelegramBotAPIWrapper) *Router {
	return &Router{
		// bot
		bot: bot,
		// demoCommander
		demoCommander: demo.NewDemoCommander(bot),
		// user
		// access
		// buy
		// delivery
		// recommendation
		// travel
		travelCommander: travel.NewTravelCommander(
			bot.Ctx,
			bot.TravelTicketApiClient,
			bot.TravelTicketFacadeClient,
			bot,
		),
		// loyalty
		// bank
		// subscription
		// license
		// insurance
		// payment
		// storage
		// streaming
		// business
		// work
		// service
		// exchange
		// estate
		// rating
		// security
		// cinema
		// logistic
		// product
		// education
	}
}

func (c *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		c.demoCommander.HandleCallback(callback, callbackPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		c.travelCommander.HandleCallback(callback, callbackPath)
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", callbackPath.Domain)
	}
}

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Domain {
	case "demo":
		c.demoCommander.HandleCommand(msg, commandPath)
	case "user":
		break
	case "access":
		break
	case "buy":
		break
	case "delivery":
		break
	case "recommendation":
		break
	case "travel":
		c.travelCommander.HandleCommand(msg, commandPath)
	case "loyalty":
		break
	case "bank":
		break
	case "subscription":
		break
	case "license":
		break
	case "insurance":
		break
	case "payment":
		break
	case "storage":
		break
	case "streaming":
		break
	case "business":
		break
	case "work":
		break
	case "service":
		break
	case "exchange":
		break
	case "estate":
		break
	case "rating":
		break
	case "security":
		break
	case "cinema":
		break
	case "logistic":
		break
	case "product":
		break
	case "education":
		break
	default:
		log.Printf("Router.handleCallback: unknown domain - %s", commandPath.Domain)
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command}__{domain}__{subdomain}")

	c.bot.Send(outputMsg)
}
