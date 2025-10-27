package repository

import (
	"product/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetProductByID(ID uint) (*model.Product, error)
	UpdateProductByID(ID uint, updateProduct map[string]interface{}) (*model.Product, error)
	DeleteProductByID(ID uint) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	err := r.db.Create(product).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetProductByID(ID uint) (*model.Product, error) {
	var product model.Product

	err := r.db.First(&product, ID).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) UpdateProductByID(ID uint, updateProduct map[string]interface{}) (*model.Product, error) {
	var product model.Product

	err := r.db.Model(&product).Where("product_id = ?", ID).Updates(updateProduct).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Where("product_id = ?", ID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) DeleteProductByID(ID uint) (*model.Product, error) {
	var product model.Product

	err := r.db.First(&product, ID).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Delete(&product).Error

	if err != nil {
		return nil, err
	}

	return &product, err
}