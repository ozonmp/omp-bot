package office

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

func (c *OfficeCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	parsedData := business.Office{}

	err := json.Unmarshal([]byte(args), &parsedData)

	if err != nil {
		log.Printf("OfficeCommander.CallbackList: "+
			"error reading json data for type Office from "+
			"input string %v - %v", args, err)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong json struct. Please send like this: {\"name\":\"name\", \"description\":\"description\"}",
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("OfficeCommander.Create: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)

	id, err := c.officeService.Create(parsedData)

	if err != nil {
		log.Printf("fail to create entity %v", err)
		msg.Text = err.Error()
	} else {
		msg.Text = fmt.Sprintf("Entity was added, id:%d", id)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Create: error sending reply message to chat - %v", err)
	}
}
