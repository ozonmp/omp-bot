package bank

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/bank/operation"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BankCommander struct {
	bot                *tgbotapi.BotAPI
	operationCommander Commander
}

func NewBankCommander(
	bot *tgbotapi.BotAPI,
) *BankCommander {
	return &BankCommander{
		bot: bot,
		// operationCommander
		operationCommander: operation.NewBankOperationCommander(bot),
	}
}

func (c *BankCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "operation":
		c.operationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BankCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BankCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "operation":
		c.operationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BankCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}