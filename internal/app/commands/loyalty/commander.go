package loyalty

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/loyalty/certificate"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LoyaltyCommander struct {
	bot                	 *tgbotapi.BotAPI
	certificateCommander Commander
}

func NewLoyaltyCommander(
	bot *tgbotapi.BotAPI,
) *LoyaltyCommander {
	return &LoyaltyCommander{
		bot: bot,
		certificateCommander: certificate.NewLoyaltyCertificateCommander(bot),
	}
}

func (c *LoyaltyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "certificate":
		c.certificateCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LoyaltyCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LoyaltyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "certificate":
		c.certificateCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LoyaltyCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
