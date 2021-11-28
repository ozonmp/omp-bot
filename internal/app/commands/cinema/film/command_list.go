package film

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/paginator"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	"strings"
)

const defaultPageLength int64 = 10
const firstIndexFilm int64 = 1

func (c *CinemaFilmCommander) List(ctx context.Context, inputMessage *tgbotapi.Message, p *paginator.Paginator) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - List - Start")

	if p == nil {
		p = paginator.NewPaginator("")
	} else if p.Direction == "next" {
		p.Page += 1
	} else {
		p.Page -= 1
	}

	startIndex := p.Page*defaultPageLength + 1
	if startIndex < 0 {
		startIndex = firstIndexFilm
		p.Page = 0
	}

	strFilms, outOfBound, err := c.getFilmsToPage(ctx, startIndex, defaultPageLength)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - List - Get films error")
		return
	}
	if outOfBound {
		p.Page = 0
	}
	if len(strFilms) == 0 {
		_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, "You dont have films. Add at least one"))
		return
	}

	outputMsgText := "Films: \n\n"

	outputMsgText = strings.Join(strFilms, "\n")
	outputMsgText += fmt.Sprintf("\npage %d", p.Page+1)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	msg.ReplyMarkup = p.NewKeyBoard()
	_ = c.sendMessage(msg)
	p = nil
}

func (c *CinemaFilmCommander) getFilmsToPage(ctx context.Context, cursor int64, limit int64) ([]string, bool, error) {
	log := logger.LoggerFromContext(ctx)
	log.Debug().Msg("CinemaFilmCommander - getFilmsToPage - Start")

	films, err := c.filmService.List(ctx, cursor, limit)
	log.Debug().Msg(fmt.Sprintf("CinemaFilmCommander - getFilmsToPage - Got films %v", films))
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - getFilmsToPage - Error")
		return nil, false, err
	}

	if len(films) != 0 {
		return c.fromFilmsToStrings(films), false, nil
	}

	log.Debug().Err(err).Msg("CinemaFilmCommander - getFilmsToPage - Second chance to get films")

	films, err = c.filmService.List(ctx, firstIndexFilm, limit)
	if err != nil {
		log.Debug().Err(err).Msg("CinemaFilmCommander - getFilmsToPage - Second chance Error")
		return nil, true, err
	}

	return c.fromFilmsToStrings(films), true, nil
}

func (c *CinemaFilmCommander) fromFilmsToStrings(films []cinema.Film) []string {
	result := make([]string, 0, len(films))
	for _, film := range films {
		result = append(result, fmt.Sprintf("%d|'%s'", film.ID, film.Name))
	}
	return result
}
