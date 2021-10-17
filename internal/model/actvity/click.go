package actvity

type Click struct {
	Title string
}

func (c *Click) String() string {
	return c.Title
}
