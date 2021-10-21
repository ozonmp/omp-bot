package transition

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityTransitionCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__activity__transition - help\n"+
			"/list__activity__transition [offset limit] - list transitions\n"+
			"/get__activity__transition transitionId - get transition with transition id}\n"+
			"/delete__activity__transition transitionId - delete transition with transition id\n"+
			"/create__activity__transition json - create the new transition from json string, example {\"name\":\"test\",\"from\":\"xxx\",\"to\":\"yyy\"}\n"+
			"/update__activity__transition transitionId json - update the existings transition with transition id from json string, example {\"name\":\"test2\",\"from\":\"aaa\",\"to\":\"bbb\"}\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Help: error sending reply message to chat - %v", err)
	}
}
