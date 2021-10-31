package intern

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	intern2 "github.com/ozonmp/omp-bot/internal/service/work/intern"
	"log"
)

func (c *WorkInternCommander) New(inputMessage *tgbotapi.Message) {
	name := inputMessage.CommandArguments()

	if len(name) == 0 {
		log.Printf("Empty name of student")
		return
	}

	newIntern := intern2.NewIntern(name)
	internID, err := c.internService.Create(*newIntern)
	if err != nil {
		log.Printf("fail to create new intern with name %s", name)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Intern %s was successfully created and assigned with id=%d", name, internID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternCommander.Get: error sending reply message to chat - %v", err)
	}
}
