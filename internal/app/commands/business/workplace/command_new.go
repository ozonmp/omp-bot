package workplace

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *BusinessWorkplaceCommander) New(inputMessage *tgbotapi.Message) {

	var workplace = business.Workplace{}

	if err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &workplace); err != nil {
		var description = fmt.Sprintf("Fail to unmarshal input data: %s ", inputMessage.CommandArguments())
		c.processError(inputMessage.Chat.ID, description, "Correct workplace example: "+WorkplaceJsonExample)
		return
	}

	if workplace.ID == 0 {
		c.processError(inputMessage.Chat.ID, "ID in workplace data is absent", "Correct workplace example: "+WorkplaceJsonExample)
		return
	}

	newWorkplaceID, err := c.workplaceService.Create(workplace)
	if err != nil {
		var description = fmt.Sprintf("Fail to create new workplace %v", err)
		c.processError(inputMessage.Chat.ID, description, "")
		return
	}

	var msg = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("New warkplace was created. ID - %d", newWorkplaceID))
	c.bot.Send(msg)
}
