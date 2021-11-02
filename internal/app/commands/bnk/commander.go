package demo

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/bnk/assets"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type AssetsCommander struct {
	bot             *tgbotapi.BotAPI
	assetsCommander Commander
}

func NewAssetsCommander(
	bot *tgbotapi.BotAPI,
) *AssetsCommander {
	return &AssetsCommander{
		bot: bot,
		// assetsCommander
		assetsCommander: assets.NewAssetsCommander(bot),
	}
}

func (c *AssetsCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "assets":
		c.assetsCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("AssetsCommander.HandleCallback: unknown assets - %s", callbackPath.Subdomain)
	}
}

func (c *AssetsCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "assets":
		c.assetsCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("AssetsCommander.HandleCommand: unknown assets - %s", commandPath.Subdomain)
	}
}
