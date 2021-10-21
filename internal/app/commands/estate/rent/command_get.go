package rent

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"regexp"
	"strconv"
)

var command_get = func(cmnd CommandName) CommandMetadata {
	sampleRegexp := regexp.MustCompile(`\d+`)
	cmdMetadata := newCommandMetadata(cmnd, "get product by id")
	cmdMetadata.Exec = func(c *Commander, chatId int64, command string, commandArgs string) {
		if len(commandArgs) == 0 {
			commandArgs = sampleRegexp.FindString(command)
		}

		withArguments := func(id string) {
			var modelId uint64
			var err error
			if modelId, err = strconv.ParseUint(id, 10, 64); err != nil {
				c.sendError(chatId, err, cmdMetadata.Command)
				return
			}

			if rent, err := c.productService.Describe(modelId); err == nil {
				editData := templateFactory(
					fmt.Sprintf("%s %v", CmndEdit.ToDomainCommand(), modelId),
					rent,
				)

				deleteArgs := commandDeleteArgs{ID: modelId}
				deleteArgsData, _ := deleteArgs.ToJsonString()
				deleteData := fmt.Sprintf(`%s__%s`, CmndDelete.ToCallbackPath(), deleteArgsData)

				msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("%s (ID %v):\n%s\n%s",
					rent.ObjectType.Full(), modelId, rent.ToFormatRowsString(), CmndList.ToDomainCommand()))
				msg.ParseMode = "HTML"

				appendEditAndDeleteButtons(&msg, *editData, deleteData)

				c.sendWithLog(msg, cmdMetadata.Command)
			} else {
				c.sendError(chatId, err, cmdMetadata.Command)
			}
		}

		noArguments := func() {
			msg := tgbotapi.NewMessage(chatId, fmt.Sprintf(`No ID specified! Command format is:
<code>
%s 123
</code>
or
<code>
%s123
</code>
try %s
or 
show %s`,
				cmdMetadata.Name.ToDomainCommand(),
				cmdMetadata.Name.ToDomainCommand(),
				CmndHelp.ToDomainCommand(),
				CmndList.ToDomainCommand()),
			)
			msg.ParseMode = "HTML"
			c.sendWithLog(msg, cmdMetadata.Command)
		}

		if len(commandArgs) == 0 {
			noArguments()
		} else {
			withArguments(commandArgs)
		}
	}

	return cmdMetadata
}(CmndGet)
