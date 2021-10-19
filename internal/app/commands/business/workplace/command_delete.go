package workplace

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Delete(inputMessage *tgbotapi.Message) {
	var args = inputMessage.CommandArguments()

	workplaceID, err := strconv.Atoi(args)
	if err != nil {
		description := fmt.Sprintf("Args( %s) processing error", args)
		c.processError(inputMessage.Chat.ID, description, "Args must be number. Ex.: /delete__business__workplace 1")
		return
	}

	deleted, err := c.workplaceService.Remove(uint64(workplaceID))
	if err != nil {
		description := fmt.Sprintf("Fail to delete workplace with ID = %d", workplaceID)
		c.processError(inputMessage.Chat.ID, description, "")
		return
	}

	var msgValue = fmt.Sprintf("Workplace with ID %d is deleted from DB", workplaceID)
	if !deleted {
		msgValue = fmt.Sprintf("Workplace with ID %d is not deleted from DB", workplaceID)
	}

	var msg = tgbotapi.NewMessage(inputMessage.Chat.ID, msgValue)
	c.bot.Send(msg)
}
