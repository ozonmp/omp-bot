package task

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskStruct) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	pos, err := strconv.Atoi(args)
	if err != nil || pos == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "id not correct. id > 0")
		c.SendMessage(msg)
		return
	}

	product, _ := c.taskService.Describe(uint64(pos))

	if product.IsEmpty() {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "id not found.")
		c.SendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.String())

	c.SendMessage(msg)
}
