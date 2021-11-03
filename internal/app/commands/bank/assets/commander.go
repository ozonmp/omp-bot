package assets

import (
	"github.com/ozonmp/omp-bot/internal/service/bank/assets"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type AssetsCommander struct {
	Bot           *tgbotapi.BotAPI
	assetsService *assets.DummyAssetsService
}

func NewAssetsCommander(
	bot *tgbotapi.BotAPI,
) *AssetsCommander {
	assetsService := assets.NewDummyAssetsService()

	return &AssetsCommander{
		Bot:           bot,
		assetsService: assetsService,
	}
}

func (c *AssetsCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("AssetsCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *AssetsCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *AssetsCommander) Send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err := c.Bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.Get: error sending reply message to chat - %v", err)
	}
}
