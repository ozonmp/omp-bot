package ground

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Send(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	// msg.ParseMode = "markdown"

	c.send(msg)
}

func (c *GroundCommander) SendWithReply(chatID int64, text string, reply interface{}) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = reply
	// msg.ParseMode = "markdown"

	c.send(msg)
}

func (c *GroundCommander) send(msg tgbotapi.Chattable) {
	send, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Message with text '%s' don't send!", send.Text)
		log.Panic(err)
	} else {
		log.Println(send.Text)
	}
}

func (c *GroundCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	c.Send(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	c.Help(inputMessage)
}
