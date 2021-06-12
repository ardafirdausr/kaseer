package entity

import "time"

type Order struct {
	ID        int64       `db:"id" json:"id,omitempty"`
	Total     int         `db:"total" json:"total"`
	CreatedAt time.Time   `db:"created_at" json:"created_at,omitempty"`
	Items     []OrderItem `json:"order_items" validate:"required"`
}

type OrderItem struct {
	ID        int64     `db:"id" json:"id,omitempty"`
	OrderId   int64     `db:"order_id" json:"order_id,omitempty"`
	ProductId int64     `db:"product_id" json:"product_id" validate:"required"`
	Quantity  int       `db:"quantity" json:"quantity" validate:"required"`
	Subtotal  int       `db:"subtotal" json:"subtotal,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	Product   Product
}

type AnnualIncome struct {
	Year   int
	Month  string
	Income int
}

type CreateOrderParam struct {
	Total int                    `db:"total" json:"total"`
	Items []CreateOrderItemParam `json:"order_items" validate:"required"`
}

type CreateOrderItemParam struct {
	ProductId int64 `db:"product_id" json:"product_id" validate:"required"`
	Quantity  int   `db:"quantity" json:"quantity" validate:"required"`
	Subtotal  int   `db:"subtotal" json:"subtotal,omitempty"`
	OrderId   int64
}
