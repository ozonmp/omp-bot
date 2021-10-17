package subdomain

import "fmt"

var (
	CommandHelp   = "help__domain__subdomain"
	CommandGet    = "get__domain__subdomain"
	CommandList   = "list__domain__subdomain"
	CommandDelete = "delete__domain__subdomain"
	CommandNew    = "new__domain__subdomain"
	CommandEdit   = "edit__domain__subdomain"

	UsageGet  = fmt.Sprintf("/%s <subdomain_id>\n\nsubdomain_id - id элемента (строго больше 0)", CommandGet)
	UsageList = fmt.Sprintf(
		"/%s <cursor> <limit>\n\n"+
			"cursor - курсор в базе (строго больше 0)\n"+
			"limit - кол-во элементов в странице (строго больше 0)", CommandList)
	UsageDelete = fmt.Sprintf("/%s <subdomain_id>\n\nsubdomain_id - id элемента (строго больше 0)", CommandDelete)
	UsageNew    = fmt.Sprintf("/%s <subdomain_json>\n\nExample: `{\"name\": \"Batman\"}`", CommandNew)
	UsageEdit   = fmt.Sprintf(
		"/%s <subdomain_json>\n\n"+
			"Edits by specified id\n"+
			"Example: `{\"id\": 1, \"name\": \"Batman\"}`", CommandEdit)

	ErrNotFound  = "Не найдено"
	ErrOnDelete  = "Не удалось удалить"
	ErrOnEdit    = "Не удалось изменить"
	ErrEmptyList = "Пустой список 😢. Добавьте элементы."

	SuccessDelete = "Успешно удалено!"
	SuccessNew    = "Успешно добавлено!"
	SuccessEdit   = "Успешно изменено!"

	CallbackNameList = "list"
)
