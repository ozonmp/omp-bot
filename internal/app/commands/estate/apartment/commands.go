package apartment

type CommandName string

const (
	CNHelp   CommandName = "help"
	CNList   CommandName = "list"
	CNGet    CommandName = "get"
	CNDelete CommandName = "delete"
	CNNew    CommandName = "new"
	CNEdit   CommandName = "edit"
)

type CallbackName string

const (
	CbNList CallbackName = CallbackName(CNList)
)
