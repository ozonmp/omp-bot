package film

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaFilmCommander) Get(ctx context.Context, inputMessage *tgbotapi.Message) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - Get - Start")

	entityArgsString := strings.TrimSpace(inputMessage.CommandArguments())
	film, err := c.getCommandLogic(ctx, entityArgsString)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - Get - Error")
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	log.Debug().Err(err).Msg("CinemaFilmCommander - Get - Success")
	_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, film.String()))
}

func (c *CinemaFilmCommander) getCommandLogic(ctx context.Context, entityArgsString string) (*cinema.Film, error) {
	idx, err := strconv.Atoi(entityArgsString)
	if err != nil || idx < 0 {
		return nil, fmt.Errorf("You need to input film ID, it must be a number and can't be lower than 0")
	}

	film, err := c.filmService.Describe(ctx, int64(idx))
	if err != nil {
		return nil, fmt.Errorf("fail to get film: %v", err)
	}

	return film, nil
}
