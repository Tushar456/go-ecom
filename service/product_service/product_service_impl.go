package product_service

import (
	"context"
	"go-ecom/model"
	"go-ecom/repository/product_repository"
)

type ProductServiceImpl struct {
	repo product_repository.ProductRepository
}

func NerProductServiceImpl(r product_repository.ProductRepository) ProductService {

	return &ProductServiceImpl{
		repo: r,
	}

}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	return s.repo.CreateProduct(ctx, p)
}

func (s *ProductServiceImpl) GetProduct(ctx context.Context, id int64) (*model.Product, error) {

	return s.repo.GetProduct(ctx, id)

}

func (s *ProductServiceImpl) ListProducts(ctx context.Context) ([]model.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	return s.repo.UpdateProduct(ctx, p)
}

func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.DeleteProduct(ctx, id)
}
