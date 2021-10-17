package production

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationProductionCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		help("list", "")+
			help("create", `{"description":"test","type":"x"}`)+
			help("read", "1")+
			help("update", `{"id":1,"description":"testy","type":"y"}`)+
			help("delete", "1"))

	c.sendMessage(msg)
}

func help(commandName, args string) string {
	return fmt.Sprintf("/%v__recommendation__production %v\n", commandName, args)
}
