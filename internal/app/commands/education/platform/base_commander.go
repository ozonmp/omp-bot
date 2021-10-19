package platform

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
	"net/url"
)

type PlatformService interface {
	Describe(platformID uint64) (*education.Platform, error)
	List(cursor uint64, limit uint64) ([]education.Platform, error)
	Create(platform education.Platform) (uint64, error)
	Update(platformID uint64, platform education.Platform) error
	Remove(platformID uint64) (bool, error)
}

type PlatformBaseCommander struct {
	bot     *tgbotapi.BotAPI
	service PlatformService
}

type PlatformBaseCallbackHandler struct {
	bot     *tgbotapi.BotAPI
	service PlatformService
}

func NewPlatformCommander(
	bot *tgbotapi.BotAPI,
	service PlatformService,
) PlatformCommander {
	return &PlatformBaseCommander{
		bot:     bot,
		service: service,
	}
}

func NewPlatformCallbackHandler(
	bot *tgbotapi.BotAPI,
	service PlatformService,
) *PlatformBaseCallbackHandler {
	return &PlatformBaseCallbackHandler{bot: bot, service: service}
}

func (c *PlatformBaseCommander) validate(p PlatformInput) error {
	if len(p.Title) < 2 || len(p.Title) > 255 {
		return errors.New("title must be between 2-255 chars")
	}

	_, err := url.ParseRequestURI(p.SiteUrl)
	if err != nil {
		return errors.New("invalid site_url")
	}

	return nil
}

func (c *PlatformBaseCommander) convertToEntity(p PlatformInput) education.Platform {
	return education.Platform{
		Title:       p.Title,
		Description: p.Description,
		SiteUrl:     p.SiteUrl,
		Enabled:     p.Enabled,
	}
}

func (c *PlatformBaseCommander) sendMessage(msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PlatformBaseCommander: error sending reply message to chat - %v", err)
	}
}
