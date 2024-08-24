package order_repository

import (
	"context"
	"go-ecom/model"

	"github.com/jmoiron/sqlx"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, o *model.Order) (*model.Order, error)
	ListOrders(ctx context.Context) ([]*model.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
	execTx(ctx context.Context, fn func(*sqlx.Tx) error) error
}
