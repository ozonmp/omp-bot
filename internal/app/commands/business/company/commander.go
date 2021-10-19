package company

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/company"
)

type CompanyCommander struct {
	bot			   *tgbotapi.BotAPI
	companyService *company.DummyCompanyService
}

func NewCompanyCommander(
	bot *tgbotapi.BotAPI,
) *CompanyCommander {
	companyService := company.NewDummyCompanyService()

	return &CompanyCommander{
		bot:			bot,
		companyService: companyService,
	}
}

func (c *CompanyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CompanyCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CompanyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	default:
		c.Default(msg)
	}
}