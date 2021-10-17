package task

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type createProductData struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *TaskCommander) Create(inputMessage *tgbotapi.Message) {

	messageText := "Create new product - "

	args := inputMessage.CommandArguments()

	var createData createProductData

	err := json.Unmarshal([]byte(args), &createData)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Format command is bad.\nFormat: {\"id\":5,\"title\":\"New title product\",\"description\":\"New descriptions\"}")
		c.SendMessage(msg)
		return
	}

	err = c.taskService.Create(createData.Id, createData.Title, createData.Description)
	if err == nil {
		messageText += fmt.Sprintf("success. ProductID = %d\n", createData.Id)
	} else {
		messageText += fmt.Sprintf("error - %s", err.Error())
		fmt.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		messageText,
	)

	c.SendMessage(msg)

}
