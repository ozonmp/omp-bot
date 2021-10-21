package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/paginator"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strings"
)

const defaultPageLength = 10

func (c *CinemaFilmCommander) List(inputMessage *tgbotapi.Message, p *paginator.Paginator) {
	if len(c.filmService.Films) == 0 {
		_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, "You dont have films. Add at least one"))
		return
	}

	outputMsgText := "Films: \n\n"

	if p == nil {
		p = paginator.NewPaginator("")
	} else if p.Direction == "next" {
		p.Page += 1
	} else {
		p.Page -= 1
	}
	maxPage := (c.filmService.NumberOfFilms())/defaultPageLength + 1
	if p.Page >= maxPage && p.Direction == "next" {
		p.Page = 0
	} else if p.Page == -1 && p.Direction == "prev" {
		p.Page = maxPage - 1
	}

	startIndex := p.Page * defaultPageLength

	filmsToOutput, _ := c.filmService.List(uint64(startIndex), uint64(defaultPageLength))
	filmsOnPage := c.fromFilmsToStrings(filmsToOutput)
	outputMsgText = strings.Join(filmsOnPage, "\n")
	outputMsgText += fmt.Sprintf("\npage %d/%d", p.Page+1, maxPage)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	msg.ReplyMarkup = p.NewKeyBoard()
	_ = c.sendMessage(msg)
	p = nil
}

func (c *CinemaFilmCommander) fromFilmsToStrings(films []cinema.Film) []string {
	result := make([]string, 0, len(films))
	for _, film := range films {
		result = append(result, fmt.Sprintf("%d|'%s'", film.ID, film.Name))
	}
	return result
}
