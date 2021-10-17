package streaming

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/streaming/announcement"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type StreamingCommander struct {
	bot                *tgbotapi.BotAPI
	announcementCommander Commander
}

func NewStreamingCommander(
	bot *tgbotapi.BotAPI,
) *StreamingCommander {
	return &StreamingCommander{
		bot: bot,
		announcementCommander: announcement.NewStreamingAnnouncementCommander(bot),
	}
}

func (c *StreamingCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "announcement":
		c.announcementCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("StreamingCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *StreamingCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "announcement":
		c.announcementCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("StreamingCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

