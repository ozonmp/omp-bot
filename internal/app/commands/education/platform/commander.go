package platform

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/platform"
)

type PlatformCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	Create(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type PlatformCallbackHandler interface {
	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type EducationPlatformCommander struct {
	bot                     *tgbotapi.BotAPI
	platformCommander       PlatformCommander
	platformCallbackHandler PlatformCallbackHandler
}

func NewEducationPlatformCommander(
	bot *tgbotapi.BotAPI,
) *EducationPlatformCommander {
	return &EducationPlatformCommander{
		bot:                     bot,
		platformCommander:       NewPlatformCommander(bot, platform.NewDummyPlatformService()),
		platformCallbackHandler: NewPlatformCallbackHandler(bot, platform.NewDummyPlatformService()),
	}
}

func (c *EducationPlatformCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.platformCallbackHandler.CallbackList(callback, callbackPath)
	default:
		log.Printf("EducationPlatformCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EducationPlatformCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.platformCommander.Help(msg)
	case "list":
		c.platformCommander.List(msg)
	case "get":
		c.platformCommander.Get(msg)
	case "delete":
		c.platformCommander.Delete(msg)
	case "new":
		c.platformCommander.Create(msg)
	case "edit":
		c.platformCommander.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *EducationPlatformCommander) sendMessage(msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EducationPlatformCommander: error sending reply message to chat - %v", err)
	}
}
