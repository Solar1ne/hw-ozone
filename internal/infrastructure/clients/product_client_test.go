package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"awesomeProject4/internal/domain/models"

	"github.com/stretchr/testify/assert"
)

func TestProductClient_GetProduct(t *testing.T) {
	tests := []struct {
		name            string
		sku             int64
		serverResponse  string
		serverStatus    int
		expectedProduct *models.Product
		expectedError   string
	}{
		{
			name:           "Product found",
			sku:            123,
			serverResponse: `{"name": "Test Product", "price": 100}`,
			serverStatus:   http.StatusOK,
			expectedProduct: &models.Product{
				SKU:   123,
				Name:  "Test Product",
				Price: 100,
			},
			expectedError: "",
		},
		{
			name:            "Product not found",
			sku:             456,
			serverResponse:  `{"error": "not found"}`,
			serverStatus:    http.StatusNotFound,
			expectedProduct: nil,
			expectedError:   "product not found",
		},
		{
			name:            "Server error",
			sku:             789,
			serverResponse:  `{"error": "internal server error"}`,
			serverStatus:    http.StatusInternalServerError,
			expectedProduct: nil,
			expectedError:   "failed to get product info: {\"error\": \"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/get_product", r.URL.Path)
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

				w.WriteHeader(tt.serverStatus)
				w.Write([]byte(tt.serverResponse))
			}))

			client := NewProductClient(server.URL, "test-token", server.Client())

			product, err := client.GetProduct(context.Background(), tt.sku)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedProduct, product)
			}
		})
	}
}
