package platform

const (
	PrevButtonText = "Prev page"
	NextButtonText = "Next page"
)

const (
	DefaultListLimit = 3
)

type CallbackListData struct {
	Cursor uint64
	Limit  uint64
}

type PlatformInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SiteUrl     string `json:"site_url"`
	Enabled     bool   `json:"enabled"`
}
