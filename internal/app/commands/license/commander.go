package license

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/license/author"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LicenseCommander struct {
	bot             *tgbotapi.BotAPI
	authorCommander *author.LicenseAuthorCommander
}

func NewLicenseCommander(
	bot *tgbotapi.BotAPI,
) *LicenseCommander {
	return &LicenseCommander{
		bot:             bot,
		authorCommander: author.NewLicenseAuthorCommander(bot),
	}
}

func (c *LicenseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "author":
		c.authorCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LicenseCommander.HandleCallback: unknown author - %s", callbackPath.Subdomain)
	}
}

func (c *LicenseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "author":
		c.authorCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LicenseCommander.HandleCommand: unknown author - %s", commandPath.Subdomain)
	}
}
