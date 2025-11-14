package validate

import (
	"product/exceptions"
	"product/model"
)

func ValidateInputProduct(inputProduct model.CreateProductInput) error {
	errors := make(map[string]string)

	if inputProduct.Price < 0 {
		errors["price"] = "price can't be negative"
	}

	if len(errors) > 0 {
		return &exceptions.ErrValidation{
			Details: errors,
		}
	}

	return nil
}

func ValidateUpdateProduct(inputProduct model.UpdateProductInput) error {
	errors := make(map[string]string)

	if inputProduct.Price != nil && *inputProduct.Price < 0 {
		errors["price"] = "price can't be negative"
	}

	if len(errors) > 0 {
		return &exceptions.ErrValidation{
			Details: errors,
		}
	}

	return nil
}

func ValidateInputInventory(inputInventory model.InputInventory) error {
	errors := make(map[string]string)

	if inputInventory.StockLevel < 0 {
		errors["stocklevel"] = "stock can't be negative"
	}

	if len(errors) > 0 {
		return &exceptions.ErrValidation{
			Details: errors,
		}
	}

	return nil
}

func ValidateUpdateInventory(inputInventory *model.UpdateInventory) error {
	errors := make(map[string]string)

	if *inputInventory.StockLevel < 0 {
		errors["stocklevel"] = "stock can't be negative"
	}

	if len(errors) > 0 {
		return &exceptions.ErrValidation{
			Details: errors,
		}
	}

	return nil
}

func ValidateUpdateMap(updatesMap map[string]interface{}) error {
	errors := make(map[string]string)

	if len(updatesMap) == 0 {
		errors["updatesMap"] = "no change updated"
		return &exceptions.ErrValidation{
			Details: errors,
		}
	}

	return nil
}
