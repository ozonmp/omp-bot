package course

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/work"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	args=strings.TrimSpace(args)
	splittedArgs  := strings.Split(args, " ")

	if len(splittedArgs) != 5 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("WorkCourseCommander.Edit: wrong number of arguments"))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Edit: error in sending reply %v", err)
		}
		return
	}

	id, err := strconv.Atoi(splittedArgs[0])
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("WorkCourseCommander.Edit: wrong id - %v", err))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.Edit: error in sending reply %v", err)
		}
		return
	}
	err = c.courseService.Update(
		work.Course{
			Id:          uint64(id),
			Name:        splittedArgs[1],
			Timing:      splittedArgs[2],
			Description: splittedArgs[3],
			Level:       splittedArgs[4],
		})

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("WorkCourseCommander.Edit: error in courseService.Update - %v", err))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.New: error in courseService.Edit - %v\nAnd error in sending error reply", err)
		}
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Edited course with id: %v", id))
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.Edit: error sending reply - %v", err)
	}

}
