package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/paginator"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strings"
)

func (c *CinemaFilmCommander) List(inputMessage *tgbotapi.Message, next, init bool) {
	tempPaginator, ok := c.paginators[inputMessage.Chat.ID]
	if !ok {
		c.paginators[inputMessage.Chat.ID] = paginator.NewCinemaPaginator()
		tempPaginator = c.paginators[inputMessage.Chat.ID]
	}

	answerText := "You dont have films. Add at least one"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, answerText)
	msg.ReplyMarkup = tempPaginator.Keyboard
	if len(c.filmService.Films) == 0 {
		_ = c.sendMessage(msg)
		return
	}

	outputMsgText := "Films: \n\n"

	if init {
		tempPaginator.CurrentPage = 0
	} else if next {
		tempPaginator.CurrentPage += 1
	} else {
		tempPaginator.CurrentPage -= 1
	}
	maxPage := (len(c.filmService.Films)-1)/tempPaginator.PageLength + 1
	if tempPaginator.CurrentPage == maxPage && next {
		tempPaginator.CurrentPage = 0
	} else if tempPaginator.CurrentPage == -1 && !next {
		tempPaginator.CurrentPage = maxPage - 1
	}

	startIndex := tempPaginator.CurrentPage * tempPaginator.PageLength

	filmsToOutput, _ := c.filmService.List(uint64(startIndex), uint64(tempPaginator.PageLength))
	filmsOnPage := c.fromFilmsToStrings(filmsToOutput)
	outputMsgText = strings.Join(filmsOnPage, "\n")
	outputMsgText += fmt.Sprintf("\npage %d/%d", tempPaginator.CurrentPage+1, maxPage)

	msg.Text = outputMsgText
	_ = c.sendMessage(msg)
}

func (c *CinemaFilmCommander) fromFilmsToStrings(films []cinema.Film) []string {
	result := make([]string, 0, len(films))
	for _, film := range films {
		result = append(result, fmt.Sprintf("%d|'%s'", film.ID, film.Name))
	}
	return result
}
