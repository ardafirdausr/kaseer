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
	ID   int64
	Code string
	Name string
	Sale int
}

type CreateProductParam struct {
	Code  string `db:"code" validate:"required"`
	Name  string `db:"name" validate:"required"`
	Price int    `db:"price" validate:"required,numeric,gt=0"`
	Stock int    `db:"stock" validate:"required,numeric,gte=0"`
}

type UpdateProductParam struct {
	Code  string `db:"code"`
	Name  string `db:"name"`
	Price int    `db:"price" validate:"numeric,gt=0"`
	Stock int    `db:"stock" validate:"numeric,gte=0"`
}
