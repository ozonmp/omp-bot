package championat

var allChampionats = []Championat{
	{Title: "ZeroChampionat"},
	{Title: "FirstChampionat"},
	{Title: "SecondChampionat"},
	{Title: "ThirdChampionat"},
	{Title: "FourthChampionat"},
	{Title: "FifthChampionat"},
}

type Championat struct {
	Title string
}

func newEntity(Title string) {
	allChampionats = append(allChampionats, Championat{
		Title: Title,
	})
}

func deleteEntity(i uint64) {
	allChampionats = append(allChampionats[:i], allChampionats[i+1:]...)
}

func editEntity(i uint64, title string) {
	allChampionats[i] = Championat{
		Title: title,
	}
}
