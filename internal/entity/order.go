package entity

import "time"

type Order struct {
	ID        int64       `db:"id" json:"id,omitempty"`
	Total     int         `db:"total" json:"total"`
	CreatedAt time.Time   `db:"created_at" json:"created_at,omitempty"`
	Items     []OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID           int64     `db:"id" json:"id,omitempty"`
	OrderID      int64     `db:"order_id" json:"order_id,omitempty"`
	ProductID    int64     `db:"product_id" json:"product_id"`
	ProductCode  string    `db:"product_code" json:"product_code"`
	ProductName  string    `db:"product_name" json:"product_name"`
	ProductPrice int       `db:"product_price" json:"product_price"`
	Quantity     int       `db:"quantity" json:"quantity"`
	Subtotal     int       `db:"subtotal" json:"subtotal"`
	CreatedAt    time.Time `db:"created_at" json:"created_at,omitempty"`
}

type AnnualIncome struct {
	Year   int    `json:"year"`
	Month  string `json:"month"`
	Income int    `json:"income"`
}

type CreateOrderParam struct {
	Total int                    `db:"total" json:"total,omitempty"`
	Items []CreateOrderItemParam `json:"order_items" validate:"required"`
}

type CreateOrderItemParam struct {
	ProductId int64 `db:"product_id" json:"product_id" validate:"required"`
	Quantity  int   `db:"quantity" json:"quantity" validate:"required"`
	Subtotal  int   `db:"subtotal" json:"subtotal,omitempty"`
	OrderId   int64
}
