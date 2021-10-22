package point

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
	"log"
	"strings"
)

func (c *PointCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	parsedData := loyalty.Point{}

	s := strings.Split(args, "|")

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"",
	)

	if len(s) == 2 {
		parsedData = loyalty.Point{Name: s[0], Description: s[1]}

		id, err := c.pointService.Create(parsedData)

		if err != nil {
			log.Printf("fail to create entity %v", err)
			msg.Text = err.Error()
		} else {
			msg.Text = fmt.Sprintf("Entity was added, entity id: %d", (id+1))
		}
	}else{
		log.Printf("PointCommander.CallbackList: wrong number = %v of arguments in string: %v", len(s), args)
		msg = tgbotapi.NewMessage(
					inputMessage.Chat.ID,
			 		"Wrong args. Please send like this: {name | description}")
	}
	

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PointCommander.Create: error sending reply message to chat - %v", err)
	}
}