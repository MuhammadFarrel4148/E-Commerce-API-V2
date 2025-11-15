package model

import "time"

type CartItem struct {
	CartItemID uint `gorm:"primaryKey"`
	CartID     uint `gorm:"not null"`
	Cart       Cart
	ProductID  uint `gorm:"not null"`
	Quantity   uint `gorm:"not null"`
	CreatedAt  time.Time
}
