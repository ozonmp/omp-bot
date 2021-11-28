package film

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	"strings"
)

func (c *CinemaFilmCommander) New(ctx context.Context, inputMsg *tgbotapi.Message) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - New - Start")

	entityArgsString := strings.TrimSpace(inputMsg.CommandArguments())
	film, err := c.newCommandLogic(ctx, entityArgsString)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - New - Error")
		_ = c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	log.Debug().Msg("CinemaFilmCommander - New - Success")
	_ = c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Film created:\n%s", film.String())))
}

func (c *CinemaFilmCommander) newCommandLogic(ctx context.Context, entityArgsString string) (*cinema.Film, error) {
	params, err := getParameters(entityArgsString)
	if err != nil {
		return nil, err
	}

	film, err := filmFromParameters(params)
	if err != nil {
		return nil, err
	}

	if _, err := c.filmService.Create(ctx, film); err != nil {
		return nil, err
	}

	return film, nil
}
