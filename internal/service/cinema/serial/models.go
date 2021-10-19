package serial

type Serial struct {
	ID    int
	Title string
	Year  int
}

var allEntities = []Serial{
	{ID: 1, Title: "Serial one", Year: 2001},
	{ID: 2, Title: "Serial two", Year: 2002},
	{ID: 3, Title: "Serial three", Year: 2003},
	{ID: 4, Title: "Serial four", Year: 2004},
	{ID: 5, Title: "Serial five", Year: 2005},
	{ID: 6, Title: "Serial six", Year: 2006},
	{ID: 7, Title: "Serial seven", Year: 2007},
	{ID: 11, Title: "Serial eleven", Year: 2011},
	{ID: 15, Title: "Serial fifteen", Year: 2015},
	{ID: 16, Title: "Serial sixteen", Year: 2016},
	{ID: 17, Title: "Serial seventeen", Year: 2017},
	{ID: 25, Title: "The Simpsons", Year: 1989},
}
