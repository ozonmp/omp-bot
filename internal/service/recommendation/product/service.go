package product

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	"log"
	"sort"
)

type Service interface {
	Describe(id uint64) (*recomendation.Product, error)
	List(cursor uint64, limit uint64) ([]recomendation.Product, error)
	Create(recomendation.Product) (uint64, error)
	Update(id uint64, product recomendation.Product) error
	Remove(id uint64) (bool, error)
	Size() int
}

type DummyProductService struct {
	products map[uint64]recomendation.Product
}

func NewDummyProductService() *DummyProductService {
	products := make(map[uint64]recomendation.Product)
	return &DummyProductService{products: products}
}

func (service *DummyProductService) Create(product recomendation.Product) (uint64, error) {
	if _, ok := service.products[product.Id]; !ok {
		service.products[product.Id] = product
		return product.Id, nil
	} else {
		msg := fmt.Sprintf("product with id %d already exist", product.Id)
		log.Printf("DummyProductService.Create: %s", msg)
		return 0, errors.New(msg)
	}
}

func (service *DummyProductService) Update(id uint64, product recomendation.Product) error {
	if _, ok := service.products[id]; !ok {
		msg := fmt.Sprintf("Product with id %d not found", id)
		log.Printf("DummyProductService.Update: %s\n", msg)
		return fmt.Errorf(msg)
	} else {
		if id != product.Id {
			if _, errCreate := service.Create(product); errCreate != nil {
				return errCreate
			}
			if _, errRemove := service.Remove(id); errRemove != nil {
				return errRemove
			}
		} else {
			service.products[id] = product
			return nil
		}
	}
	return nil
}

func (service *DummyProductService) Describe(id uint64) (*recomendation.Product, error) {
	if product, ok := service.products[id]; !ok {
		err := fmt.Errorf("Product with id %d not found", id)
		log.Printf("DummyProductService.Describe: %s", err.Error())
		return nil, err
	} else {
		log.Printf("Get product with id %d", id)
		return &product, nil
	}
}

func (service *DummyProductService) List(cursor uint64, limit uint64) ([]recomendation.Product, error) {
	mapSize := uint64(len(service.products))
	if cursor > mapSize {
		return nil, fmt.Errorf("DummyProductService.List: cursor %d out of range", cursor)
	}
	productsList := make([]recomendation.Product, 0, mapSize)
	for _, product := range service.products {
		productsList = append(productsList, product)
	}
	sort.SliceStable(productsList, func(i, j int) bool {
		return productsList[i].Id < productsList[j].Id
	})
	if cursor+limit < mapSize {
		return productsList[cursor : cursor+limit : cursor+limit], nil
	} else {
		return productsList[cursor:], nil
	}
}

func (service *DummyProductService) Size() int {
	return len(service.products)
}

func (service *DummyProductService) Remove(id uint64) (bool, error) {
	if _, ok := service.products[id]; !ok {
		log.Printf("DummyProductService.Remove: Product with id %d not found", id)
		return false, fmt.Errorf("Product with id %d not found", id)
	} else {
		log.Printf("Product with id %d removed", id)
		delete(service.products, id)
		return true, nil
	}
}
