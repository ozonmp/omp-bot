package workplace

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

func (c *BusinessWorkplaceCommander) New(inputMessage *tgbotapi.Message) {

	workplace := business.Workplace{}
	if err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &workplace); err != nil {
		log.Printf("fail to unmarshall input workplace  %v", workplace)
		exampleWorkplace, _ := (json.Marshal(business.Workplace{1, "Title Example"}))
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("New workplace is not valid. Example:  "+string(exampleWorkplace)),
		)
		c.bot.Send(msg)
		return
	}

	newWorkplaceID, err := c.workplaceService.Create(workplace)
	if err != nil {
		log.Printf("fail to create new workplace %v", err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("fail to create new workplace %v", err),
		)
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("New warkplace was created. ID - %d", newWorkplaceID),
	)

	c.bot.Send(msg)
}
