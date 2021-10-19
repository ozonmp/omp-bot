package workplace

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Get(inputMessage *tgbotapi.Message) {
	var args = inputMessage.CommandArguments()

	workplaceID, err := strconv.Atoi(args)
	if err != nil {
		description := fmt.Sprintf("Args( %s) processing error", args)
		c.processError(inputMessage.Chat.ID, description, "Args must be number. Ex.: /get__business__workplace 1")
		return
	}

	workplace, err := c.workplaceService.Describe(uint64(workplaceID))
	if err != nil {
		description := fmt.Sprintf("Fail to get workplace with workplaceID %d: %v", workplaceID, err)
		c.processError(inputMessage.Chat.ID, description, "")
		return
	}

	var msg = tgbotapi.NewMessage(inputMessage.Chat.ID, workplace.String())
	c.bot.Send(msg)
}
