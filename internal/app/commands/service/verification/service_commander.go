package verification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/service/verification"
	"log"
)

type VerificationCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type ServiceVerificationCommander struct {
	bot                 *tgbotapi.BotAPI
	verificationService *verification.DummyVerificationService
}

func NewVerificationCommander(bot *tgbotapi.BotAPI,) *ServiceVerificationCommander {
	verificationService := verification.NewDummyVerificationService()

	return &ServiceVerificationCommander{
		bot:                 bot,
		verificationService: verificationService,
	}
}

func (c *ServiceVerificationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ServiceVerificationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ServiceVerificationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(0, msg)
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



