package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/delivery/common"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/delivery/common"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DeliveryCommander struct {
	bot             *tgbotapi.BotAPI
	commonCommander common.DummyCommonCommander
}

func NewDeliveryCommander(
	bot *tgbotapi.BotAPI,
) *DeliveryCommander {
	return &DeliveryCommander{
		bot:             bot,
		commonCommander: common.NewCommonCommander(bot, service.NewDummyCommonService()),
	}
}

func (c *DeliveryCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "common":
		c.commonCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DeliveryCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *DeliveryCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "common":
		c.commonCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DeliveryCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
