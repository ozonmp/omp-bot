package product

import (
	"fmt"
	"strconv"
)

type Product struct {
	id   uint64
	name string
}

func (p *Product) String() string {
	return fmt.Sprintf("id: %d name: %s", p.id, p.name)
}

var AllProducts []Product

func init() {
	for i := 0; i < 30; i++ {
		AllProducts = append(AllProducts, Product{id: uint64(i), name: "Product №" + strconv.Itoa(i)})
	}
}
