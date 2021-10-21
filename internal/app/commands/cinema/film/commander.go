package film

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/cinema/film"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CinemaFilmCommander struct {
	bot         *tgbotapi.BotAPI
	filmService *film.DummyFilmService
}

func NewCinemaFilmCommander(bot *tgbotapi.BotAPI) *CinemaFilmCommander {
	filmService := film.NewDummyFilmService()
	return &CinemaFilmCommander{
		bot:         bot,
		filmService: filmService,
	}
}

func (c *CinemaFilmCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CinemaFilmCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaFilmCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg, nil)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *CinemaFilmCommander) sendMessage(msg tgbotapi.MessageConfig) error {
	_, err := c.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("CinemaFilmCommander.List: error sending reply message to chat - %v", err)
	}
	return err
}
