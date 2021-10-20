package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (c *CinemaFilmCommander) Delete(inputMessage *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMessage.CommandArguments())
	filmToDelete, err := c.deleteCommandLogic(entityArgsString)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("Film '%s' succesfully deleted", *filmToDelete)))
}

func (c *CinemaFilmCommander) deleteCommandLogic(entityArgsString string) (*string, error) {
	id, err := strconv.Atoi(entityArgsString)
	if err != nil || id < 0 {
		return nil, fmt.Errorf("You need to input film ID, it must be a number and can't be lower than 0")
	}

	film, err := c.filmService.Describe(uint64(id))
	if err != nil {
		return nil, err
	}

	filmToDelete := film.String()
	_, err = c.filmService.Remove(film.ID)
	if err != nil {
		return nil, err
	}

	return &filmToDelete, nil
}
