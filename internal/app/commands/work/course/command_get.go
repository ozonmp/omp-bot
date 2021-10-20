package course

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	index, err := strconv.Atoi(args)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Error:\n%s", err))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Get: error sending error - %v", err)
		}
		return
	}
	if index <= 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Error: course ID is not correct")
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Get: error sending error - %v", err)
		}
		return
	}

	course, err := c.courseService.Describe(uint64(index))

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Error:\n%s", err))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Get: error sending error - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, course.String())
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.Get: error sending reply message to chat - %v", err)
	}
}
