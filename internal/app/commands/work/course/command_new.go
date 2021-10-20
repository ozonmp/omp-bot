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

	splittedArgs := strings.Split(args, " ")

	id, err := c.courseService.Create(
		work.Course{
			//Id:          uint64(courseId),
			Name:        splittedArgs[1],
			Timing:      splittedArgs[2],
			Description: splittedArgs[3],
			Level:       splittedArgs[4],
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
