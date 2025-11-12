package repository

import (
	"context"
	"product/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	PreloadProduct(ctx context.Context, ID uint) (*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	GetProductByID(ctx context.Context, ID uint) (*model.Product, error)
	UpdateProductByID(ctx context.Context, ID uint, updateProduct map[string]interface{}) (*model.Product, error)
	DeleteProductByID(ctx context.Context, ID uint) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// Helper Get Product By Preload Category
func (r *productRepository) findProductByID(ctx context.Context, ID uint, preload bool) (*model.Product, error) {
	var product model.Product
	db := r.db.WithContext(ctx)

	if preload {
		db = db.Preload("Category")
	}

	if err := db.First(&product, ID).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) PreloadProduct(ctx context.Context, ID uint) (*model.Product, error) {
	return r.findProductByID(ctx, ID, true)
}

func (r *productRepository) CreateProduct(ctx context.Context, product *model.Product) error {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) GetProductByID(ctx context.Context, ID uint) (*model.Product, error) {	
	product, err := r.PreloadProduct(ctx, ID)
	if err != nil{
		return nil, err
	}

	return product, nil
}

func (r *productRepository) UpdateProductByID(ctx context.Context, ID uint, updateProduct map[string]interface{}) (*model.Product, error) {
	if err := r.db.Model(&model.Product{}).Where("product_id = ?", ID).Updates(updateProduct).Error; err != nil {
		return nil, err
	}

	product, err := r.PreloadProduct(ctx, ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepository) DeleteProductByID(ctx context.Context, ID uint) (*model.Product, error) {
	product, err := r.PreloadProduct(ctx, ID)
	if err != nil {
		return nil, err
	}
	
	if err := r.db.Delete(&model.Product{}, ID).Error; err != nil {
		return nil, err
	}

	return product, nil
}
