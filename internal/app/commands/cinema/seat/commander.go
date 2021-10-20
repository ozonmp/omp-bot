package seat

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/cinema/seat"
)

type SeatCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type CinemaSeatCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *seat.DummySeatService
}

func NewCinemaSeatCommander(
	bot *tgbotapi.BotAPI,
) *CinemaSeatCommander {
	seatService := seat.NewDummySeatService()

	return &CinemaSeatCommander{
		bot:              bot,
		subdomainService: seatService,
	}
}

func (c *CinemaSeatCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CinemaSeatCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaSeatCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
