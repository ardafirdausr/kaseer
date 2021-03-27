package main

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type M map[string]interface{}

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

	_, err = DB.Exec(
		"UPDATE products SET code = ?, name = ?, stock = ?, price = ? WHERE id = ?",
		p.Code, p.Name, p.Stock, p.Price, p.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (p *Product) Delete() error {
	_, err := DB.Exec("DELETE FROM products WHERE id = ?", p.ID)
	if err != nil {
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

type Order struct {
	ID        int64     `db:"id"`
	Code      string    `db:"code"`
	Total     string    `db:"total"`
	CreatedAt time.Time `db:"created_at"`
	Items     []OrderItem
}

type OrderItem struct {
	ID        int64     `db:"id"`
	OrderId   int64     `db:"order_id" validate:"required" `
	ProductId int64     `db:"product_id" validate:"required"`
	Quantity  string    `db:"quantity" validate:"required"`
	Subtotal  int       `db:"subtotal" validate:"required"`
	CreatedAt time.Time `db:"created_at"`
}

func (o *Order) Save() error {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	ors, err := tx.ExecContext(ctx, "INSERT INTO orders(code, total) VALUES(?, ?)", o.Code, o.Total)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}

	orderId, err := ors.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}

	sqlStr := "INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES "
	vals := []interface{}{}

	for _, item := range o.Items {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, orderId, item.ProductId, item.Quantity, item.Subtotal)
		item.OrderId = orderId
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, err := DB.PrepareContext(ctx, sqlStr)
	_, err = stmt.Exec(vals...)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	return nil
}

func GetAllOrders() ([]Order, error) {
	rows, err := DB.Query("SELECT * FROM orders")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order = Order{}
		var err = rows.Scan(
			&order.ID,
			&order.Code,
			&order.Total,
			&order.CreatedAt,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return orders, nil
}

type User struct {
	ID        int64
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	PhotoUrl  *string   `db:"omitempty,photo_url"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) changePassword(password string) {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	u.Password = fmt.Sprintf("%x", hashed)
}

func (u *User) CheckPassword(password string) bool {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashed := hash.Sum(nil)
	return fmt.Sprintf("%x", hashed) == u.Password
}

func (u *User) Update() error {
	_, err := DB.Exec(
		"UPDATE users SET name = ?, email = ?, photo_url = ? WHERE id = ?",
		u.Name, u.Email, u.PhotoUrl, u.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func findUserById(userId int64) (*User, error) {
	row := DB.QueryRow("SELECT * FROM users WHERE id = ?", userId)

	var user User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func findUserByEmail(email string) (*User, error) {
	row := DB.QueryRow("SELECT * FROM users WHERE email = ?", email)

	var user User
	var err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PhotoUrl,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
