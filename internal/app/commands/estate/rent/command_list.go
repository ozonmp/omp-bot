package rent

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var command_list = func(cmnd CommandName) CommandMetadata {
	argsFactory := func (page uint64, commandPath string) string {
		prevArgs := commandListArgs{Page: page}
		prevJson, _ := json.Marshal(prevArgs)
		return fmt.Sprintf("%s__%s", commandPath, string(prevJson))
	}

	cmdMetadata := newCommandMetadata(cmnd, "get product list")
	cmdMetadata.Exec = func (c *Commander, chatId int64, command string, commandArgs string) {
		args := commandListArgs{}
		_ = args.Unmarshal(commandArgs)

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Here all the rents (page: %v): \n", args.Page + 1))

		if products, err := c.productService.List(args.Page*c.offset, c.offset); err == nil {
			productsCount := uint64(len(products))

			for _, p := range products {
				sb.WriteString(fmt.Sprintf("\n%s%v %s", CmndGet.ToDomainCommand(), p.ID, p.ObjectType.Full()))
				sb.WriteString(p.ToFormatRowsString())
			}

			if productsCount == 0 {
				sb.WriteString("\n<b>no entities found</b>")
			}

			msgText := sb.String()
			msg := tgbotapi.NewMessage(chatId, msgText)
			msg.ParseMode = "HTML"

			var nextArgs string
			if productsCount >= c.offset {
				nextArgs = argsFactory(args.Page + uint64(1), cmdMetadata.Name.ToCallbackPath())
			}

			var prevArgs string
			if args.Page > 0 {
				prevArgs = argsFactory(args.Page - uint64(1), cmdMetadata.Name.ToCallbackPath())
			}

			appendPaginationButtons(&msg, prevArgs, nextArgs)

			c.sendWithLog(msg, "Command")
		} else {
			c.sendError(chatId, err, command)
		}
	}
	return cmdMetadata
}(CmndList)


