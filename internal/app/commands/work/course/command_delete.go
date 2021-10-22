package course

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	args=strings.TrimSpace(args)
	index, err := strconv.Atoi(args)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Error: wrong ID\n"))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Delete: error sending error - %v", err)
		}
		return
	}

	result, err := c.courseService.Remove(uint64(index))

	if err != nil {
		log.Printf("WorkCourseCommander.Delete: error while deleting - %v", err)
	}
	var text string
	if result {
		text = "Course info was deleted"
	} else {
		text = "Course id does not exist"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.Delete: error sending reply message to chat - %v", err)
	}
}
