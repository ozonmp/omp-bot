package education

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/test"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type EducationCommander struct {
	bot           *tgbotapi.BotAPI
	testCommander Commander
}

func NewEducationCommander(
	bot *tgbotapi.BotAPI,
) *EducationCommander {
	return &EducationCommander{
		bot:           bot,
		testCommander: test.NewEducationTestCommander(bot),
	}
}

func (c *EducationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "test":
		c.testCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("EducationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *EducationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "test":
		c.testCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("EducationCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
