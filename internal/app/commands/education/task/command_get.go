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

	outputMsgText := fmt.Sprintf("ProductID: %d Name: %s Description: %s", product.Id, product.Title, product.Description)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	c.SendMessage(msg)
}
