package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"strings"
)

func (c *CinemaFilmCommander) New(inputMsg *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMsg.CommandArguments())
	film, err := c.newCommandLogic(entityArgsString)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}
	c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Film created:\n%s", film.String())))
}

func (c *CinemaFilmCommander) newCommandLogic(entityArgsString string) (*cinema.Film, error) {
	params, err := getParameters(entityArgsString)
	if err != nil {
		return nil, err
	}

	film, err := filmFromParameters(params)
	if err != nil {
		return nil, err
	}

	if _, err := c.filmService.Create(film); err != nil {
		return nil, err
	}
	return film, nil
}
