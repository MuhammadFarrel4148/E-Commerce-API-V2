package model

type InputCategory struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCategory struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type CategoryOutput struct {
	CategoryID	uint	`json:"category_id"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

type Category struct {
	CategoryID  uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null;unique"`
	Description string
}
