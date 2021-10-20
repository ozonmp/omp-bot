package bank

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/bank/card"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/bank/card"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DummyBankCommander struct {
	bot           *tgbotapi.BotAPI
	cardCommander Commander
}

func NewBankCommander(
	bot *tgbotapi.BotAPI,
) *DummyBankCommander {
	service:= service.NewDummyCardService()
	return &DummyBankCommander{
		bot:           bot,
		cardCommander: card.NewCardCommander(bot, service),
	}
}

func (c *DummyBankCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "card":
		c.cardCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BankCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *DummyBankCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "card":
		c.cardCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BankCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}