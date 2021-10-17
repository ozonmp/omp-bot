package office

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

type EditRequest struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *OfficeCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	parsedData := EditRequest{}

	err := json.Unmarshal([]byte(args), &parsedData)

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error reading json data for type Office from "+
			"input string %v - %v", args, err)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			`Wrong json struct. Please send like this: {"id": 1, name":"name", "description":"description"}`,
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("OfficeCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)

	err = c.officeService.Update(parsedData.Id, business.Office{
		Name:        parsedData.Name,
		Description: parsedData.Description,
	})

	if err != nil {
		log.Printf("fail to edit entity %v", err)
		msg.Text = err.Error()
	} else {
		msg.Text = fmt.Sprintf("Entity was updated, id:%d", parsedData.Id)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Edit: error sending reply message to chat - %v", err)
	}
}
