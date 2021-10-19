package life

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

const getErrorMassage = "incorrect input\n\n" + "method signature: /get__insurance__life {uint}"

func (telegramLifeCommander *TelegramLifeCommander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil || idx < 0 {
		telegramLifeCommander.sendError(inputMessage, getErrorMassage)
		return
	}

	if life, err := telegramLifeCommander.lifeService.Describe(uint64(idx)); err == nil {
		lifeJson, _ := json.Marshal(life)
		telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, string(lifeJson)))
	} else {
		telegramLifeCommander.sendError(inputMessage, err.Error())
	}
}
