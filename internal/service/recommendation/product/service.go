package product

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
)

type Service interface {
	Describe(id uint64) (*recomendation.Product, error)
	List(cursor uint64, limit uint64) ([]recomendation.Product, error)
	Create(recomendation.Product) (uint64, error)
	Update(id uint64, product recomendation.Product) error
	Remove(id uint64) (bool, error)
}

type DummyProductService struct {
	products map[uint64]recomendation.Product
}

func NewDummyProductService() *DummyProductService {
	products := make(map[uint64]recomendation.Product)
	return &DummyProductService{products: products}
}

func (service *DummyProductService) Describe(id uint64) (*recomendation.Product, error) {
	if product, err := service.products[id]; err {
		return nil, fmt.Errorf("DummyProductService.Describe: product with id %d not found", id)
	} else{
		return &product, nil
	}
}

func (service * DummyProductService) List(cursor uint64, limit uint64) ([]recomendation.Product, error){
	productsList := make([]recomendation.Product, limit)
	mapSize := len(service.products)
	if cursor > uint64(mapSize){
		return nil, fmt.Errorf("DummyProductService.List: cursor %d out of range", cursor)
	}
	var pos uint64
	for _, product := range service.products{
		if pos == cursor + limit{
			break
		}
		if pos >= cursor && pos < cursor + limit{
			productsList = append(productsList, product)
		}
		pos++
	}
	return  productsList, nil
}

func (service * DummyProductService) Create(product recomendation.Product) (uint64, error){
	if product, err := service.products[product.Id]; err {
		service.products[product.Id] = product
		return product.Id, nil
	} else{
		return 0, fmt.Errorf("DummyProductService.Create: product with id %d already exist", product.Id)
	}
}

func (service * DummyProductService) Update(id uint64, product recomendation.Product) error {
	if _, err := service.products[id]; err {
		return fmt.Errorf("DummyProductService.Update: product with id %d not found", id)
	} else {
		if id != product.Id {
			if _, errRemove := service.Remove(id); errRemove != nil {
				return errRemove
			}
			if _, errCreate := service.Create(product); errCreate != nil{
				return errCreate
			}
		} else {
			service.products[id] = product
			return nil
		}
	}
	return nil
}

func (service * DummyProductService) Remove(id uint64) (bool, error){
	if _, err := service.products[id]; err {
		return false, fmt.Errorf("DummyProductService.Remove: product with id %d not found", id)
	} else{
		service.Remove(id)
		return true, nil
	}
}