package announcement

const pageLimit = 5

type AnouncementData struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TimePlanned  uint64 `json:"time_planned"`
	ThumbnailUrl string `json:"thumbnail_url"`
}