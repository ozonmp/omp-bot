package exchange

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/exchange/return1"
	"github.com/ozonmp/omp-bot/internal/app/path"
	return1Service "github.com/ozonmp/omp-bot/internal/service/exchange/return1"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ExchangeCommander struct {
	bot              *tgbotapi.BotAPI
	return1Commander Commander
}

func NewExchangeCommander(
	bot *tgbotapi.BotAPI,
) *ExchangeCommander {
	return &ExchangeCommander{
		bot:              bot,
		return1Commander: return1.NewReturn1Commander(bot, *return1Service.NewDummyReturn1Service()),
	}
}

func (c *ExchangeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "return1":
		c.return1Commander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ExchangeCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ExchangeCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "return1":
		c.return1Commander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ExchangeCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
