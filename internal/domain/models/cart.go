package models

type CartItem struct {
	SKU   *int64  `json:"sku_id"`
	Name  string  `json:"name"`
	Count *uint16 `json:"count"`
	Price *uint32 `json:"price"`
}

type Cart struct {
	UserID *int64
	Items  map[int64]*CartItem // Ключ - SKU товара
}

type OrderID int64

type OrderStatus string
