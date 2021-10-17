package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/streaming/announcement"
	"log"
	"strconv"
)

func (c *StreamingAnnouncementCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Usage: /edit__streaming__announcement {announcement index}",
		)
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	a := announcement.Announcement{}
	if c.announcementService.Update(uint64(idx), a) != nil {
		log.Printf("fail to update announcement with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Failed to update announcement with id: " + strconv.Itoa(idx),
		)

		_, err = c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement was updated successfully",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.Get: error sending reply message to chat - %v", err)
	}
}