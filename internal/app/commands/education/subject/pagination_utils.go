package subject

import (
	"encoding/json"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/subject"
)

const (
	DefaultSubjectPerPage = 3
)

type CursorData struct {
	Cursor uint64 `json:"offset"`
}

type ListService interface {
	List(cursor uint64, limit uint64) ([]subject.Subject, error)
	SubjectsCount() uint64
}

func getPaginatedMessage(listService ListService, cursor uint64, limit uint64) (string, *tgbotapi.InlineKeyboardMarkup, error) {
	subjects, err := listService.List(cursor, limit)
	if err != nil {
		return "", nil, err
	}

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	if cursor > 0 {
		var newCursor uint64
		if limit > cursor {
			newCursor = 0
		} else {
			newCursor = cursor - limit
		}
		button, err := makeButton("К предыдущей странице", newCursor)
		if err != nil {
			return "", nil, err
		}
		buttons = append(buttons, *button)
	}

	if cursor+limit < listService.SubjectsCount() {
		newCursor := cursor + limit
		button, err := makeButton("К слудующей странице", newCursor)
		if err != nil {
			return "", nil, err
		}
		buttons = append(buttons, *button)
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)
	return formatSubjects(subjects), &keyboard, nil
}

func makeButton(buttonText string, buttonCursor uint64) (*tgbotapi.InlineKeyboardButton, error) {
	offsetData := CursorData{
		Cursor: buttonCursor,
	}

	serCursorData, err := json.Marshal(offsetData)
	if err != nil {
		return nil, err
	}

	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "subject",
		CallbackName: "list",
		CallbackData: string(serCursorData),
	}
	button := tgbotapi.NewInlineKeyboardButtonData(buttonText, callbackPath.String())
	return &button, nil
}

func formatSubjects(subjects []subject.Subject) string {
	if len(subjects) == 0 {
		return "Ни одного элемента"
	}
	var res strings.Builder
	for _, curSubject := range subjects {
		res.WriteString(curSubject.String() + ";\n")
	}
	return res.String()
}
