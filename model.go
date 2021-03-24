package main

import (
	"log"
	"time"
)

type Product struct {
	ID        int       `db:"id"`
	Code      string    `db:"code"`
	Name      string    `db:"name"`
	Price     int       `db:"price"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (p *Product) GetAllProducts() ([]Product, error) {
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
			&product.Quantity,
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
