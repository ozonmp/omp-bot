package transition

type CreateTransition struct {
	Name string `json:"name"`
	From string `json:"from"`
	To   string `json:"to"`
}
