package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/streaming/announcement"
	"log"
)

type StreamingAnnouncementCommander struct {
	bot              *tgbotapi.BotAPI
	announcementService *announcement.Service
}

func NewStreamingAnnouncementCommander(
	bot *tgbotapi.BotAPI,
) *StreamingAnnouncementCommander {
	announcementService := announcement.NewService()

	return &StreamingAnnouncementCommander{
		bot:              bot,
		announcementService: announcementService,
	}
}

func (c *StreamingAnnouncementCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("StreamingAnnouncementCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *StreamingAnnouncementCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help": c.Help(msg)
	case "list": c.List(msg)
	case "get": c.Get(msg)
	case "new": c.New(msg)
	case "edit": c.Edit(msg)
	case "delete": c.Delete(msg)
	default: c.Default(msg)
	}
}
