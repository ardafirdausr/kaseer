package main

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Product struct {
	ID        int64     `db:"id" validate:"omitempty,required,numeric"`
	Code      string    `db:"code" validate:"required"`
	Name      string    `db:"name" validate:"required"`
	Price     int       `db:"price" validate:"required,numeric,gt=0"`
	Stock     int       `db:"stock" validate:"required,numeric,gte=0"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (p *Product) Save() error {
	existProduct, err := FindProductByCode(p.Code)
	if err != nil {
		log.Println(err.Error())
		return err
	} else if existProduct != nil {
		return errors.New("Product code is already exists")
	}

	res, err := DB.Exec(
		"INSERT INTO products(code, name, stock, price) VALUES(?, ?, ?, ?)",
		p.Code, p.Name, p.Stock, p.Price)
	if err != nil {
		return err
	}

	p.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Update() error {
	existProduct, err := FindProductByCode(p.Code)
	if err != nil {
		log.Println(err.Error())
		return err
	} else if existProduct != nil && existProduct.ID != p.ID {
		return errors.New("Product code is already exists")
	}

	res, err := DB.Exec(
		"UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?",
		p.Code, p.Name, p.Stock, p.Price, p.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return err
	} else if rowAffected < 1 {
		return errors.New("Failed to update data")
	}

	return nil
}

func (p *Product) Delete() error {
	res, err := DB.Exec("DELETE FROM products WHERE id = ?", p.ID)
	if err != nil {
		return err
	}

	p.ID, err = res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	p = nil
	return nil
}

func GetAllProducts() ([]Product, error) {
	rows, err := DB.Query("SELECT * FROM products")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product = Product{}
		var err = rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func FindProductByCode(code string) (*Product, error) {
	row := DB.QueryRow("SELECT * FROM products WHERE code = ?", code)

	var product Product
	var err = row.Scan(
		&product.ID,
		&product.Code,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func FindProductById(id int) (*Product, error) {
	row := DB.QueryRow("SELECT * FROM products WHERE id = ?", id)

	var product Product
	var err = row.Scan(
		&product.ID,
		&product.Code,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}
