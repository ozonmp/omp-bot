package transaction

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/bank/transaction"
)

type BankTransactionCommander struct {
	bot                *tgbotapi.BotAPI
	transactionService transaction.Service
	cursor             uint64
	limit              uint64
}

func NewBankTransactionCommander(
	bot *tgbotapi.BotAPI,
) *BankTransactionCommander {
	TransactionService := transaction.NewDummyTransactionService()

	return &BankTransactionCommander{
		bot:                bot,
		transactionService: TransactionService,
		cursor:             0,
		limit:              3,
	}
}

func (c *BankTransactionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BankTransactionCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BankTransactionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
