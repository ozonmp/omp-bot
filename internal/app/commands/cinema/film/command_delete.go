package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (c *CinemaFilmCommander) Delete(inputMessage *tgbotapi.Message) {
	args := strings.TrimSpace(inputMessage.CommandArguments())
	id, err := strconv.Atoi(args)
	if err != nil || id < 0 {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
			"You need to input film ID, it must be a number and can't be lower than 0"))
		return
	}

	film, err := c.filmService.Describe(uint64(id))
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	filmNameToDelete := film.Name
	_, err = c.filmService.Remove(film.ID)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("Film '%s' succesfully deleted", filmNameToDelete)))
}

