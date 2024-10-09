package repositories

import (
	"awesomeProject4/internal/domain/models"
	"sync"
)

type InMemoryCartRepository struct {
	carts map[int64]*models.Cart
	mu    sync.RWMutex
}

func NewInMemoryCartRepository() *InMemoryCartRepository {
	return &InMemoryCartRepository{
		carts: make(map[int64]*models.Cart),
	}
}

func (repo *InMemoryCartRepository) GetCart(userID int64) (*models.Cart, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if cart, exists := repo.carts[userID]; exists {
		return cart, nil
	}
	return nil, nil
}

func (repo *InMemoryCartRepository) SaveCart(cart *models.Cart) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.carts[*cart.UserID] = cart
	return nil
}

func (repo *InMemoryCartRepository) DeleteCart(userID int64) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.carts, userID)
	return nil
}
