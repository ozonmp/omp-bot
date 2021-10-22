package product

import (
	model "github.com/ozonmp/omp-bot/internal/model/lists/product"
)

import "errors"

type ProductService interface {
	Describe(productID uint64) (*model.Product, error)
	List(cursor uint64, limit uint64) ([]model.Product, error)
	Create(product model.Product) (uint64, error)
	Update(productID uint64, product model.Product) error
	Remove(productID uint64) (bool, error)
	Count() uint64
}
type DummyProductService struct{}

func (p *DummyProductService) Describe(productID uint64) (*model.Product, error) {
	if productID > p.Count()-1 {
		return nil, errors.New("index out of bounds")
	}
	return &model.AllProducts[productID], nil
}

func (p DummyProductService) List(cursor uint64, limit uint64) ([]model.Product, error) {
	var maxInd = p.Count()
	if cursor > maxInd {
		return nil, errors.New("index out of bounds")
	}
	if maxInd < cursor+limit {
		return model.AllProducts[cursor:maxInd], nil
	}
	return model.AllProducts[cursor : cursor+limit], nil
}

func (p DummyProductService) Create(product model.Product) (uint64, error) {
	return 0, errors.New("unimplemented")
}

func (p DummyProductService) Update(productID uint64, product model.Product) error {
	return errors.New("unimplemented")
}

func (p DummyProductService) Remove(productID uint64) (bool, error) {
	if productID > p.Count()-1 {
		return false, errors.New("index out of bounds")
	}
	model.AllProducts = append(model.AllProducts[:productID], model.AllProducts[productID+1:]...)
	return true, nil
}

func (p DummyProductService) Count() uint64 {
	return uint64(len(model.AllProducts))
}

func NewDummyProductService() *DummyProductService {
	return &DummyProductService{}
}

//var _ ProductService = (*DummyProductService)(nil)
