package service

import (
	"errors"
	"product/model"
	"product/repository"
)

type InputInventory struct {
	ProductID  uint `json:"product_id"`
	StockLevel int  `json:"stock_level"`
}

type UpdateInventory struct {
	ProductID  *uint `json:"product_id"`
	StockLevel *int  `json:"stock_level"`
}

type InventoryService interface {
	CreateInventory(inputInventory InputInventory) (*model.Inventory, error)
	GetInventoryByID(ID uint) (*model.Inventory, error)
	UpdateInventoryByID(ID uint, updateInventory *UpdateInventory) (*model.Inventory, error)
	DeleteInventoryByID(ID uint) (*model.Inventory, error)
}

type inventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repo}
}

func (s *inventoryService) CreateInventory(inputInventory InputInventory) (*model.Inventory, error) {
	if inputInventory.StockLevel < 0 {
		return nil, errors.New("stock level can't be negative")
	}

	inventory := &model.Inventory{
		ProductID:  inputInventory.ProductID,
		StockLevel: inputInventory.StockLevel,
	}

	err := s.repo.CreateInventory(inventory)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) GetInventoryByID(ID uint) (*model.Inventory, error) {
	inventory, err := s.repo.GetInventoryByID(ID)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) UpdateInventoryByID(ID uint, updateInventory *UpdateInventory) (*model.Inventory, error) {
	if *updateInventory.StockLevel < 0 {
		return nil, errors.New("stock level can't be negative")
	}

	updatesMap := make(map[string]interface{})

	if updateInventory.ProductID != nil {
		updatesMap["product_id"] = *updateInventory.ProductID
	}

	if updateInventory.StockLevel != nil {
		updatesMap["stock_level"] = *updateInventory.StockLevel
	}

	if len(updatesMap) == 0 {
		return nil, errors.New("no change updated")
	}

	inventory, err := s.repo.UpdateInventoryByID(ID, updatesMap)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *inventoryService) DeleteInventoryByID(ID uint) (*model.Inventory, error) {
	inventory, err := s.repo.DeleteInventoryByID(ID)

	if err != nil {
		return nil, err
	}

	return inventory, nil
}
