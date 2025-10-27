package model

import "time"

type Category struct {
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string
}

type Product struct {
	ProductID  uint    `gorm:"primaryKey"`
	Name       string  `gorm:"type:varchar(255);not null"`
	Price      float64 `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}

type Inventory struct {
	InventoryID uint `gorm:"primaryKey"`
	StockLevel  int  `gorm:"not null;default:0"`
	UpdatedAt   time.Time
	ProductID   uint    `gorm:"not null;unique"`
	Product     Product `gorm:"foreignKey:ProductID"`
}
