package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/streaming/announcement"
	"log"
	"strconv"
	"strings"
)

func (c *StreamingAnnouncementCommander) New(inputMessage *tgbotapi.Message) {
	args := strings.Trim(inputMessage.CommandArguments(), "")

	if args == "" {
		c.sendNewFormatMessage(inputMessage)
		return
	}

	parsedData := AnouncementData{}
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.New: "+
			"error reading json data for type AnouncementData from "+
			"input string %v - %v", args, err)
		c.sendNewFormatMessage(inputMessage)
		return
	}
	newItem := announcement.Announcement{
		Author:       parsedData.Author,
		TimePlanned:  parsedData.TimePlanned,
		Title:        parsedData.Title,
		Description:  parsedData.Description,
		ThumbnailUrl: parsedData.ThumbnailUrl,
	}

	created, err := c.announcementService.Create(newItem)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Unable to create new announcement",
		)
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Announcement with id " + strconv.FormatUint(created, 10) + " created successfully",
	)
	c.bot.Send(msg)
}

func (c *StreamingAnnouncementCommander) sendNewFormatMessage(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Command format is: /new__streaming__announcement {announcement json}\n" +
			"JSON fields are:\nauthor(string),\n" +
			"title(string),\n" +
			"description(string),\n" +
			"time_planned(timestamp),\n" +
			"thumbnail_url(string)",
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.New: error sending reply message to chat - %v", err)
	}
}