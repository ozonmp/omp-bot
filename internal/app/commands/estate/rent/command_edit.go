package rent

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/estate"
	"strconv"
	"strings"
)

var command_edit = func(cmnd CommandName) CommandMetadata {

	cmdMetadata := newCommandMetadata(cmnd, "change product")

	cmdMetadata.Exec = func(c *Commander, chatId int64, commander string, commandArgs string) {
		var modelId uint64
		var err error

		parts := strings.Split(commandArgs, "\n")

		if len(parts) <= 1 {
			err = fmt.Errorf("Edit.Command: error nothing to set")
			c.sendError(chatId, err, cmdMetadata.Command)
			return
		}

		if modelId, err = strconv.ParseUint(parts[0], 10, 64); err != nil {
			c.sendError(chatId, err, cmdMetadata.Command)
			return
		}

		var rent *estate.Rent
		if rent, err = c.productService.Describe(modelId); err != nil {
			c.sendError(chatId, err, cmdMetadata.Command)
			return
		}

		if err = rent.Patch(parts[1:]); err != nil {
			c.sendError(chatId, err, cmdMetadata.Command)
			return
		}

		if err := c.productService.Update(modelId, *rent); err != nil {
			c.sendError(chatId, err, cmdMetadata.Command)
			return
		}

		editData := templateFactory(
			fmt.Sprintf("%s %v", CmndEdit.ToDomainCommand(), modelId),
			rent,
		)

		deleteArgs := commandDeleteArgs{ID: modelId}
		deleteArgsData, _ := deleteArgs.ToJsonString()
		deleteData := fmt.Sprintf(`%s__%s`, CmndDelete.ToDomainCommand(), deleteArgsData)

		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(`%s edited (ID: %v)
check %s
or
show %s`, rent.ObjectType.Full(), rent.ID, CmndGet.ToDomainCommand(), CmndList.ToDomainCommand()))

		appendEditAndDeleteButtons(&msg, *editData, deleteData)

		c.sendWithLog(msg, cmdMetadata.Command)
	}

	return cmdMetadata
}(CmndEdit)
