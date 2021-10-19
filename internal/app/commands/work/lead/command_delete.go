package lead

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LeadCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		c.sendMsg(inputMessage.Chat.ID, "Command format: /delete__work__lead ID")
		return
	}

	_, err = c.leadService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to remove lead with idx %d: %v", idx, err)
		c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf("Error: %v", err))
		return
	}

	c.sendMsg(inputMessage.Chat.ID, "Successfully deleted")

}
