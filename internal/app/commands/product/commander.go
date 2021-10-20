package product

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/product/group"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ProductCommander struct {
	bot            *tgbotapi.BotAPI
	groupCommander group.GroupCommander
}

func NewProductCommander(bot *tgbotapi.BotAPI) *ProductCommander {
	return &ProductCommander{
		bot:            bot,
		groupCommander: group.NewProductGroupCommander(bot),
	}
}

func (c *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "group":
		c.groupCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ProductCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "group":
		c.groupCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Commander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
