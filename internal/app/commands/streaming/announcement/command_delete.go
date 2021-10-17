package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *StreamingAnnouncementCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Usage: /delete__streaming__announcement {announcement index}",
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	a, err := c.announcementService.Remove(uint64(idx))
	if err != nil || !a {
		log.Printf("fail to remove announcement with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Failed to remove announcement with id: " + strconv.Itoa(idx),
		)

		_, err = c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement was removed successfully",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
	}
}