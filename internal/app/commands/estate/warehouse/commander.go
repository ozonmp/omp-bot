package warehouse

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/estate/warehouse"
)

type EstateWarehouseCommander struct {
	bot              *tgbotapi.BotAPI
	warehouseService *warehouse.DummyEstateService
}

func NewEstateWarehouseCommander(bot *tgbotapi.BotAPI) *EstateWarehouseCommander {
	dws := warehouse.NewDummyEstateService()
	return &EstateWarehouseCommander{
		bot:              bot,
		warehouseService: dws,
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
		c.Describe(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *EstateWarehouseCommander) sendMsg(chatID int64, msg string) {
	m := tgbotapi.NewMessage(chatID, msg)
	_, err := c.bot.Send(m)
	if err != nil {
		log.Printf("Estate.WarehouseCommander.sendInternal: error sending reply message to chat - %v", err)
	}
}
