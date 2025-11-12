package model

import "time"

type Inventory struct {
	InventoryID uint `gorm:"primaryKey"`
	StockLevel  int  `gorm:"not null;default:0"`
	UpdatedAt   time.Time
	ProductID   uint    `gorm:"not null;unique"`
	Product     Product
}
