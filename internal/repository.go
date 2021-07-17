package internal

import "github.com/ardafirdausr/kaseer/internal/entity"

type UserRepository interface {
	// Get User by user ID.
	// if the user found, return user and nil
	// if the user not found, return nil and ErrNotFound
	// if other an error happens, return nil and the error
	GetUserByID(ID int64) (*entity.User, error)

	// Get User by email.
	// if the user found, return user and nil
	// if the user not found, return nil and ErrNotFound
	// if other an error happens, return nil and the error
	GetUserByEmail(email string) (*entity.User, error)

	// Update User data by user ID.
	// if update success, return true and nil
	// if update an error happens, return false and the error
	UpdateByID(ID int64, param entity.UpdateUserParam) (bool, error)

	// Update User password by user ID.
	// if update success, return true and nil
	// if update an error happens, return false and the error
	UpdatePasswordByID(ID int64, password string) (bool, error)
}

type ProductRepository interface {

	// Get All Products
	// if success, returns slice of products
	// if an error happened, returns nil and the error
	GetAllProducts() ([]*entity.Product, error)

	// Get Best Seller Products
	// if success, return slice of products
	// if an error happened, return 0 and the error
	GetBestSellerProducts() ([]*entity.ProductSale, error)

	// Get Products by ids.
	// if success, return the slice of products and nil.
	// if not found, return nil and ErrNotFound.
	// if an error happened, return nil and the error.
	GetProductsByIDs(IDs ...int64) ([]*entity.Product, error)

	// Get Product by product code.
	// if success, return the product and nil.
	// if not found, return nil and ErrNotFound.
	// if an error happened, return nil and the error.
	GetProductByCode(code string) (*entity.Product, error)

	// Get Product by product ID
	// if success, return the product and nil
	// if not found, return nil and ErrNotFound
	// if an error happened, return nil and the error
	GetProductByID(ID int64) (*entity.Product, error)

	// Create Product
	// if success, return the product and nil
	// if an error happened, return nil and the error
	Create(param entity.CreateProductParam) (*entity.Product, error)

	// Update Product
	// if update success, return true and nil
	// if an error happened, return false and the error
	UpdateByID(ID int64, param entity.UpdateProductParam) (bool, error)

	// Create Product
	// if update success, return true and nil
	// if an error happened, return false and the error
	DeleteByID(ID int64) (bool, error)
}

type OrderRepository interface {
	GetAllOrders() ([]*entity.Order, error)
	GetAnnualIncome() ([]*entity.AnnualIncome, error)
	GetDailyOrderCount() (int, error)
	GetTotalOrderCount() (int, error)
	GetLastDayIncome() (int, error)
	GetLastMonthIncome() (int, error)
	GetOrderItemsByID(ID int64) ([]*entity.OrderItem, error)
	Create(param entity.CreateOrderParam) (*entity.Order, error)
}
