package model

import "time"

type Order struct {
	ID            int64      `db:"id"`
	PaymentMethod string     `db:"payment_method"`
	TaxPrice      float32    `db:"tax_price"`
	ShippingPrice float32    `db:"shipping_price"`
	TotalPrice    float32    `db:"total_price"`
	UserID        int64      `db:"user_id"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	Items         []OrderItem
}

type OrderItem struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Quantity  int64   `db:"quantity"`
	Image     string  `db:"image"`
	Price     float32 `db:"price"`
	ProductID int64   `db:"product_id"`
	OrderID   int64   `db:"order_id"`
}
