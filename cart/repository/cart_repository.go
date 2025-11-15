package repository

import (
	"cart/exceptions"
	"cart/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(ctx context.Context, cart *model.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(ctx context.Context, cart *model.Cart) error {
	if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exceptions.ErrCartFound
		} else {
			return err
		}
	}

	return nil
}