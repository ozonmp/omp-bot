package cinema

import (
	"context"
	cnmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/film"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot           *tgbotapi.BotAPI
	filmCommander Commander
}

func NewCinemaCommander(
	ctx context.Context,
	bot *tgbotapi.BotAPI,
	film *cnmApi.CnmFilmApiServiceClient,
) *CinemaCommander {
	return &CinemaCommander{
		bot: bot,
		// subdomainCommander
		filmCommander: film.NewCinemaFilmCommander(ctx, bot, film),
	}
}

func (c *CinemaCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "film":
		c.filmCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown film - %s", callbackPath.Subdomain)
	}
}

func (c *CinemaCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "film":
		c.filmCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown film - %s", commandPath.Subdomain)
	}
}
