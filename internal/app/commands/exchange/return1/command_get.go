package return1

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Return1CommanderImpl) Get(inputMsg *tgbotapi.Message) {
	//reply := func(text string, other ...interface{}) {
	//for _, v := range other {
	//log.Println("Return1CommanderImpl.Get:", v)
	//}

	//msg := tgbotapi.NewMessage(
	//inputMsg.Chat.ID,
	//text,
	//)

	//_, err := c.bot.Send(msg)
	//if err != nil {
	//log.Printf("Return1CommanderImpl.Get: error sending reply message to chat [%v]", err)
	//}
	//}

	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		replyToUser("error wrong arguments. should be 1 number - element id.", inputMsg, c.bot, err)

		return
	}

	element, err := c.service.Describe(uint64(idx))
	if err != nil {
		replyToUser(fmt.Sprintf("failed get with idx[%d]", idx), inputMsg, c.bot, err)

		return
	}

	replyToUser("Successfully got element: "+element.String(), inputMsg, c.bot)
}
