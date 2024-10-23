package repositories

import (
	"awesomeProject4/internal/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryCartRepository_GetCart(t *testing.T) {
	repo := NewInMemoryCartRepository()

	userID := int64(1)
	cart := &models.Cart{
		UserID: &userID,
		Items:  make(map[int64]*models.CartItem),
	}

	err := repo.SaveCart(cart)
	assert.NoError(t, err)

	gotCart, err := repo.GetCart(userID)
	assert.NoError(t, err)
	assert.Equal(t, cart, gotCart)

	nonExistentUserID := int64(2)
	gotCart, err = repo.GetCart(nonExistentUserID)
	assert.NoError(t, err)
	assert.Nil(t, gotCart)
}

func TestInMemoryCartRepository_SaveCart(t *testing.T) {
	repo := NewInMemoryCartRepository()
	userID := int64(1)
	cart := &models.Cart{
		UserID: &userID,
		Items:  make(map[int64]*models.CartItem),
	}

	err := repo.SaveCart(cart)
	assert.NoError(t, err)

	gotCart, err := repo.GetCart(userID)
	assert.NoError(t, err)
	assert.Equal(t, cart, gotCart)

	var a int64 = 100
	var b uint16 = 2
	cart.Items[100] = &models.CartItem{SKU: &a, Count: &b}
	err = repo.SaveCart(cart)
	assert.NoError(t, err)

	gotCart, err = repo.GetCart(userID)
	assert.NoError(t, err)
	assert.Equal(t, cart, gotCart)
}

func TestInMemoryCartRepository_DeleteCart(t *testing.T) {
	repo := NewInMemoryCartRepository()

	userID := int64(1)
	cart := &models.Cart{
		UserID: &userID,
		Items:  make(map[int64]*models.CartItem),
	}

	err := repo.SaveCart(cart)
	assert.NoError(t, err)

	err = repo.DeleteCart(userID)
	assert.NoError(t, err)

	gotCart, err := repo.GetCart(userID)
	assert.NoError(t, err)
	assert.Nil(t, gotCart)
}
