package assets

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AssetsCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	commands := strings.Split(args, "|")

	for _, command := range commands {
		if command == "" {
			c.Send(inputMessage.Chat.ID, "Один из элементов пуст")
			return
		}
	}

	ID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Неверный аргумент %s: %v", commands[0], err))
		return
	}
	username := commands[1]
	money, err := strconv.ParseFloat(commands[2], 64)
	if err != nil || money < 0 {
		c.Send(inputMessage.Chat.ID, "Проблемы со считывание баланса, будет установлен счёт = 0")
		return
	}

	err = c.assetsService.Edit(ID, username, money)
	if err != nil {
		c.Send(inputMessage.Chat.ID, fmt.Sprintf("Ошибки обновления актива: %v", err))
		return
	}

	c.Send(inputMessage.Chat.ID, fmt.Sprintf("Актив %s успешно изменен %d", username, ID))
}
