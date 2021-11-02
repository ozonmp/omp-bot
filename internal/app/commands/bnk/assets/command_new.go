package assets

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commands := strings.Split(args, "|")

	for _, command := range commands {
		if command == "" {
			c.Send(inputMessage.Chat.ID, "Один из элементов пуст")
			return
		}
	}

	username := commands[0]
	money, err := strconv.ParseFloat(commands[1], 64)
	if err != nil || money < 0 {
		c.Send(inputMessage.Chat.ID, "Проблемы со считывание баланса, будет установлен счёт = 0")
		return
	}

	newID, err := c.assetsService.New(username, money)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Ошибки создания актива: %v", err))
		return
	}

	c.Send(inputMessage.Chat.ID, fmt.Sprintf("Актив %s успешно добавлен %d", username, newID))
}
