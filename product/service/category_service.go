package service

import (
	"context"
	"product/model"
	"product/repository"
)

var ctx context.Context

type CategoryService interface {
	CreateCategoryService(inputCategory model.InputCategory) (*model.Category, error)
	GetCategoryByID(ID uint) (*model.Category, error)
	UpdateCategoryByID(ID uint, updateCategory *model.UpdateCategory) (*model.Category, error)
	DeleteCategoryByID(ID uint) (*model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategoryService(InputCategory model.InputCategory) (*model.Category, error) {
	category := &model.Category{
		Name:        InputCategory.Name,
		Description: InputCategory.Description,
	}

	if err := s.repo.CreateCategory(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategoryByID(ID uint) (*model.Category, error) {
	category, err := s.repo.GetCategoryByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategoryByID(ID uint, updateCategory *model.UpdateCategory) (*model.Category, error) {
	updatesMap := make(map[string]interface{})

	if updateCategory.Name != nil {
		updatesMap["name"] = *updateCategory.Name
	}

	if updateCategory.Description != nil {
		updatesMap["description"] = *updateCategory.Description
	}

	category, err := s.repo.UpdateCategoryByID(ctx, ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategoryByID(ID uint) (*model.Category, error) {
	category, err := s.repo.DeleteCategoryByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}
