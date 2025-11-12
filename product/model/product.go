package model

import "time"

// Input Product
type CreateProductInput struct {
	Name       string  `json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
}

type UpdateProductInput struct {
	Name       *string  `json:"name"`
	Price      *float64 `json:"price"`
	CategoryID *uint    `json:"category_id"`
}

// Response product
type CategoryResponse struct {
	CategoryID          uint
	CategoryName        string
	CategoryDescription string
}

type ProductResponse struct {
	ProductID uint
	Name      string
	Price     float64
	Category  CategoryResponse
}

// Table product database
type Product struct {
	ProductID  uint    `gorm:"primaryKey"`
	Name       string  `gorm:"type:varchar(255);not null"`
	Price      float64 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CategoryID uint     `gorm:"not null"`
	Category   Category
}

func FormatProduct(product Product) ProductResponse {
	categoryResponse := CategoryResponse{
		CategoryID:          product.Category.CategoryID,
		CategoryName:        product.Category.Name,
		CategoryDescription: product.Category.Description,
	}

	productResponse := ProductResponse{
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		Category:  categoryResponse,
	}

	return productResponse
}
