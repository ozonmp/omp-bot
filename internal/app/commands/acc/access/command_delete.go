package access

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AccAccessCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var outMsgText string
	st, err := c.accessService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to delete access with idx %d: %v", idx, err)
		outMsgText = "fail to delete access"
	} else {
		if st {
			outMsgText = "Access with ID " + strconv.Itoa(idx) + " deleted"
		} else {
			outMsgText = "Access with ID " + strconv.Itoa(idx) + " not found"
		}

	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMsgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Delete: error sending reply message to chat - %v", err)
	}
}
