package service

import (
	"context"
	"product/model"
	"product/repository"
	"product/validate"
)

type ProductService interface {
	CreateProductService(ctx context.Context, inputProduct model.CreateProductInput) (*model.ProductResponse, error)
	GetProductServiceByID(ctx context.Context, ID uint) (*model.Product, error)
	UpdateProductServiceByID(ctx context.Context, ID uint, inputProduct model.UpdateProductInput) (*model.Product, error)
	DeleteProductServiceByID(ctx context.Context, ID uint) (*model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProductService(ctx context.Context, inputProduct model.CreateProductInput) (*model.ProductResponse, error) {
	if err := validate.ValidateInputProduct(inputProduct); err != nil {
		return nil, err
	}

	product := &model.Product{
		Name:       inputProduct.Name,
		Price:      inputProduct.Price,
		CategoryID: inputProduct.CategoryID,
	}

	if err := s.repo.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	productOutput, err := s.repo.PreloadProduct(ctx, product.ProductID)

	if err != nil {
		return nil, err
	}

	productResponseValue := model.FormatProduct(*productOutput)
	productResponsePointer := &productResponseValue

	return productResponsePointer, nil
}

func (s *productService) GetProductServiceByID(ctx context.Context, ID uint) (*model.Product, error) {
	product, err := s.repo.PreloadProduct(ctx, ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) UpdateProductServiceByID(ctx context.Context, ID uint, inputProduct model.UpdateProductInput) (*model.Product, error) {
	if err := validate.ValidateUpdateProduct(inputProduct); err != nil {
		return nil, err
	}

	updatesMap := make(map[string]interface{})

	if inputProduct.Name != nil {
		updatesMap["name"] = inputProduct.Name
	}
	if inputProduct.Price != nil {
		updatesMap["price"] = inputProduct.Price
	}
	if inputProduct.CategoryID != nil {
		updatesMap["category_id"] = inputProduct.CategoryID
	}

	if len(updatesMap) == 0 {
		return s.repo.PreloadProduct(ctx, ID)
	}

	updatedProduct, err := s.repo.UpdateProductByID(ctx, ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (s *productService) DeleteProductServiceByID(ctx context.Context, ID uint) (*model.Product, error) {
	product, err := s.repo.DeleteProductByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}
