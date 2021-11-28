package film

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	"strconv"
	"strings"
)

func (c *CinemaFilmCommander) Delete(ctx context.Context, inputMessage *tgbotapi.Message) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - Delete - Start")

	entityArgsString := strings.TrimSpace(inputMessage.CommandArguments())
	filmToDelete, err := c.deleteCommandLogic(ctx, entityArgsString)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - Delete - Error")
		_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	log.Debug().Msg("CinemaFilmCommander - Delete - Success")
	_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("Film '%s' succesfully deleted", *filmToDelete)))
}

func (c *CinemaFilmCommander) deleteCommandLogic(ctx context.Context, entityArgsString string) (*string, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - deleteCommandLogic - Start")

	id, err := strconv.Atoi(entityArgsString)
	if err != nil || id < 0 {
		err := fmt.Errorf("you need to input film ID, it must be a number and can't be lower than 0")
		log.Debug().Err(err).Msg("CinemaFilmCommander - deleteCommandLogic - Cast error")
		return nil, err
	}

	film, err := c.filmService.Describe(ctx, int64(id))
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - deleteCommandLogic - Describe error")
		return nil, err
	}

	filmToDelete := film.String()
	_, err = c.filmService.Remove(ctx, film.ID)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - deleteCommandLogic - Remove error")
		return nil, err
	}

	log.Debug().Msg("CinemaFilmCommander - deleteCommandLogic - Success")
	return &filmToDelete, nil
}
