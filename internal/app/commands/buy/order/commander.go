package order

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/buy"
	"github.com/ozonmp/omp-bot/internal/service/buy/order"
)

type OrderService interface {
	Describe(orderID uint64) (*buy.Order, error)
	List(cursor uint64, limit uint64) ([]buy.Order, error)
	Create(buy.Order) (uint64, error)
	Update(orderID uint64, order buy.Order) error
	Remove(orderID uint64) (bool, error)
}

type OrderCommander struct {
	bot          *tgbotapi.BotAPI
	orderService OrderService
}

func NewOrderCommander(bot *tgbotapi.BotAPI) *OrderCommander {
	orderService := order.NewDummyOrderService()

	return &OrderCommander{
		bot:          bot,
		orderService: orderService,
	}
}

func (c *OrderCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("OrderCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *OrderCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *OrderCommander) Reply(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OrderCommander: error sending reply message to chat - %v", err)
	}
}
