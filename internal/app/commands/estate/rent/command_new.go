package rent

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/estate"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var command_new = func(cmnd CommandName) CommandMetadata {
	newMetadata := newCommandMetadata(cmnd, "create new product")

	newResponseHandler := func(c *Commander, chatId int64, command string, commandArgs string) {
		if rent, err := estate.ParseRent(commandArgs); err != nil {
			c.sendError(chatId, err, newMetadata.Command)
		} else if modelId, err := c.productService.Create(*rent); err != nil {
			c.sendError(chatId, err, newMetadata.Command)
		} else {
			editData := templateFactory(
				fmt.Sprintf("%s %v", CmndEdit.ToDomainCommand(), modelId),
				rent,
			)

			deleteArgs := commandDeleteArgs{ID: modelId}
			deleteArgsData, _ := deleteArgs.ToJsonString()
			deleteData := fmt.Sprintf(`%s__%s`, CmndDelete.ToCallbackPath(), deleteArgsData)

			msgText := fmt.Sprintf(`%s created (ID: %v)
show %s`, rent.ObjectType.Full(), modelId, CmndList.ToDomainCommand())

			msg := tgbotapi.NewMessage(chatId, msgText)
			msg.ParseMode = "HTML"

			appendEditAndDeleteButtons(&msg, *editData, deleteData)

			c.sendWithLog(msg, "")
		}
	}

	newRequestHandler := func(c *Commander, chatId int64, command string) {
		msg := tgbotapi.NewMessage(chatId, `Choose type and fill template by rows:
<code>
ROW  FIELD        TYPE      EXAMPLE
1    object_type  string    car/house
2    object_info  string    Description text
3    renter_id    uint64    1
4    period       duration  4h/3d/2m/1y
5    price        decimal   9.99
</code>
<b>remove first row which start with @bot_name and send message</b>`)
		msg.ParseMode = "HTML"

		appendCreateNewButtons(&msg, newMetadata.Name.ToDomainCommand())
		c.sendWithLog(msg, newMetadata.Command)
	}

	newMetadata.Exec = func(c *Commander, chatId int64, command string, commandArgs string) {
		if len(commandArgs) > 0 {
			newResponseHandler(c, chatId, command, commandArgs)
		} else {
			newRequestHandler(c, chatId, command)
		}
	}

	return newMetadata
}(CmndNew)


