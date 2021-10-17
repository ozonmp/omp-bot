package task

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type updateProductData struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *TaskCommander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	var updateData updateProductData

	err := json.Unmarshal([]byte(args), &updateData)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Format command is bad.\nFormat: {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}")
		c.SendMessage(msg)
		return
	}

	err = c.taskService.Update(updateData.Id, updateData.Title, updateData.Description)
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
