package exchange

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/app/commands/exchange/exchange"
	"github.com/ozonmp/omp-bot/internal/app/path"
	exchange2 "github.com/ozonmp/omp-bot/internal/service/exchange/exchange"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ExchangeDomainCommander struct {
	bot               *tgbotapi.BotAPI
	exchangeCommander Commander
}

func NewExchangeDomainCommander(bot *tgbotapi.BotAPI) *ExchangeDomainCommander {

	return &ExchangeDomainCommander{
		bot:               bot,
		exchangeCommander: exchange.NewExchangeCommander(bot, exchange2.NewDummyExchangeService()),
	}
}

func (c *ExchangeDomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "exchange":
		c.exchangeCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ExchangeDomainCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ExchangeDomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "exchange":
		c.exchangeCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ExchangeDomainCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
