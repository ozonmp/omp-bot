package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/loyalty/certificate"
	"log"
)

type LoyaltyCertificateCommander struct {
	bot                *tgbotapi.BotAPI
	certificateService *certificate.DummyCertificateService
}

func NewLoyaltyCertificateCommander(
	bot *tgbotapi.BotAPI,
) *LoyaltyCertificateCommander {
	certificateService := certificate.NewDummyСertificateService()

	return &LoyaltyCertificateCommander{
		bot:                bot,
		certificateService: certificateService,
	}
}

func (c *LoyaltyCertificateCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LoyaltyCertificateCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LoyaltyCertificateCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
