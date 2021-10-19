package workplace

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *BusinessWorkplaceCommander) Edit(inputMessage *tgbotapi.Message) {
	var workplace = business.Workplace{}

	if err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &workplace); err != nil {
		var description = fmt.Sprintf("Passed workplace is not valid: %s", inputMessage.CommandArguments())
		c.processError(inputMessage.Chat.ID, description, "Correct workplace example: "+WorkplaceJsonExample)
		return
	}

	if workplace.ID == 0 {
		c.processError(inputMessage.Chat.ID, "ID in workplace data is absent", "Correct workplace example: "+WorkplaceJsonExample)
		return
	}

	if err := c.workplaceService.Update(workplace.ID, workplace); err != nil {
		var description = fmt.Sprintf("Fail to update workplace %d, %v", workplace.ID, err)
		c.processError(inputMessage.Chat.ID, description, "")
		return
	}

	var msg = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Warkplace with ID %d was updated.", workplace.ID))
	c.bot.Send(msg)
}
