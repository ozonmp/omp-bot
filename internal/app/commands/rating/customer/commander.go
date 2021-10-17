package customer

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/rating/customer/paginator"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/rating"
	"github.com/ozonmp/omp-bot/internal/service/rating/customer"
)

const paginatorChunkSize = 2

type CustomerCommander struct {
	bot             *tgbotapi.BotAPI
	customerService DummyService
	paginator       *paginator.Paginator
}

func NewCustomerCommander(
	bot *tgbotapi.BotAPI,
) *CustomerCommander {
	customerService := customer.NewDummyService()

	return &CustomerCommander{
		bot:             bot,
		customerService: customerService,
		paginator:       paginator.NewPaginator(customerService, paginatorChunkSize),
	}
}

func (c *CustomerCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var err error
	switch callbackPath.CallbackName {
	case "list":
		err = c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CustomerCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}

	c.handleError(callback.Message.Chat.ID, err)
}

func (c *CustomerCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		err = c.Help(msg)
	case "list":
		err = c.List(msg)
	case "get":
		err = c.Get(msg)
	case "delete":
		err = c.Delete(msg)
	case "new":
		err = c.New(msg)
	case "edit":
		err = rating.NewUserError("not implemented")
	default:
		err = c.Default(msg)
	}

	c.handleError(msg.Chat.ID, err)
}

func (c *CustomerCommander) handleError(chatID int64, err error) {
	if err == nil {
		return
	}

	var outMsg tgbotapi.MessageConfig

	var userError rating.UserError
	if errors.As(err, &userError) {
		outMsg = tgbotapi.NewMessage(chatID, "Input data were wrong: "+userError.Error())
	} else {
		log.Printf("Internal error %v", err)
		outMsg = tgbotapi.NewMessage(chatID, "Got internal exception")
	}

	if _, err := c.bot.Send(outMsg); err != nil {
		log.Printf("Send error message error %v", err)
	}
}
