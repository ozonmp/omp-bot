package referral

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type DummyService interface {
	List(cursor uint64, limit uint64) ([]Referral, error)
	Count() int
	Describe(idx uint64) (*Referral, error)
	Remove(idx uint64) (bool, error)
	Create(referral Referral) (uint64, error)
	Update(idx uint64, referral Referral) error
}

type ReferralCommander struct {
	bot             *tgbotapi.BotAPI
	referralService *DummyReferralService
}

func NewReferralCommander(bot *tgbotapi.BotAPI) *ReferralCommander {
	referralService := NewDummyDummyReferralService()

	return &ReferralCommander{
		bot:             bot,
		referralService: referralService,
	}
}

func (c *ReferralCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ReferralCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ReferralCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(0, msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		log.Printf("Invalid command %s", commandPath.CommandName)
	}
}
