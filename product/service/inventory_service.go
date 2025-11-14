package service

import (
	"context"
	"product/model"
	"product/repository"
	"product/validate"
)

var ctx context.Context

type InventoryService interface {
	CreateInventory(inputInventory model.InputInventory) (*model.Inventory, error)
	GetInventoryByID(ID uint) (*model.Inventory, error)
	UpdateInventoryByID(ID uint, updateInventory *model.UpdateInventory) (*model.Inventory, error)
	DeleteInventoryByID(ID uint) (*model.Inventory, error)
}

type inventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repo}
}

func (s *inventoryService) CreateInventory(inputInventory model.InputInventory) (*model.Inventory, error) {
	if err := validate.ValidateInputInventory(inputInventory); err != nil {
		return nil, err
	}

	inventory := &model.Inventory{
		ProductID:  inputInventory.ProductID,
		StockLevel: inputInventory.StockLevel,
	}

	err := s.repo.CreateInventory(ctx, inventory)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) GetInventoryByID(ID uint) (*model.Inventory, error) {
	inventory, err := s.repo.GetInventoryByID(ctx, ID)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) UpdateInventoryByID(ID uint, updateInventory *model.UpdateInventory) (*model.Inventory, error) {
	if err := validate.ValidateUpdateInventory(updateInventory); err != nil {
		return nil, err
	}

	updatesMap := make(map[string]interface{})

	if updateInventory.ProductID != nil {
		updatesMap["product_id"] = *updateInventory.ProductID
	}

	if updateInventory.StockLevel != nil {
		updatesMap["stock_level"] = *updateInventory.StockLevel
	}

	if err := validate.ValidateUpdateMap(updatesMap); err != nil {
		return nil, err
	}

	inventory, err := s.repo.UpdateInventoryByID(ctx, ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) DeleteInventoryByID(ID uint) (*model.Inventory, error) {
	inventory, err := s.repo.DeleteInventoryByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return inventory, nil
}
