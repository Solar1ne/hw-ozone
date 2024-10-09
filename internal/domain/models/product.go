package models

type Product struct {
	SKU   int64  `json:"sku"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}
