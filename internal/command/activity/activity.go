package activity

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/command/activity/click"
	service "github.com/ozonmp/omp-bot/internal/service/activity/click"
	"log"
)

type ActivityCommander struct {
	bot            *tgbotapi.BotAPI
	clickCommander click.ClickCommander
}

func NewActivityCommander(bot *tgbotapi.BotAPI) *ActivityCommander {
	s := service.NewActivityClickService()

	return &ActivityCommander{
		bot:            bot,
		clickCommander: click.NewActivityClickCommander(bot, s),
	}
}

func (c *ActivityCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "click":
		c.clickCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ActivityCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ActivityCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "click":
		c.clickCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ActivityCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
