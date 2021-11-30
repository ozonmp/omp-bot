package task

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TaskCommandStruct) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	task := make(map[string]interface{})

	err := json.Unmarshal([]byte(args), &task)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Format command is bad.\nFormat: {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}")
		c.SendMessage(msg)
		return
	}

	err = c.taskService.Update(task)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			fmt.Sprintf("Update error. Error = %s", err),
		)
		c.SendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Update product - success")

	c.SendMessage(msg)

}
