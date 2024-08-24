package product_repository

import (
	"context"
	"fmt"
	"go-ecom/model"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepositoryImpl(db *sqlx.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (ms *ProductRepositoryImpl) CreateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {

	res, err := ms.db.NamedExecContext(ctx,
		"INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)", p)
	if err != nil {
		return nil, fmt.Errorf("error inserting product: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	p.ID = id
	return p, nil

}

func (ms *ProductRepositoryImpl) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	var p model.Product
	err := ms.db.GetContext(ctx, &p, "SELECT * FROM products WHERE id=?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	return &p, nil
}

func (ms *ProductRepositoryImpl) ListProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := ms.db.SelectContext(ctx, &products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}

	return products, nil
}

func (ms *ProductRepositoryImpl) UpdateProduct(ctx context.Context, p *model.Product) (*model.Product, error) {
	_, err := ms.db.NamedExecContext(ctx, "UPDATE products SET name=:name, image=:image, category=:category, description=:description, rating=:rating, num_reviews=:num_reviews, price=:price, count_in_stock=:count_in_stock, updated_at=:updated_at WHERE id=:id", p)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return p, nil
}

func (ms *ProductRepositoryImpl) DeleteProduct(ctx context.Context, id int64) error {
	_, err := ms.db.ExecContext(ctx, "DELETE FROM products WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
