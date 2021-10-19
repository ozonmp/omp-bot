package recomendation

import "fmt"

type Product struct {
	Id uint64
	Title string
	Description string
	Rating float64
}

func (product *Product) String() string{
	return fmt.Sprintf( "%d %s", product.Id, product.Title)
}