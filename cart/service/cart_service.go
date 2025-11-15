package service

import (
	"cart/model"
	"cart/repository"
	"context"
)

var ctx context.Context

type CartService interface {
	CreateCartService(UserID string) (*model.Cart, error)
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo}
}

func (r *cartService) CreateCartService(UserID string) (*model.Cart, error) {
	cart := &model.Cart{
		UserID: UserID,
	}

	if err := r.cartRepo.CreateCart(ctx, cart); err != nil {
		return nil, err
	}

	return cart, nil
}