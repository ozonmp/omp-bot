package championat

type ChampionatEditData struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ChampionatCreateData struct {
	Title string `json:"title"`
}
