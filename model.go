package main

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
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

type ProductOrderDetail struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Subtotal int    `json:"subtotal"`
}

func (o *Order) FormatDate() string {
	return o.CreatedAt.Format("02 January 2006 - 15:04 WIB")
}

func (o *Order) Save() error {
	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	ors, err := tx.ExecContext(ctx, "INSERT INTO orders(total) VALUES(?)", o.Total)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
		return err
	}

	orderId, err := ors.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
		return err
	}

	params := []string{}
	vals := []interface{}{}
	for _, item := range o.Items {
		params = append(params, "(?, ?, ?, ?)")
		vals = append(vals, orderId, item.ProductId, item.Quantity, item.Subtotal)
		item.OrderId = orderId
	}

	query := fmt.Sprintf("INSERT INTO order_items(order_id, product_id, quantity, subtotal) VALUES %s", strings.Join(params, ", "))
	_, err = tx.ExecContext(ctx, query, vals...)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	return nil
}

func GetAllOrders() ([]Order, error) {
	rows, err := DB.Query("SELECT * from orders ORDER BY created_at DESC")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		var err = rows.Scan(
			&order.ID,
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

func GetOrderDetail(orderId int) ([]ProductOrderDetail, error) {
	rows, err := DB.Query("SELECT p.id, p.code, p.name, p.price, oi.quantity, oi.subtotal FROM order_items AS oi LEFT JOIN products AS p ON oi.product_id = p.id WHERE oi.order_id = ?", orderId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var products []ProductOrderDetail

	for rows.Next() {
		var product ProductOrderDetail
		var err = rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.Subtotal,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
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
