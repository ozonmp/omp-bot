package ticket

const (
	CallbackListPrefix = "travel__ticket__list__"
)

const (
	PreviousPageText = "Prev page"
	NextPageText     = "Next page"
)

const (
	ListLimit = 5
)

type CallbackListData struct {
	Cursor uint64
	Limit  uint64
}
