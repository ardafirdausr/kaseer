package entity

import "time"

type Product struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductSale struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Sale int    `json:"sale"`
}

type CreateProductParam struct {
	Code  string `json:"code" form:"code" validate:"required"`
	Name  string `json:"name" form:"name" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required,numeric,gt=0"`
	Stock int    `json:"stock" form:"stock" validate:"required,numeric,gte=0"`
}

type UpdateProductParam struct {
	Code  string `form:"code"`
	Name  string `form:"name"`
	Price int    `form:"price" validate:"numeric,gt=0"`
	Stock int    `form:"stock" validate:"numeric,gte=0"`
}
