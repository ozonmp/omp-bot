package group

import "github.com/ozonmp/omp-bot/internal/model/product"

var groups = []product.Group{
	{
		id:    1,
		owner: "Owner 1",
		items: "Item 1",
	},
	{
		id:    1,
		owner: "Owner 1",
		items: "Item 2",
	},
	{
		id:    1,
		owner: "Owner 1",
		items: "Item 3",
	},
	{
		id:    1,
		owner: "Owner 2",
		items: "Item 1",
	},
	{
		id:    1,
		owner: "Owner 2",
		items: "Item 2",
	},
	{
		id:    1,
		owner: "Owner 2",
		items: "Item 3",
	},
}

type Subdomain struct {
	Title string
}
