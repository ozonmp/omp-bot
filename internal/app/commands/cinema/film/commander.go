package film

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/service/cinema/film"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	cnmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
)

type CinemaFilmCommander struct {
	bot         *tgbotapi.BotAPI
	filmService *film.DummyFilmService
}

func NewCinemaFilmCommander(bot *tgbotapi.BotAPI, filmApi cnmApi.CnmFilmApiServiceClient) *CinemaFilmCommander {
	filmService := film.NewDummyFilmService(filmApi)
	return &CinemaFilmCommander{
		bot:         bot,
		filmService: filmService,
	}
}

func (c *CinemaFilmCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(ctx, callback, callbackPath)
	default:
		log.Printf("CinemaFilmCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CinemaFilmCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(ctx, msg, nil)
	case "get":
		c.Get(ctx, msg)
	case "delete":
		c.Delete(ctx, msg)
	case "new":
		c.New(ctx, msg)
	case "edit":
		c.Edit(ctx, msg)
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
