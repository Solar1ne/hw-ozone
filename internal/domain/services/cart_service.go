package services

import (
	"context"
	"errors"
	"sort"

	"awesomeProject4/internal/domain/models"
)

type CartRepository interface {
	GetCart(userID int64) (*models.Cart, error)
	SaveCart(cart *models.Cart) error
	DeleteCart(userID int64) error
	Checkout(ctx context.Context, userID int64) (OrderID OrderID, error)
}

type ProductService interface {
	GetProduct(ctx context.Context, sku int64) (*models.Product, error)
}

type CartService struct {
	CartRepo      CartRepository
	ProductClient ProductService
}

func NewCartService(repo CartRepository, client ProductService) *CartService {
	return &CartService{
		CartRepo:      repo,
		ProductClient: client,
	}
}

func (cs *CartService) AddItemToCart(ctx context.Context, userID, skuID int64, count uint16) error {
	product, err := cs.ProductClient.GetProduct(ctx, skuID)
	if err != nil {
		return err
	}

	cart, _ := cs.CartRepo.GetCart(userID)
	if cart == nil {
		cart = &models.Cart{
			UserID: &(userID),
			Items:  make(map[int64]*models.CartItem),
		}
	}

	if item, exists := cart.Items[skuID]; exists {
		*item.Count += count
	} else {
		cart.Items[skuID] = &models.CartItem{
			SKU:   &skuID,
			Name:  product.Name,
			Count: &count,
			Price: &product.Price,
		}
	}

	return cs.CartRepo.SaveCart(cart)
}

func (cs *CartService) RemoveItemFromCart(userID, skuID int64) error {
	cart, _ := cs.CartRepo.GetCart(userID)
	if cart == nil {
		return nil
	}

	delete(cart.Items, skuID)

	if len(cart.Items) == 0 {
		return cs.CartRepo.DeleteCart(userID)
	}

	return cs.CartRepo.SaveCart(cart)
}

func (cs *CartService) ClearCart(userID int64) error {
	return cs.CartRepo.DeleteCart(userID)
}

func (cs *CartService) GetCart(userID int64) (*models.Cart, error) {
	cart, _ := cs.CartRepo.GetCart(userID)
	if cart == nil || len(cart.Items) == 0 {
		return nil, errors.New("cart not found")
	}
	return cart, nil
}

func (cs *CartService) GetSortedItems(cart *models.Cart) []models.CartItem {
	items := make([]models.CartItem, 0, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, *item)
	}

	sort.Slice(items, func(i, j int) bool {
		return *items[i].SKU < *items[j].SKU
	})

	return items
}
