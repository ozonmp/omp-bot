package customer

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/general_errors"
	"github.com/ozonmp/omp-bot/internal/service/rating/customer"
)

type CustomerCommander struct {
	bot             *tgbotapi.BotAPI
	customerService *customer.DummyService
}

func NewCustomerCommander(
	bot *tgbotapi.BotAPI,
) *CustomerCommander {
	customerService := customer.NewDummyService()

	return &CustomerCommander{
		bot:             bot,
		customerService: customerService,
	}
}

func (c *CustomerCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CustomerCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CustomerCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		err = c.Get(msg)
	case "delete":
		err = c.Delete(msg)
	case "new":
		err = general_errors.NewUserError("not implemented")
	case "edit":
		err = general_errors.NewUserError("not implemented")
	default:
		c.Default(msg)
	}
	if err == nil {
		return
	}
	var userError general_errors.UserError
	var outMsg tgbotapi.MessageConfig
	if errors.As(err, &userError) {
		outMsg = tgbotapi.NewMessage(msg.Chat.ID, "Input data were wrong: "+userError.Error())
	} else {
		log.Printf("Internal error %v", err)
		outMsg = tgbotapi.NewMessage(msg.Chat.ID, "Got internal exception")
	}

	if _, err := c.bot.Send(outMsg); err != nil {
		log.Printf("Send message errro %v", err)
	}

}
