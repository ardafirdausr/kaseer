package entity

import "time"

type Order struct {
	ID        int64        `json:"id,omitempty"`
	Total     int          `json:"total"`
	CreatedAt time.Time    `json:"created_at,omitempty"`
	Items     []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID           int64     `json:"id,omitempty"`
	OrderID      int64     `json:"order_id,omitempty"`
	ProductID    int64     `json:"product_id"`
	ProductCode  string    `json:"product_code"`
	ProductName  string    `json:"product_name"`
	ProductPrice int       `json:"product_price"`
	Quantity     int       `json:"quantity"`
	Subtotal     int       `json:"subtotal"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

type AnnualIncome struct {
	Year   int    `json:"year"`
	Month  string `json:"month"`
	Income int    `json:"income"`
}

type CreateOrderParam struct {
	Total int                     `json:"total,omitempty"`
	Items []*CreateOrderItemParam `json:"order_items" validate:"required"`
}

type CreateOrderItemParam struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required"`
	Subtotal  int   `json:"subtotal,omitempty"`
	OrderId   int64
}
