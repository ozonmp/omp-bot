package workplace

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	workplaceID, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"args must be number. Ex.: /get__business__workplace 1",
		)
		c.bot.Send(msg)
		return
	}

	workplace, err := c.workplaceService.Describe(uint64(workplaceID))
	if err != nil {
		log.Printf("fail to get workplace with workplaceID %d: %v", workplaceID, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Workplace with ID %d is not found", workplaceID),
		)
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		workplace.String(),
	)

	c.bot.Send(msg)
}
