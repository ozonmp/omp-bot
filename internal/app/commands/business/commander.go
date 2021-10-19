package business

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/company"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessCommander struct {
	bot				 *tgbotapi.BotAPI
	companyCommander Commander
}

func NewBusinessCommander (
	bot *tgbotapi.BotAPI,
) *BusinessCommander {
	return &BusinessCommander{
		bot: bot,
		// companyCommander
		companyCommander: company.NewCompanyCommander(bot),
	}
}

func (c *BusinessCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "company":
		c.companyCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BusinessCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "company":
		c.companyCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}