package announcement

import "time"

var allEntities = []Announcement{
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample 1",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 2",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample 3",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 4",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample 5",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 6",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample 7",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 8",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "John Doe",
		TimePlanned: 1634488911,
		Title: "Sample 9",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 10",
		Description: "Sample description",
		ThumbnailUrl: "example.com",
	},
	{
		Author: "Jane Doe",
		TimePlanned: 1634488911,
		Title: "Sample 11",
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

func (a *Announcement) FormattedTime() string {
	return time.Unix(int64(a.TimePlanned), 0).String()
}

func (a *Announcement) String() string {
	return "Author: " + a.Author + "\n" +
		"Time: " + a.FormattedTime() + "\n" +
		"Title: " + a.Title + "\n" +
		"Description: " + a.Description + "\n" +
		"Thubnail: " + a.ThumbnailUrl
}