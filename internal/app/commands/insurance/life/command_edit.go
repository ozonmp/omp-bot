package life

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/insurance"
	"strconv"
	"strings"
)

const editErrorMassage = "incorrect input\n\n" + "method signature: /edit__insurance__life {uint} {life in JSON}"

func (telegramLifeCommander *TelegramLifeCommander) Edit(inputMessage *tgbotapi.Message) {

	args := strings.Split(inputMessage.CommandArguments(), " ")

	if len(args) != 2 {
		telegramLifeCommander.sendError(inputMessage, editErrorMassage)
		return
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil || idx < 0 {
		telegramLifeCommander.sendError(inputMessage, editErrorMassage)
		return
	}

	var life insurance.Life
	err = json.Unmarshal([]byte(args[1]), &life)
	if err != nil {
		telegramLifeCommander.sendError(inputMessage, editErrorMassage)
		return
	}

	err = telegramLifeCommander.lifeService.Update(uint64(idx), life)
	if err == nil {
		telegramLifeCommander.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Success"))
	} else {
		telegramLifeCommander.sendError(inputMessage, err.Error())
	}
}
