package course

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/work"
	"log"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	args=strings.TrimSpace(args)
	splittedArgs := strings.Split(args, " ")

	if len(splittedArgs) != 4 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("WorkCourseCommander.New: wrong number of arguments"))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.New: error in sending error: wrong number of arguments")
		}
		return
	}

	id, err := c.courseService.Create(
		work.Course{
			//Id:          uint64(courseId),
			Name:        splittedArgs[0],
			Timing:      splittedArgs[1],
			Description: splittedArgs[2],
			Level:       splittedArgs[3],
		})

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("WorkCourseCommander.New: error in courseService.Create - %v", err))
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("WorkCourseCommander.New: error in courseService.Create - %v\nAnd error in sending error reply", err)
		}
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("New course added with id: %v", id))
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkCourseCommander.New: error sending reply - %v", err)
	}
}
