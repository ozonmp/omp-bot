package bank

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/bank/transaction"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BankCommander struct {
	bot                  *tgbotapi.BotAPI
	transactionCommander Commander
}

func NewBankTransactionCommander(
	bot *tgbotapi.BotAPI,
) *BankCommander {
	return &BankCommander{
		bot: bot,
		// transactionCommander
		transactionCommander: transaction.NewBankTransactionCommander(bot),
	}
}

func (c *BankCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "transaction":
		c.transactionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BankCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BankCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "transaction":
		c.transactionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BankCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
