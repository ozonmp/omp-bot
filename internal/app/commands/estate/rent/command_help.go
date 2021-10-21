package rent

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var command_help = func(cmnd CommandName) CommandMetadata {
	var helpMessageText string

	cmdMetadata := newCommandMetadata(cmnd, "print commands")
	cmdMetadata.Exec = func (c *Commander, chatId int64, command string, commandArgs string) {
		if len(helpMessageText) == 0 {
			helpMessageText = func(c *Commander) string {
				availableCommands := []CommandMetadata{
					cmdMetadata,
					command_get,
					command_new,
					command_list,
					command_edit,
					command_delete}
				sb := strings.Builder{}
				for i, command := range availableCommands {
					if i > 0 {
						sb.WriteString("\n")
					}
					sb.WriteString(command.Name.ToDomainCommand())
					sb.WriteString(" - ")
					sb.WriteString(command.Description)
				}

				return sb.String()
			}(c)
		}
		msg := tgbotapi.NewMessage(chatId, helpMessageText)
		c.sendWithLog(msg, cmdMetadata.Command)
	}

	return cmdMetadata
}(CmndHelp)
