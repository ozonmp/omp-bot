package task

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskStruct) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	task_id, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Remove error - id is not valid")
		c.SendMessage(msg)
		return
	}

	_, err = c.taskService.Remove(uint64(task_id))
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf("Remove error. Error = %s", err),
		)
		c.SendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Remove product - success")

	c.SendMessage(msg)

}
