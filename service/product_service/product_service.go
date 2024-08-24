package product_service

import (
	"context"
	"go-ecom/model"
)

type ProductService interface {
	CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error)
	GetProduct(ctx context.Context, id int64) (*model.Product, error)
	ListProducts(ctx context.Context) ([]model.Product, error)
	UpdateProduct(ctx context.Context, p *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}
