package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strconv"
	"strings"
)

func (c *WorkCourseCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	splittedArgs := strings.Split(callbackPath.CallbackData, " ")

	if len(splittedArgs) != 3 {
		log.Printf("WorkCourseCommander.Callbacklist: error in button args (len)")
		return
	}

	cursorInt, err := strconv.Atoi(splittedArgs[1])
	if err != nil {
		log.Printf("WorkCourseCommander.Callbacklist: error in button args (convert arg1)")
		return
	}

	limitInt, err := strconv.Atoi(splittedArgs[2])
	if err != nil {
		log.Printf("WorkCourseCommander.Callbacklist: error in button args (convert arg2)")
		return
	}

	cursor := uint64(cursorInt)
	limit := uint64(limitInt)

	switch splittedArgs[0] {
	case "next":
		cursor += limit
		c.List(callback.Message, cursor, limit)
	case "prev":
		cursor -= limit
		if cursor < 0 {
			cursor = 0
		}
		c.List(callback.Message, cursor, limit)
	}
}
