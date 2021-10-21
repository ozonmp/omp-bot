package rent

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

var command_delete = func(cmnd CommandName) CommandMetadata {
	deleteMetadata := newCommandMetadata(cmnd, "delete product")
	deleteMetadata.Exec = func (c *Commander, chatId int64, command string, commandArgs string) {
		var modelId uint64
		var err error
		if modelId, err = strconv.ParseUint(commandArgs, 10, 64); err != nil {
			c.sendError(chatId, err, deleteMetadata.Command)
			return
		}

		var msgTxt string

		if ok, err := c.productService.Remove(modelId); err != nil {
			c.sendError(chatId, err, deleteMetadata.Command)
			return
		} else if ok {
			msgTxt = fmt.Sprintf(`❌ rent deleted (ID: %v)
show %s`, modelId, CmndList.ToDomainCommand())
		} else {
			msgTxt = fmt.Sprintf(`❗ rent not found (ID: %v)
show %s`, modelId, CmndList.ToDomainCommand())
		}

		msg := tgbotapi.NewMessage(chatId, msgTxt)
		c.sendWithLog(msg, "CallbackDelete")
	}

	return deleteMetadata
}(CmndDelete)


