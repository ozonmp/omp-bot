package product

type JsonProductModel struct {
	Id          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}
