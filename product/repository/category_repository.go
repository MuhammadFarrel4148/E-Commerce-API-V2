package repository

import (
	"context"
	"errors"
	"product/exceptions"
	"product/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategoryByID(ctx context.Context, ID uint) (*model.Category, error)
	UpdateCategoryByID(ctx context.Context, ID uint, updatesCategory map[string]interface{}) (*model.Category, error)
	DeleteCategoryByID(ctx context.Context, ID uint) (*model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

// Helper find product
func (r *categoryRepository) findProduct(ctx context.Context, ID uint) (*model.Category, error) {
	var category model.Category

	if err := r.db.WithContext(ctx).First(&category, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrNotFound
		} else {
			return nil, err
		}
	}

	return &category, nil
};

func (r *categoryRepository) CreateCategory(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exceptions.ErrNameFound
		} else {
			return err
		}
	}

	return nil
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, ID uint) (*model.Category, error) {
	category, err := r.findProduct(ctx, ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategoryByID(ctx context.Context, ID uint, updatesCategory map[string]interface{}) (*model.Category, error) {
	var category model.Category

	if err := r.db.WithContext(ctx).Model(&category).Where("category_id = ?", ID).Updates(updatesCategory).Error; err != nil {
		return nil, err
	}

	categories, err := r.findProduct(ctx, ID)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) DeleteCategoryByID(ctx context.Context, ID uint) (*model.Category, error) {
	categories, err := r.findProduct(ctx, ID)
	if err != nil {
		return nil, err
	}

	if err := r.db.Delete(categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}