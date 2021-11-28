package cinema

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/film"
	"github.com/ozonmp/omp-bot/internal/app/path"
	cnmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
	"log"
)

type Commander interface {
	HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(ctx context.Context, message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot           *tgbotapi.BotAPI
	filmCommander Commander
}

func NewCinemaCommander(
	bot *tgbotapi.BotAPI,
	filmApi cnmApi.CnmFilmApiServiceClient,
) *CinemaCommander {
	return &CinemaCommander{
		bot: bot,
		// subdomainCommander
		filmCommander: film.NewCinemaFilmCommander(bot, filmApi),
	}
}

func (c *CinemaCommander) HandleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "film":
		c.filmCommander.HandleCallback(ctx, callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown film - %s", callbackPath.Subdomain)
	}
}

func (c *CinemaCommander) HandleCommand(ctx context.Context, msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "film":
		c.filmCommander.HandleCommand(ctx, msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown film - %s", commandPath.Subdomain)
	}
}
