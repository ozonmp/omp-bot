package workplace

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"
)

func (c *BusinessWorkplaceCommander) Edit(inputMessage *tgbotapi.Message) {
	workplace := business.Workplace{}
	if err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &workplace); err != nil {
		log.Printf("fail to unmarshall input workplace  %v", workplace)
		exampleWorkplace, _ := (json.Marshal(business.Workplace{1, "Title Example"}))
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Workplace is not valid. Example:  "+string(exampleWorkplace)),
		)
		c.bot.Send(msg)
		return
	}

	err := c.workplaceService.Update(workplace.ID, workplace)
	if err != nil {
		log.Printf("fail to update workplace %v", err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("fail to update workplace %d, %v", workplace.ID, err),
		)
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Warkplace with ID %d was updated.", workplace.ID),
	)

	c.bot.Send(msg)
}
