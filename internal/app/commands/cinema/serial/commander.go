package serial

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/cinema/serial"
)

type CinemaSerialCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *serial.Service
}

func NewCinemaSerialCommander(bot *tgbotapi.BotAPI) *CinemaSerialCommander {
	subdomainService := serial.NewService()
	return &CinemaSerialCommander{
		bot:              bot,
		subdomainService: subdomainService,
	}
}

func (c *CinemaSerialCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CinemaSerialCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaSerialCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "edit":
		c.Edit(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
