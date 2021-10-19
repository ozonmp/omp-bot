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

	Help(inputMessage *tgbotapi.Message)
	Get(inputMessage *tgbotapi.Message)
	List(inputMessage *tgbotapi.Message)
	Delete(inputMessage *tgbotapi.Message)
	New(inputMessage *tgbotapi.Message)
	Edit(inputMessage *tgbotapi.Message)
}

type EstateCommander struct {
	bot                *tgbotapi.BotAPI
	warehouseCommander Commander
}

func NewEstateCommander(
	bot *tgbotapi.BotAPI,
) *EstateCommander {
	return &EstateCommander{
		bot: bot,
		// warehouseCommander
		warehouseCommander: warehouse.NewEstateWarehouseCommander(bot),
	}
}

func (c *EstateCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Warehouse {
	case "warehouse":
		c.warehouseCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("EstateCommander.HandleCallback: unknown warehouse - %s", callbackPath.Warehouse)
	}
}

func (c *EstateCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Warehouse {
	case "warehouse":
		c.warehouseCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("EstateCommander.HandleCommand: unknown warehouse - %s", commandPath.Warehouse)
	}
}
