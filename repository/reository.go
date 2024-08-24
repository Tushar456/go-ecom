package repository

import (
	"go-ecom/repository/order_repository"
	"go-ecom/repository/product_repository"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	OrderRepo   order_repository.OrderRepository
	ProductRepo product_repository.ProductRepository
}

func NewRepository(db *sqlx.DB) Repository {

	return Repository{
		OrderRepo:   order_repository.NewOrderRepositoryImpl(db),
		ProductRepo: product_repository.NewProductRepositoryImpl(db),
	}

}
