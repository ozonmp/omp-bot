package film

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaFilmCommander) Get(inputMessage *tgbotapi.Message) {

	args := strings.TrimSpace(inputMessage.CommandArguments())
	idx, err := strconv.Atoi(args)
	if err != nil || idx < 0 {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
			"You need to input film ID, it must be a number and can't be lower than 0"))
		return
	}

	film, err := c.filmService.Describe(uint64(idx))
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf("fail to get film: %v", err)))
		return
	}

	_ = c.sendMessage(tgbotapi.NewMessage(inputMessage.Chat.ID, film.String()))
}
