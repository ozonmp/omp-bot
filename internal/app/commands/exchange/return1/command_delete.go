package return1

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Return1CommanderImpl) Delete(inputMsg *tgbotapi.Message) {
	//reply := func(text string, other ...interface{}) {
	//for _, v := range other {
	//log.Println("Return1CommanderImpl.Delete:", v)
	//}

	//msg := tgbotapi.NewMessage(
	//inputMsg.Chat.ID,
	//text,
	//)

	//_, err := c.bot.Send(msg)
	//if err != nil {
	//log.Printf("Return1CommanderImpl.Delete: error sending reply message to chat [%v]", err)
	//}
	//}

	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		replyToUser("error wrong arguments. should be 1 number - element id. error[%v]", inputMsg, c.bot, err)

		return
	}

	success, err := c.service.Remove(uint64(idx))
	if err != nil {
		replyToUser(fmt.Sprintf("failed delete with idx[%d]", idx), inputMsg, c.bot, err)

		return
	}

	if success {
		replyToUser(
			fmt.Sprintf("Successfully removed element with id[%d]."+
				"\n\nCarefull now! other element indexes may be changed!",
				idx),
			inputMsg,
			c.bot,
		)
	} else {
		replyToUser(fmt.Sprintf("Can't remove element with id[%d]", idx), inputMsg, c.bot)
	}
}
