package announcement

var allEntities = []Announcement{
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
}

type Announcement struct {
	Author string
	TimePlanned uint64
	Title string
	Description string
	ThumbnailUrl string
}
