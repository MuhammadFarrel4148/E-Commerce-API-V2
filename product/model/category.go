package model

type Category struct {
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null;unique"`
	Description string
}
