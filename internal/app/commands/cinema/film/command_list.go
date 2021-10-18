package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strings"
)

const pageLength = 10

var currentPage = 0

var callbackPathNext = path.CallbackPath{
	Domain:       "cinema",
	Subdomain:    "film",
	CallbackName: "list",
	CallbackData: "next",
}

var callbackPathPrev = path.CallbackPath{
	Domain:       "cinema",
	Subdomain:    "film",
	CallbackName: "list",
	CallbackData: "prev",
}

var keyBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Next", callbackPathNext.String()),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Prev", callbackPathPrev.String()),
	),
)

func (c *CinemaFilmCommander) List(inputMessage *tgbotapi.Message, next, init bool) {
	answerText := "You dont have films. Add at least one"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, answerText)
	msg.ReplyMarkup = keyBoard
	if len(c.filmService.Films) == 0 {
		_ = c.sendMessage(msg)
		return
	}

	outputMsgText := "Films: \n\n"

	if init {
		currentPage = 0
	} else if next {
		currentPage += 1
	} else {
		currentPage -= 1
	}
	maxPage := (len(c.filmService.Films)-1)/pageLength + 1
	if currentPage == maxPage && next {
		currentPage = 0
	} else if currentPage == -1 && !next {
		currentPage = maxPage - 1
	}

	startIndex := currentPage * pageLength
	endIndex := (currentPage + 1) * pageLength

	filmsToOutput, _ := c.filmService.List(uint64(startIndex), uint64(endIndex), true)
	filmsOnPage := c.fromFilmsToStrings(filmsToOutput)
	outputMsgText = strings.Join(filmsOnPage, "\n")
	outputMsgText += fmt.Sprintf("\npage %d/%d", currentPage+1, maxPage)

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
