package insurance

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/insurance/life"
	"github.com/ozonmp/omp-bot/internal/app/path"
	insurance_life "github.com/ozonmp/omp-bot/internal/service/insurance/life"
	"log"
)

type InsuranceSubdomainCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type InsuranceCommander struct {
	bot           *tgbotapi.BotAPI
	lifeCommander InsuranceSubdomainCommander
}

func NewInsuranceCommander(
	bot *tgbotapi.BotAPI,
) *InsuranceCommander {
	return &InsuranceCommander{
		bot:           bot,
		lifeCommander: life.NewTelegramLifeCommander(bot, insurance_life.NewDummyLifeService()),
	}
}

func (c *InsuranceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "life":
		c.lifeCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *InsuranceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "life":
		c.lifeCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
