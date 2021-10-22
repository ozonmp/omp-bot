package course

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *WorkCourseCommander) Help(inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__course 		- this help\n"+
			"/list__work__course 		- list of courses\n"+
			"/get__work__course id 		- get course by id\n"+
			"/delete__work__course id 	- delete course by id\n"+
			"/edit__work__course	 	- edit course info by id\n"+
			"example: 3 Edited Edited Edited Edited\n"+
			"/new__work__course 		- add new course\n"+
			"example: New New New New",
	)

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Printf("WorkCourseCommander.Help: error sending reply message to chat - %v", err)
	}
}
