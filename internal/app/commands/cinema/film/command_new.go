package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func (c *CinemaFilmCommander) New(inputMsg *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMsg.CommandArguments())
	params, err := getParameters(entityArgsString)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	film, err := filmFromParameters(params)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	if _, err := c.filmService.Create(film); err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Film created:\n%s", film.String())))
}
