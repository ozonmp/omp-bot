package estate

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/estate/warehouse"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type EstateCommander struct {
	bot                *tgbotapi.BotAPI
	warehouseCommander Commander
}

func NewEstateCommander(
	bot *tgbotapi.BotAPI,
) *EstateCommander {
	return &EstateCommander{
		bot:                bot,
		warehouseCommander: warehouse.NewEstateWarehouseCommander(bot),
	}
}

func (c *EstateCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "warehouse":
		c.warehouseCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("EstateCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *EstateCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "warehouse":
		c.warehouseCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("EstateCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
