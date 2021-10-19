package ground

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	c.Send(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	c.Help(inputMessage)
}

func (c *GroundCommander) Send(chatID int64, text string) {
	if len(text) == 0 {
		log.Printf("Message text empty!!!")
	} else {
		msg := tgbotapi.NewMessage(chatID, text)
		c.send(msg)
	}
}

func (c *GroundCommander) SendWithReply(chatID int64, text string, reply interface{}) {
	if len(text) == 0 {
		log.Printf("Message text empty!!!")
	} else {
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = reply

		c.send(msg)
	}
}

func (c *GroundCommander) send(msg tgbotapi.Chattable) {
	send, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Message with text '%s' don't send! Error: %v", send.Text, err)
	} else {
		log.Println(send.Text)
	}
}
