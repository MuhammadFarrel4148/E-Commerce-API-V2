package service

import (
	"errors"
	"product/model"
	"product/repository"
)

type CreateProductInput struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}

type UpdateProductInput struct {
	Name       *string  `json:"name"`
	Price      *float64 `json:"price"`
	CategoryID *uint    `json:"category_id"`
}

type ProductService interface {
	CreateProductService(inputProduct CreateProductInput) (*model.Product, error)
	GetProductServiceByID(ID uint) (*model.Product, error)
	UpdateProductServiceByID(ID uint, inputProduct *UpdateProductInput) (*model.Product, error)
	DeleteProductServiceByID(ID uint) (*model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProductService(inputProduct CreateProductInput) (*model.Product, error) {
	if inputProduct.Price < 0 {
		return nil, errors.New("price can't be negative")
	}
	if inputProduct.Name == "" {
		return nil, errors.New("name can't be empty")
	}

	product := &model.Product{
		Name:       inputProduct.Name,
		Price:      inputProduct.Price,
		CategoryID: inputProduct.CategoryID,
	}

	err := s.repo.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProductServiceByID(ID uint) (*model.Product, error) {
	product, err := s.repo.GetProductByID(ID)

	if err != nil {
		return nil, err
	}

	return product, err
}

func (s *productService) UpdateProductServiceByID(ID uint, inputProduct *UpdateProductInput) (*model.Product, error) {
	if inputProduct.Price != nil && *inputProduct.Price < 0 {
		return nil, errors.New("price can't be negative")
	}

	updatesMap := make(map[string]interface{})

	if inputProduct.Name != nil {
		updatesMap["name"] = *inputProduct.Name
	}
	if inputProduct.Price != nil {
		updatesMap["price"] = *inputProduct.Price
	}
	if inputProduct.CategoryID != nil {
		updatesMap["category_id"] = *inputProduct.CategoryID
	}

	if len(updatesMap) == 0 {
		return s.repo.GetProductByID(ID)
	}

	updatedProduct, err := s.repo.UpdateProductByID(ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (s *productService) DeleteProductServiceByID(ID uint) (*model.Product, error) {
	product, err := s.repo.DeleteProductByID(ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}
