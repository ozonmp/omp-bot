package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/streaming/announcement"
	"log"
	"strconv"
	"strings"
)

type AnouncementData struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TimePlanned  uint64 `json:"time_planned"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

func (c *StreamingAnnouncementCommander) New(inputMessage *tgbotapi.Message) {
	args := strings.Trim(inputMessage.CommandArguments(), "")

	if args == "" {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Command format is: /new__streaming__announcement {announcement json}\n" +
				"JSON fields are:\nauthor(string),\n" +
				"title(string),\n" +
				"description(string),\n" +
				"time_planned(timestamp),\n" +
				"thumbnail_url(string)",
		)
		c.bot.Send(msg)
		return
	}

	parsedData := AnouncementData{}
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.New: "+
			"error reading json data for type AnouncementData from "+
			"input string %v - %v", args, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Command format is: /new__streaming__announcement {announcement json}\n" +
				"JSON fields are:\nauthor(string),\n" +
				"title(string),\n" +
				"description(string),\n" +
				"time_planned(timestamp),\n" +
				"thumbnail_url(string)",
		)
		c.bot.Send(msg)
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
