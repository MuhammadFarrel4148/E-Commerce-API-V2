package service

import (
	"product/model"
	"product/repository"
)

type InputCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategory struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CategoryService interface {
	CreateCategoryService(inputCategory InputCategory) (*model.Category, error)
	GetCategoryByID(ID uint) (*model.Category, error)
	UpdateCategoryByID(ID uint, updateCategory *UpdateCategory) (*model.Category, error)
	DeleteCategoryByID(ID uint) (*model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) CreateCategoryService(InputCategory InputCategory) (*model.Category, error) {
	category := &model.Category{
		Name:        InputCategory.Name,
		Description: InputCategory.Description,
	}

	err := s.repo.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategoryByID(ID uint) (*model.Category, error) {
	category, err := s.repo.GetCategoryByID(ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategoryByID(ID uint, updateCategory *UpdateCategory) (*model.Category, error) {
	updatesMap := make(map[string]interface{})

	if updateCategory.Name != nil {
		updatesMap["name"] = *updateCategory.Name
	}

	if updateCategory.Description != nil {
		updatesMap["description"] = *updateCategory.Description
	}

	category, err := s.repo.UpdateCategoryByID(ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategoryByID(ID uint) (*model.Category, error) {
	category, err := s.repo.DeleteCategoryByID(ID)

	if err != nil {
		return nil, err
	}

	return category, nil
}
