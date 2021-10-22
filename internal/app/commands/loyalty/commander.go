package loyalty

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/loyalty/referral"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LoyaltyCommander struct {
	bot               *tgbotapi.BotAPI
	referralCommander *referral.ReferralCommander
}

func NewLoyaltyCommanderCommander(
	bot *tgbotapi.BotAPI,
) *LoyaltyCommander {
	referralCommander := referral.NewReferralCommander(bot)

	return &LoyaltyCommander{
		bot:               bot,
		referralCommander: referralCommander,
	}
}

func (c *LoyaltyCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "referral":
		c.referralCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LoyaltyCommander.HandleCommand: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LoyaltyCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "referral":
		c.referralCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LoyaltyCommander.HandleCommand unknown subdomain %s", commandPath.Subdomain)
	}
}
