package film

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

func (c CinemaFilmCommander) Edit(inputMsg *tgbotapi.Message) {
	entityArgsString := strings.TrimSpace(inputMsg.CommandArguments())
	params, err := getParameters(entityArgsString)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	strIDToFind := ""
	for _, arg := range params {
		if arg[0] == "search" { // search:film name to find
			strIDToFind = arg[1]
			break
		}
	}

	idToFind, err := strconv.Atoi(strIDToFind)
	if err != nil || idToFind < 0 {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, "Message must contain 'search' argument with film ID, to find a film for update.\nExample: search:10;name:Harry Potter;rating:9.8;description:good story"))
		return
	}

	newFilm, err := filmFromParameters(params)
	if err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	if err := c.filmService.Update(uint64(idToFind), newFilm); err != nil {
		c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("%s", err)))
		return
	}

	c.sendMessage(tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Updated film:\n%s", newFilm)))
}

