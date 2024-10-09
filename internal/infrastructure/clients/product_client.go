package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"awesomeProject4/internal/domain/models"
)

type ProductClient struct {
	BaseURL string
	Client  *http.Client
	Token   string
}

func NewProductClient(baseURL, token string, client *http.Client) *ProductClient {
	return &ProductClient{
		BaseURL: baseURL,
		Client:  client,
		Token:   token,
	}
}

func (pc *ProductClient) GetProduct(ctx context.Context, sku int64) (*models.Product, error) {
	url := fmt.Sprintf("%s/get_product", pc.BaseURL)

	requestBody, err := json.Marshal(map[string]interface{}{
		"token": pc.Token,
		"sku":   sku,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := pc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var response struct {
			Name  string `json:"name"`
			Price uint32 `json:"price"`
		}
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return nil, err
		}

		product := &models.Product{
			SKU:   sku,
			Name:  response.Name,
			Price: response.Price,
		}
		return product, nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("product not found")
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	return nil, fmt.Errorf("failed to get product info: %s", string(bodyBytes))
}
