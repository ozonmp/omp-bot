package router

import (
	"fmt"
	"log"
	"runtime/debug"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/demo"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

var (
	ErrDomainNotImplemented = fmt.Errorf("domain not implemented")
	ErrUnknownDomain        = fmt.Errorf("unknown domain")
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(callback *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	// bot
	bot *tgbotapi.BotAPI

	// demoCommander
	demoCommander Commander
	// user
	// access
	// buy
	// delivery
	// recommendation
	// travel
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

func NewRouter(
	bot *tgbotapi.BotAPI,
) *Router {
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

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
			log.Printf("stack trace: %s", debug.Stack())
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "demo":
		r.demoCommander.HandleCallback(callback, callbackPath)
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
		break
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

func (r *Router) routeCommand(msg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		return
	}

	switch commandPath.Domain {
	case "demo":
		r.demoCommander.HandleCommand(msg, commandPath)
	case "user":
		err = ErrDomainNotImplemented
	case "access":
		err = ErrDomainNotImplemented
	case "buy":
		err = ErrDomainNotImplemented
	case "delivery":
		err = ErrDomainNotImplemented
	case "recommendation":
		err = ErrDomainNotImplemented
	case "travel":
		err = ErrDomainNotImplemented
	case "loyalty":
		err = ErrDomainNotImplemented
	case "bank":
		err = ErrDomainNotImplemented
	case "subscription":
		err = ErrDomainNotImplemented
	case "license":
		err = ErrDomainNotImplemented
	case "insurance":
		err = ErrDomainNotImplemented
	case "payment":
		err = ErrDomainNotImplemented
	case "storage":
		err = ErrDomainNotImplemented
	case "streaming":
		err = ErrDomainNotImplemented
	case "business":
		err = ErrDomainNotImplemented
	case "work":
		err = ErrDomainNotImplemented
	case "service":
		err = ErrDomainNotImplemented
	case "exchange":
		err = ErrDomainNotImplemented
	case "estate":
		err = ErrDomainNotImplemented
	case "rating":
		err = ErrDomainNotImplemented
	case "security":
		err = ErrDomainNotImplemented
	case "cinema":
		err = ErrDomainNotImplemented
	case "logistic":
		err = ErrDomainNotImplemented
	case "product":
		err = ErrDomainNotImplemented
	case "education":
		err = ErrDomainNotImplemented
	default:
		err = ErrUnknownDomain
	}
	return
}

func (r *Router) createResponseForCommandError(msg *tgbotapi.Message, err error) (resp tgbotapi.MessageConfig) {
	resp = tgbotapi.NewMessage(msg.Chat.ID,
		fmt.Sprintf("I'm sorry, something bad happend with your command (%s): %v", msg.Command(), err),
	)
	return
}

func (r *Router) createResponseForCommand(msg *tgbotapi.Message) (resp tgbotapi.MessageConfig) {
	resp, err := r.routeCommand(msg)
	if err != nil {
		resp = r.createResponseForCommandError(msg, err)
		return
	}
	return
}

func (r *Router) createResponseForMessage(msg *tgbotapi.Message) (resp tgbotapi.MessageConfig) {
	if !msg.IsCommand() {
		resp = createResponseWithHint(msg)
		return
	}

	resp = r.createResponseForCommand(msg)
	return
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	resp := r.createResponseForMessage(msg)

	if resp.ChatID != 0 {
		// HACK
		// TODO: remove this check after we are done with refactoring of
		// message handling, right now not all paths in the above call
		// return correct response

		_, err := r.bot.Send(resp)
		if err != nil {
			log.Printf("Failed to send response: %v", err)
		}
	}
}

func createResponseWithHint(msg *tgbotapi.Message) (resp tgbotapi.MessageConfig) {
	resp = tgbotapi.NewMessage(msg.Chat.ID,
		"Hey! I don't know how to chat. You can send me a command:\n"+
			"Command format: /{command}__{domain}__{subdomain}",
	)
	return
}
