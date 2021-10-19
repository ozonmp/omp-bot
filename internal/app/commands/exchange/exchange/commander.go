package exchange

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/exchange/exchange"
	"log"
)

type ExchangeCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)    // return error not implemented
	Edit(inputMsg *tgbotapi.Message)   // return error not implemented
}

type SubdomainCommander struct {
	bot             *tgbotapi.BotAPI
	exchangeService *exchange.ExchangeService
}

func NewExchangeCommander(bot *tgbotapi.BotAPI, service *exchange.ExchangeService) *SubdomainCommander {
	return &SubdomainCommander {
		bot:              bot,
		exchangeService : service,
	}
}

func (c SubdomainCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		    "/help__exchange__exchange — print list of commands\n"+
			"/get__exchange__exchange — get an entity\n"+
			"/list__exchange__exchange — get a list of your entities\n"+
			"/delete__exchange__exchange  — delete an existing entity\n"+
			"\n"+
			"/new__exchange__exchange  — create a new entity\n"+
			"/edit__exchange__exchange  — edit an entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ExchangeCommander.Help: error sending reply message to chat - %v", err)
	}
}

func (c *SubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	}
}

func (c *SubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	fmt.Println("TBD")
}
