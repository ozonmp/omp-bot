package task

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
)

type createProductData struct {
	Championat_id uint64 `json:"championatID"`
	Difficulty    uint64 `json:"difficulty"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}

func (c *TaskCommandStruct) New(inputMessage *tgbotapi.Message) {

	messageText := "Create new product - "

	args := inputMessage.CommandArguments()

	var createData createProductData

	err := json.Unmarshal([]byte(args), &createData)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Format command is bad.\nFormat: {\"title\":\"New title product\",\"description\":\"New descriptions\"}")
		c.SendMessage(msg)
		return
	}

	id, err := c.taskService.Create(
		education.Task{
			Championat_id: createData.Championat_id,
			Difficulty:    createData.Difficulty,
			Title:         createData.Title,
			Description:   createData.Description,
		},
	)
	if err == nil {
		messageText += fmt.Sprintf("success. ProductID = %d\n", id)
	} else {
		messageText += fmt.Sprintf("error - %s", err.Error())
		fmt.Println(err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		messageText,
	)

	c.SendMessage(msg)

}
