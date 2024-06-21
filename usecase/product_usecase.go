package usecase

import (
	"products/model"
	"products/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductByID(id int) (model.Product, error) {
	return pu.repository.GetProductByID(id)
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (int, error) {
	return pu.repository.CreateProduct(product)
}
