package customer

var allEntities = []Customer{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Customer struct {
	Title string
}
