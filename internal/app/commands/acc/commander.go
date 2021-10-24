package acc

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/acc/access"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type AccCommander struct {
	bot             *tgbotapi.BotAPI
	accessCommander Commander
}

func NewAccCommander(
	bot *tgbotapi.BotAPI,
) *AccCommander {
	return &AccCommander{
		bot:             bot,
		accessCommander: access.NewAccAccessCommander(bot),
	}
}

func (c *AccCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "access":
		c.accessCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("AccCommander.HandleCallback: unknown Access - %s", callbackPath.Subdomain)
	}
}

func (c *AccCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "access":
		c.accessCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("AccCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
