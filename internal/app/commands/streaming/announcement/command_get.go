package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *StreamingAnnouncementCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Usage: /get__streaming__announcement {announcement index}",
			)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	a, err := c.announcementService.Describe(uint64(idx))
	if err != nil || a == nil {
		log.Printf("fail to get announcement with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Failed to get announcement with id: " + strconv.Itoa(idx),
		)

		_, err = c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		a.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
	}
}
