package championat

var allEntities = []Championat{
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
	allEntities = append(allEntities, Championat{
		Title: Title,
	})
}

func deleteEntity(i int) {
	allEntities = append(allEntities[:i], allEntities[i+1:]...)
}

func editEntity(i int, title string) {
	allEntities[i] = Championat{
		Title: title,
	}
}
