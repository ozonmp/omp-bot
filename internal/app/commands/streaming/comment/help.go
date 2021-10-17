package comment

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *StreamingCommentCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		`
		/help__streaming__comment - help
		/get__streaming__comment 'id' - get comment by 'id'
		/list__streaming__comment - list of comments (5 per page) 
		/delete__streaming__comment 'id' - delete comment by 'id'
		/new__streaming__comment 'text' - create comment with 'text'
		/edit__streaming__comment 'id' 'text' - edit comment set 'text' where 'id'
		`,
	)

	c.bot.Send(msg)
}
