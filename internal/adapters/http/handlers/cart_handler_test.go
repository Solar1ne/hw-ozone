package handlers

import (
	"awesomeProject4/internal/domain/models"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"awesomeProject4/internal/adapters/repositories"
	"awesomeProject4/internal/domain/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type StubProductService struct{}

func (s *StubProductService) GetProduct(ctx context.Context, sku int64) (*models.Product, error) {
	return &models.Product{
		SKU:   sku,
		Name:  "Test Product",
		Price: 100,
	}, nil
}

func TestCartHandler_AddItem(t *testing.T) {
	repo := repositories.NewInMemoryCartRepository()
	productService := &StubProductService{}
	cartService := services.NewCartService(repo, productService)
	handler := NewCartHandler(cartService)

	tests := []struct {
		name           string
		userID         string
		skuID          string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			userID:         "1",
			skuID:          "100",
			requestBody:    `{"count": 2}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid user_id",
			userID:         "abc",
			skuID:          "100",
			requestBody:    `{"count": 2}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid sku_id",
			userID:         "1",
			skuID:          "abc",
			requestBody:    `{"count": 2}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid request body",
			userID:         "1",
			skuID:          "100",
			requestBody:    `{"count": "two"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid count (zero)",
			userID:         "1",
			skuID:          "100",
			requestBody:    `{"count": 0}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/cart/"+tt.userID+"/item/"+tt.skuID, bytes.NewBufferString(tt.requestBody))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/cart/{user_id}/item/{sku_id}", handler.AddItem).Methods("POST")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
