package warehouse

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type EstateWarehouseCommander struct {
	bot              *tgbotapi.BotAPI
	warehouseService *warehouse.Service
}

func EstateWarehouseCommander(
	bot *tgbotapi.BotAPI,
) *EstateWarehouseCommander {
	warehouseService := warehouse.NewService()

	return &EstateWarehouseCommander{
		bot:              bot,
		warehouseService: warehouseService,
	}
}

func (c *EstateWarehouseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("EstateWarehouseCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EstateWarehouseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
