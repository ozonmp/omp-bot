package product

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/config"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	api "github.com/ozonmp/rcn-product-api/pkg/rcn-product-api"
	"google.golang.org/grpc"
)

type GrpcProductService struct {
	ctx    context.Context
	client api.RcnProductApiServiceClient
}

func NewGrpcProductService(ctx context.Context, conf config.Grpc) (*GrpcProductService, error) {
	categoryServiceConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return &GrpcProductService{
		ctx:    ctx,
		client: api.NewRcnProductApiServiceClient(categoryServiceConn),
	}, nil
}

func (s *GrpcProductService) Create(product recomendation.Product) (uint64, error) {
	req := api.CreateProductV1Request{
		Title:       product.Title,
		Description: product.Description,
		Rating:      float32(product.Rating),
	}
	resp, err := s.client.CreateProductV1(s.ctx, &req)
	if err != nil {
		return 0, nil
	}
	return resp.GetProductId(), nil
}

func (s *GrpcProductService) Describe(id uint64) (*recomendation.Product, error) {
	req := api.DescribeProductV1Request{
		ProductId: id,
	}
	resp, err := s.client.DescribeProductV1(s.ctx, &req)
	if err != nil {
		return nil, err
	}
	product := &recomendation.Product{
		Id:          resp.GetProduct().GetId(),
		Title:       resp.GetProduct().GetTitle(),
		Description: resp.GetProduct().GetDescription(),
		Rating:      float64(resp.GetProduct().GetRating()),
	}
	return product, nil
}

func (s *GrpcProductService) List(cursor uint64, limit uint64) ([]recomendation.Product, error) {
	return nil, nil
}
func (s *GrpcProductService) Update(id uint64, product recomendation.Product) error {
	req := api.UpdateProductV1Request{
		ProductId:   id,
		Title:       product.Title,
		Description: product.Description,
		Rating:      float32(product.Rating),
	}
	_, err := s.client.UpdateProductV1(s.ctx, &req)
	return err
}
func (s *GrpcProductService) Remove(id uint64) (bool, error) {
	req := api.RemoveProductV1Request{
		ProductId: id,
	}
	resp, err := s.client.RemoveProductV1(s.ctx, &req)
	if err != nil {
		return false, err
	}
	return resp.GetFound(), nil
}
func (s *GrpcProductService) Size() int {
	return 0
}
