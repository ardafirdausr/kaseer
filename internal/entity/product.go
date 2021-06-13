package entity

import "time"

type Product struct {
	ID        int64     `db:"id"`
	Code      string    `db:"code"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	Stock     int       `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductSale struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Sale int    `json:"sale"`
}

type CreateProductParam struct {
	Code  string `db:"code" form:"code" json:"code" validate:"required"`
	Name  string `db:"name" form:"name" json:"name" validate:"required"`
	Price int    `db:"price" form:"price" json:"price" validate:"required,numeric,gt=0"`
	Stock int    `db:"stock" form:"stock" json:"stock" validate:"required,numeric,gte=0"`
}

type UpdateProductParam struct {
	Code  string `db:"code" form:"code"`
	Name  string `db:"name" form:"name"`
	Price int    `db:"price" form:"price" validate:"numeric,gt=0"`
	Stock int    `db:"stock" form:"stock" validate:"numeric,gte=0"`
}
