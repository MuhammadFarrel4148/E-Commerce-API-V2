package repository

import (
	"context"
	"errors"
	"product/exceptions"
	"product/model"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	CreateInventory(ctx context.Context, inventory *model.Inventory) error
	GetInventoryByID(ctx context.Context, ID uint) (*model.Inventory, error)
	UpdateInventoryByID(ctx context.Context, ID uint, updatesInventory map[string]interface{}) (*model.Inventory, error)
	DeleteInventoryByID(ctx context.Context, ID uint) (*model.Inventory, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) findInventory(ctx context.Context, inventory *model.Inventory, ID uint) error {
	if err := r.db.WithContext(ctx).First(&inventory, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.ErrNotFound
		}
	}

	return nil
}

func (r *inventoryRepository) CreateInventory(ctx context.Context, inventory *model.Inventory) error {
	if err := r.db.WithContext(ctx).Create(inventory).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return exceptions.ErrProductIDFound
		}
	}

	return nil
}

func (r *inventoryRepository) GetInventoryByID(ctx context.Context, ID uint) (*model.Inventory, error) {
	var inventory model.Inventory

	err := r.findInventory(ctx, &inventory, ID)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) UpdateInventoryByID(ctx context.Context, ID uint, updatesInventory map[string]interface{}) (*model.Inventory, error) {
	var inventory model.Inventory

	if err := r.db.Model(&inventory).Where("category_id = ?", ID).Updates(updatesInventory).Error; err != nil {
		return nil, err
	}

	if err := r.findInventory(ctx, &inventory, ID); err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) DeleteInventoryByID(ctx context.Context, ID uint) (*model.Inventory, error) {
	var inventory model.Inventory
	if err := r.findInventory(ctx, &inventory, ID); err != nil {
		return nil, err
	}

	if err := r.db.Delete(&inventory).Error; err != nil {
		return nil, err
	}

	return &inventory, nil
}
