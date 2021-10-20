package return1

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Return1CommanderImpl) List(inputMsg *tgbotapi.Message) {
	reply := func(text string, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.List:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.List: error sending reply message to chat [%v]", err)
		}
	}

	return1Elements, err := c.service.List(0, 0) //tobedone
	if err != nil {
		reply("error during getting list from service", err)

		return
	}

	textResponse := "Here all the elements:\n\n"
	for _, r := range return1Elements {
		textResponse += r.String() + "\n"
	}

	reply(textResponse)
}
