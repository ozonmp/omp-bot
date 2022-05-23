package subject

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/subject"
)

type SubjectService interface {
	Describe(subjectID uint64) (*subject.Subject, error)
	List(cursor uint64, limit uint64) ([]subject.Subject, error)
	Create(subject.Subject) (uint64, error)
	Update(subjectID uint64, subject subject.Subject) error
	Remove(subjectID uint64) (bool, error)
	SubjectsCount() uint64
}

type SubjectCommander struct {
	bot            *tgbotapi.BotAPI
	subjectService SubjectService
}

func NewSubjectCommander(bot *tgbotapi.BotAPI) *SubjectCommander {
	service := subject.NewDummyService()
	return &SubjectCommander{
		bot:            bot,
		subjectService: service,
	}
}

func (commander *SubjectCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		commander.CallbackList(callback, callbackPath)
	default:
		log.Printf("EducationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (commander *SubjectCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		commander.Help(msg)
	case "get":
		commander.Get(msg)
	case "new":
		commander.New(msg)
	case "edit":
		commander.Edit(msg)
	case "delete":
		commander.Delete(msg)
	case "list":
		commander.List(msg)
	default:
		commander.Default(msg)
	}
}
