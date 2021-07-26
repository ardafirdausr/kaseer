package usecase

import (
	"context"
	"log"

	"github.com/ardafirdausr/kaseer/internal"
	"github.com/ardafirdausr/kaseer/internal/entity"
)

type ProductUsecase struct {
	productRepository internal.ProductRepository
}

func NewProductUsecase(productRepository internal.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepository: productRepository}
}

func (pu ProductUsecase) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	products, err := pu.productRepository.GetAllProducts(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return products, err
}

func (pu ProductUsecase) GetProductByID(ctx context.Context, ID int64) (*entity.Product, error) {
	product, err := pu.productRepository.GetProductByID(ctx, ID)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) GetProductByCode(ctx context.Context, code string) (*entity.Product, error) {
	product, err := pu.productRepository.GetProductByCode(ctx, code)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) GetBestSellerProducts(ctx context.Context) ([]*entity.ProductSale, error) {
	productSales, err := pu.productRepository.GetBestSellerProducts(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	return productSales, err
}

func (pu ProductUsecase) CreateProduct(ctx context.Context, param entity.CreateProductParam) (*entity.Product, error) {
	exProduct, _ := pu.productRepository.GetProductByCode(ctx, param.Code)
	if exProduct != nil {
		return nil, entity.ErrItemAlreadyExists{
			Message: "Product already exists",
			Err:     nil,
		}
	}

	product, err := pu.productRepository.Create(ctx, param)
	if err != nil {
		log.Println(err.Error())
	}

	return product, err
}

func (pu ProductUsecase) UpdateProduct(ctx context.Context, ID int64, param entity.UpdateProductParam) (bool, error) {
	isUpdated, err := pu.productRepository.UpdateByID(ctx, ID, param)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}

func (pu ProductUsecase) DeleteProduct(ctx context.Context, ID int64) (bool, error) {
	isUpdated, err := pu.productRepository.DeleteByID(ctx, ID)
	if err != nil {
		log.Println(err.Error())
	}

	return isUpdated, err
}
