package internship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *WorkInternshipCommander) List(inputMessage *tgbotapi.Message) {
	max := len(c.internshipService.List(0, 0))
	outputMsgText := "Interships (id&description), total: " + strconv.Itoa(max) + "\n\n"
	if max == 0 {
		outputMsgText += "empty\n"
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkInternshipCommander.List: error sending reply message to chat - %v", err)
		}
		return
	}
	var internships = c.internshipService.List(c.cursor, c.limit)
	if internships == nil {
		if c.limit > 0 {
			c.cursor = (uint64(max) / c.limit) * c.limit
		} else {
			c.cursor = 0
		}
		internships = c.internshipService.List(c.cursor, c.limit)
	}
	for _, p := range internships {
		outputMsgText += c.internshipService.ShortString(p)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if c.cursor > 0 {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Prev", "work__internship__list__prev"),
		)
	}

	if int(c.cursor+c.limit) < max {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next", "work__internship__list__next"),
		)
	}

	if len(buttons) > 0 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
		msg.ReplyMarkup = keyboard
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.List: error sending reply message to chat - %v", err)
	}

}
