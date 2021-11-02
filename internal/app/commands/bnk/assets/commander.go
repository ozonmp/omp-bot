package assets

import (
	"github.com/ozonmp/omp-bot/internal/service/bnk/assets"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type AssetsCommander struct {
	bot           *tgbotapi.BotAPI
	assetsService *assets.Service
}

func NewAssetsCommander(
	bot *tgbotapi.BotAPI,
) *AssetsCommander {
	assetsService := assets.NewService()

	return &AssetsCommander{
		bot:           bot,
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
		c.New(msg)
	default:
		c.Default(msg)
	}
}

func (c *AssetsCommander) Send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.Get: error sending reply message to chat - %v", err)
	}
}
