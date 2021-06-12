package usecase

import (
	"log"

	"github.com/ardafirdausr/go-pos/internal"
	"github.com/ardafirdausr/go-pos/internal/entity"
)

type ProductUsecase struct {
	productRepository internal.ProductRepository
}

func NewProductUsecase(productRepository internal.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepository: productRepository}
}

func (pu ProductUsecase) GetAllProducts() ([]*entity.Product, error) {
	products, err := pu.productRepository.GetAllProducts()
	if err != nil {
		log.Println(err.Error())
	}

	return products, err
}

func (pu ProductUsecase) GetProductByID(ID int64) (*entity.Product, error) {
	product, err := pu.productRepository.GetProductByID(ID)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) GetProductByCode(code string) (*entity.Product, error) {
	product, err := pu.productRepository.GetProductByCode(code)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) GetBestSellerProducts() ([]*entity.ProductSale, error) {
	productSales, err := pu.productRepository.GetBestSellerProducts()
	if err != nil {
		log.Println(err.Error())
	}

	return productSales, err
}

func (pu ProductUsecase) CreateProduct(param entity.CreateProductParam) (*entity.Product, error) {
	product, err := pu.productRepository.Create(param)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) UpdateProduct(ID int64, param entity.UpdateProductParam) (bool, error) {
	isUpdated, err := pu.productRepository.UpdateByID(ID, param)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}

func (pu ProductUsecase) DeleteProduct(ID int64) (bool, error) {
	isUpdated, err := pu.productRepository.DeleteByID(ID)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}
