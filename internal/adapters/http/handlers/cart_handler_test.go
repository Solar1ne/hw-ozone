package handlers

import (
	"awesomeProject4/internal/domain/models"
	"bytes"
	"context"
	"encoding/json"
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

func TestCartHandler_RemoveItem(t *testing.T) {
	repo := repositories.NewInMemoryCartRepository()
	productService := &StubProductService{}
	cartService := services.NewCartService(repo, productService)
	handler := NewCartHandler(cartService)

	// Предположим, что пользователь с ID 1 и товар с SKU 100 уже есть в корзине
	_ = cartService.AddItemToCart(context.Background(), 1, 100, 2)

	tests := []struct {
		name           string
		userID         string
		skuID          string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			userID:         "1",
			skuID:          "100",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid user_id",
			userID:         "abc",
			skuID:          "100",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid sku_id",
			userID:         "1",
			skuID:          "abc",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/cart/"+tt.userID+"/item/"+tt.skuID, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/cart/{user_id}/item/{sku_id}", handler.RemoveItem).Methods("DELETE")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

func TestCartHandler_ClearCart(t *testing.T) {
	repo := repositories.NewInMemoryCartRepository()
	productService := &StubProductService{}
	cartService := services.NewCartService(repo, productService)
	handler := NewCartHandler(cartService)

	// Добавляем товары в корзину пользователя
	_ = cartService.AddItemToCart(context.Background(), 1, 100, 2)
	_ = cartService.AddItemToCart(context.Background(), 1, 101, 1)

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
	}{
		{
			name:           "Valid request",
			userID:         "1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Invalid user_id",
			userID:         "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Non-existing user",
			userID:         "999",                // Предполагаем, что такого пользователя нет
			expectedStatus: http.StatusNoContent, // Если очистка пустой корзины тоже успешна
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "/cart/"+tt.userID, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/cart/{user_id}", handler.ClearCart).Methods("DELETE")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

func TestCartHandler_GetCart(t *testing.T) {
	repo := repositories.NewInMemoryCartRepository()
	productService := &StubProductService{}
	cartService := services.NewCartService(repo, productService)
	handler := NewCartHandler(cartService)

	// Добавляем товары в корзину пользователя
	_ = cartService.AddItemToCart(context.Background(), 1, 100, 2)
	_ = cartService.AddItemToCart(context.Background(), 1, 101, 1)

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid request",
			userID:         "1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"items":[{"sku_id":100,"name":"Test Product","count":2,"price":100},{"sku_id":101,"name":"Test Product","count":1,"price":100}],"total_price":300}`,
		},
		{
			name:           "Invalid user_id",
			userID:         "abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid user_id\n",
		},
		{
			name:           "Empty cart",
			userID:         "2", // Предполагаем, что у пользователя с ID 2 пустая корзина
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Cart not found\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/cart/"+tt.userID, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/cart/{user_id}", handler.GetCart).Methods("GET")
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			// Проверяем тело ответа только если статус OK
			if tt.expectedStatus == http.StatusOK {
				// Декодируем ответ
				var response CartResponse
				err = json.NewDecoder(rr.Body).Decode(&response)
				assert.NoError(t, err)

				// Вычисляем ожидаемый totalPrice и количество items
				var expectedResponse CartResponse
				_ = json.Unmarshal([]byte(tt.expectedBody), &expectedResponse)
				assert.Equal(t, expectedResponse.TotalPrice, response.TotalPrice)
				assert.Equal(t, expectedResponse.Items, response.Items)
			} else {
				assert.Equal(t, tt.expectedBody, rr.Body.String())
			}
		})
	}
}
