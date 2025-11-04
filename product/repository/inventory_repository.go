package repository

import (
	"product/model"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	// VerifyCategory(name string) error
	CreateInventory(inventory *model.Inventory) error
	GetInventoryByID(ID uint) (*model.Inventory, error)
	UpdateInventoryByID(ID uint, updatesInventory map[string]interface{}) (*model.Inventory, error)
	DeleteInventoryByID(ID uint) (*model.Inventory, error)
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db}
}

// func (r *inventoryRepository) VerifyCategory(name string) error {
// 	var category model.Category

// 	err := r.db.Where("name = ?", name).First(&category).Error

// 	if err == nil {
// 		return errors.New("category name already exists")
// 	}

// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil
// 	}

// 	return err
// }

func (r *inventoryRepository) CreateInventory(inventory *model.Inventory) error {
	err := r.db.Create(inventory).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryRepository) GetInventoryByID(ID uint) (*model.Inventory, error) {
	var inventory model.Inventory

	err := r.db.First(&inventory, ID).Error

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) UpdateInventoryByID(ID uint, updatesInventory map[string]interface{}) (*model.Inventory, error) {
	var inventory model.Inventory

	err := r.db.Model(&inventory).Where("category_id = ?", ID).Updates(updatesInventory).Error

	if err != nil {
		return nil, err
	}

	err = r.db.First(&inventory, ID).Error

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (r *inventoryRepository) DeleteInventoryByID(ID uint) (*model.Inventory, error) {
	var inventory model.Inventory

	err := r.db.First(&inventory, ID).Error

	if err != nil {
		return nil, err
	}

	err = r.db.Delete(&inventory).Error

	if err != nil {
		return nil, err
	}

	return &inventory, nil
}
