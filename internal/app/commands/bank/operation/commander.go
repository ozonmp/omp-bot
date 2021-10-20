package operation
import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/bank/operation"
)

type BankOperationCommander struct {
	bot              *tgbotapi.BotAPI
	operationService operation.ServiceInterface
}

func NewBankOperationCommander(
	bot *tgbotapi.BotAPI,
) *BankOperationCommander {
	bankOperationService := operation.NewDummyService()

	return &BankOperationCommander{
		bot:              bot,
		operationService: bankOperationService,
	}
}

func (c *BankOperationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BankOperationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BankOperationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		log.Printf("BankOperationCommander.HandleCommand: unknown command name: %s", commandPath.CommandName)
	}
}
