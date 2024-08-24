package order_repository

import (
	"context"
	"fmt"
	"go-ecom/model"

	"github.com/jmoiron/sqlx"
)

type OrderRepositoryImpl struct {
	db *sqlx.DB
}

func NewOrderRepositoryImpl(db *sqlx.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}

func (ms *OrderRepositoryImpl) CreateOrder(ctx context.Context, o *model.Order) (*model.Order, error) {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {

		order, err := createOrder(ctx, tx, o)
		if err != nil {
			return fmt.Errorf("error creating order: %w", err)
		}

		for _, oi := range o.Items {
			oi.OrderID = order.ID
			err = createOrderItem(ctx, tx, oi)
			if err != nil {
				return fmt.Errorf("error creating order item: %w", err)
			}
		}

		return nil

	})
	if err != nil {
		return nil, fmt.Errorf("error creating order: %w", err)
	}

	return o, nil
}

func (ms *OrderRepositoryImpl) ListOrders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	err := ms.db.SelectContext(ctx, &orders, "SELECT * FROM orders")
	if err != nil {
		return nil, fmt.Errorf("error listing orders: %w", err)
	}

	for i := range orders {
		var items []model.OrderItem

		err = ms.db.SelectContext(ctx, &items, "SELECT * FROM order_items WHERE order_id=?", orders[i].ID)
		if err != nil {
			return nil, fmt.Errorf("error getting order items: %w", err)
		}
		orders[i].Items = items
	}
	return orders, nil
}

func (ms *OrderRepositoryImpl) DeleteOrder(ctx context.Context, id int64) error {
	err := ms.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, "DELETE FROM order_items WHERE order_id=?", id)
		if err != nil {
			return fmt.Errorf("error deleting order items: %w", err)
		}

		_, err = tx.ExecContext(ctx, "DELETE FROM orders WHERE id=?", id)
		if err != nil {
			return fmt.Errorf("error deleting order: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error deleting order: %w", err)
	}

	return nil
}

func createOrder(ctx context.Context, tx *sqlx.Tx, o *model.Order) (*model.Order, error) {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO orders (payment_method, tax_price, shipping_price, total_price) VALUES (:payment_method, :tax_price, :shipping_price, :total_price)", o)
	if err != nil {
		return nil, fmt.Errorf("error inserting order: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}
	o.ID = id

	return o, nil
}

func createOrderItem(ctx context.Context, tx *sqlx.Tx, oi model.OrderItem) error {
	res, err := tx.NamedExecContext(ctx, "INSERT INTO order_items (name, quantity, image, price, product_id, order_id) VALUES (:name, :quantity, :image, :price, :product_id, :order_id)", oi)
	if err != nil {
		return fmt.Errorf("error inserting order item: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}
	oi.ID = id

	return nil
}

func (ms *OrderRepositoryImpl) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := ms.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", rbErr)
		}
		return fmt.Errorf("error in transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}
	return nil
}
