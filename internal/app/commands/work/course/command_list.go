package course

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/work"

	"log"
)

func (c *WorkCourseCommander) List(inputMessage *tgbotapi.Message, cursor uint64, limit uint64) {
	courses, err := c.courseService.List(cursor, limit)
	if err != nil {
		log.Printf("WorkCourseCommander.List: error - %v", err)
	}

	text := ""
	for _, v := range courses {
		text += v.String()
		text += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)

	if cursor > 0 {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Prev", fmt.Sprintf("work__course__list__prev %d %d", cursor, limit)),
		)
	}
	if int(cursor+limit) < len(work.AllCourses) {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next", fmt.Sprintf("work__course__list__next %d %d", cursor, limit)),
		)
	}
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.List: error sending reply message to chat - %v", err)
	}

}
