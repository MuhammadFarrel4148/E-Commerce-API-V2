package validate

import (
	"errors"
	"product/model"
)

func ValidateInputProduct(inputProduct model.CreateProductInput) error {
	if inputProduct.Price < 0 {
		return errors.New("price can't be negative")
	}

	if inputProduct.Name == "" {
		return errors.New("name can't be empty")
	}

	return nil
}

func ValidateUpdateProduct(inputProduct model.UpdateProductInput) error {
	if inputProduct.Price != nil && *inputProduct.Price < 0 {
		return errors.New("price can't be negative")
	}

	return nil
}
