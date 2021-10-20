package film

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaFilmCommander) Get(inputMessage *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMessage.CommandArguments())
	film, err := c.getCommandLogic(entityArgsString)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}
	_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, film.String()))
}

func (c *CinemaFilmCommander) getCommandLogic(entityArgsString string) (*cinema.Film, error) {
	idx, err := strconv.Atoi(entityArgsString)
	if err != nil || idx < 0 {
		return nil, fmt.Errorf("You need to input film ID, it must be a number and can't be lower than 0")
	}

	film, err := c.filmService.Describe(uint64(idx))
	if err != nil {
		return nil, fmt.Errorf("fail to get film: %v", err)
	}

	return film, nil
}
