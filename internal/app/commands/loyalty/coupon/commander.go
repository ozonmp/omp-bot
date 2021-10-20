package coupon

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/loyalty/coupon"
)

type LoyaltyCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type LoyaltyCouponCommander struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewLoyaltyCouponCommander(bot *tgbotapi.BotAPI) *LoyaltyCouponCommander {
	couponService := service.NewService()

	return &LoyaltyCouponCommander{
		bot:           bot,
		service: couponService,
	}
}

func (c *LoyaltyCouponCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LoyaltyCouponCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LoyaltyCouponCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
