package model

import "time"

type Cart struct {
	CartID    uint		`gorm:"primaryKey"`
	UserId    string	`gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
