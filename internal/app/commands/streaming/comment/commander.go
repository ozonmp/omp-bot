package comment

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/streaming/comment"
)

const commentsPerPage = 5

type CommentCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type StreamingCommentCommander struct {
	bot            *tgbotapi.BotAPI
	commentService comment.CommentService
}

func NewStreamingCommentCommander(bot *tgbotapi.BotAPI) CommentCommander {
	commentService := comment.NewDummyCommentService()

	return &StreamingCommentCommander{
		bot:            bot,
		commentService: commentService,
	}
}

func (c *StreamingCommentCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("StreamingCommentCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *StreamingCommentCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		//c.Default(msg)
	}
}
