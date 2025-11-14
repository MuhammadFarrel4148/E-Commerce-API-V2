package model

import "time"

type InputInventory struct {
	ProductID  uint `json:"product_id"`
	StockLevel int  `json:"stock_level"`
}

type UpdateInventory struct {
	ProductID  *uint `json:"product_id"`
	StockLevel *int  `json:"stock_level"`
}

type Inventory struct {
	InventoryID uint `gorm:"primaryKey"`
	StockLevel  int  `gorm:"not null;default:0"`
	UpdatedAt   time.Time
	ProductID   uint    `gorm:"not null;unique"`
	Product     Product
}
