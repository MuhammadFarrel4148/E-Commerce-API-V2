package model

import "time"

type Cart struct {
	CartID    uint		`gorm:"primaryKey"`
	UserID    string	`gorm:"type:varchar(50);not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
