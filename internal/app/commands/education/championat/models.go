package championat

type ChampionatEditData struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type ChampionatCreateData struct {
	Title string `json:"title"`
}
