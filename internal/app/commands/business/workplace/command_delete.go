package workplace

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	workplaceID, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"args must be number. Ex.: /delete__business__workplace 1",
		)
		c.bot.Send(msg)
		return
	}

	deleted, err := c.workplaceService.Remove(uint64(workplaceID))
	if err != nil {
		log.Printf("Fail to delete workplace with ID = %d", workplaceID)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Workplace with ID %d is not found in DB", workplaceID),
		)
		c.bot.Send(msg)
		return
	}

	msgValue := ""
	if !deleted {
		msgValue = fmt.Sprintf("Workplace with ID %d is not deleted from DB", workplaceID)
	} else {
		msgValue = fmt.Sprintf("Workplace with ID %d is deleted from DB", workplaceID)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgValue,
	)

	c.bot.Send(msg)
}
