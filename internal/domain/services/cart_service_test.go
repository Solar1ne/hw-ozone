package services

import (
	"awesomeProject4/internal/domain/models"
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestCartService_AddItemToCart(t *testing.T) {
	mc := minimock.NewController(t)
	mockRepo := NewCartRepositoryMock(mc)
	mockProductService := NewProductServiceMock(mc)

	cs := NewCartService(mockRepo, mockProductService)
	ctx := context.Background()
	userID := int64(1)
	skuID := int64(100)
	count := uint16(2)

	product := &models.Product{
		SKU:   skuID,
		Name:  "Test Product",
		Price: 10.0,
	}

	mockProductService.GetProductMock.Expect(ctx, skuID).Return(product, nil)
	mockRepo.GetCartMock.Expect(userID).Return(nil, nil)
	mockRepo.SaveCartMock.Expect(&models.Cart{
		UserID: &userID,
		Items: map[int64]*models.CartItem{
			skuID: {
				SKU:   &skuID,
				Name:  product.Name,
				Count: &count,
				Price: &product.Price,
			},
		},
	}).Return(nil)

	err := cs.AddItemToCart(ctx, userID, skuID, count)
	assert.NoError(t, err)

}

func TestCartService_RemoveItemFromCart(t *testing.T) {
	mc := minimock.NewController(t)
	mockRepo := NewCartRepositoryMock(mc)
	mockProductService := NewProductServiceMock(mc)

	cs := NewCartService(mockRepo, mockProductService)
	userID := int64(1)
	skuID := int64(100)
	cart := &models.Cart{
		UserID: &userID,
		Items: map[int64]*models.CartItem{
			skuID: {SKU: &skuID, Count: new(uint16)},
		},
	}

	mockRepo.GetCartMock.Expect(userID).Return(cart, nil)
	//mockRepo.SaveCartMock.Expect(cart).Return(nil)

	err := cs.RemoveItemFromCart(userID, skuID)
	assert.NoError(t, err)
}

func TestCartService_ClearCart(t *testing.T) {
	mc := minimock.NewController(t)
	mockRepo := NewCartRepositoryMock(mc)
	mockProductService := NewProductServiceMock(mc)

	cs := NewCartService(mockRepo, mockProductService)
	userID := int64(1)
	mockRepo.DeleteCartMock.Expect(userID).Return(nil)
	err := cs.ClearCart(userID)
	assert.NoError(t, err)
}

func TestCartService_GetCart(t *testing.T) {
	mc := minimock.NewController(t)

	mockRepo := NewCartRepositoryMock(mc)
	mockProductService := NewProductServiceMock(mc)
	cs := NewCartService(mockRepo, mockProductService)

	userID := int64(1)

	cart := &models.Cart{
		UserID: &userID,
		Items: map[int64]*models.CartItem{
			100: {SKU: new(int64)},
		},
	}

	mockRepo.GetCartMock.Expect(userID).Return(cart, nil)
	gotCart, err := cs.GetCart(userID)
	assert.NoError(t, err)
	assert.Equal(t, cart, gotCart)
}
