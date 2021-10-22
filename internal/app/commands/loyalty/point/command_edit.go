package point

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
	"log"
	"strconv"
	"strings"
)

const editCommandArgsNum = 3

func (c *PointCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	s := strings.Split(args, "|")

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)

	if len(s) == editCommandArgsNum {

		id, err := strconv.Atoi(strings.TrimSpace(s[0]))

		if err != nil {
			log.Printf("fail to conver id %v", err)
			msg = tgbotapi.NewMessage(
				inputMessage.Chat.ID,
				"fail to conver id",
			)
		} else {
			parsedData := loyalty.Point{ Name: s[1], Description: s[2]}

			err = c.pointService.Edit(uint64(id -1), parsedData)

			if err != nil {
				log.Printf("fail to edit entity %v", err)
				msg = tgbotapi.NewMessage(
					inputMessage.Chat.ID,
					"fail to edit entity",
				)
			} else{
				msg.Text = fmt.Sprintf("Entity was updates, entity id: %d", id)
			}
		}
	}else{
		log.Printf("PointCommander.CallbackList: wrong number = %v of arguments in string: %v", len(s), args)
		msg = tgbotapi.NewMessage(
					inputMessage.Chat.ID,
			 		"Wrong args. Please send like this: {id | name | description}")
	}


	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PointCommander.Edit: error sending reply message to chat - %v", err)
	}
}