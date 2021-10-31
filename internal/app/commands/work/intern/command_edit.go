package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	intern2 "github.com/ozonmp/omp-bot/internal/service/work/intern"
	"log"
	"strconv"
	"strings"
)

func (c *WorkInternCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.Fields(inputMessage.CommandArguments())

	if len(args) < 2 {
		log.Printf("wrong args")
		return
	}

	internId, err := strconv.Atoi(args[0])
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	name := strings.Join(args[1:], " ")
	newIntern := intern2.NewIntern(name)
	newIntern.InternshipID = uint64(internId)
	err = c.internService.Update(uint64(internId), *newIntern)

	if err != nil {
		log.Printf("fail to update intern with id %d and name %s", internId, name)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Intern with id %s was successfully updated", internId),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}
