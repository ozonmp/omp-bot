package task

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskStruct) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	pos, err := strconv.Atoi(args)
	if err != nil || pos == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "ProductID not correct. id > 0")
		c.SendMessage(msg)
		return
	}

	product, err := c.taskService.Describe(uint64(pos))

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf("Get error. Error = %s", err))
		c.SendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.String())

	c.SendMessage(msg)
}
