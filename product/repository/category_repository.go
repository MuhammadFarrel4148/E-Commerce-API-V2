package repository

import (
	"product/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *model.Category) error
	GetCategoryByID(ID uint) (*model.Category, error)
	UpdateCategoryByID(ID uint, updatesCategory map[string]interface{}) (*model.Category, error)
	DeleteCategoryByID(ID uint) (*model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *model.Category) error {
	err := r.db.Create(category).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) GetCategoryByID(ID uint) (*model.Category, error) {
	var category model.Category

	err := r.db.First(&category, ID).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) UpdateCategoryByID(ID uint, updatesCategory map[string]interface{}) (*model.Category, error) {
	var category model.Category

	err := r.db.Model(&category).Where("category_id = ?", ID).Updates(updatesCategory).Error

	if err != nil {
		return nil, err
	}

	err = r.db.First(&category, ID).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) DeleteCategoryByID(ID uint) (*model.Category, error) {
	var category model.Category

	err := r.db.First(&category, ID).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Delete(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}