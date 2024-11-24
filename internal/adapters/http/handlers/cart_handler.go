package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"awesomeProject4/internal/domain/services"
	"awesomeProject4/internal/infrastructure/validation"

	"github.com/gorilla/mux"
)

type CartHandler struct {
	Service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{
		Service: service,
	}
}

func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	skuIDStr := vars["sku_id"]

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	skuID, err := strconv.ParseInt(skuIDStr, 10, 64)
	if err != nil || skuID <= 0 {
		http.Error(w, "Invalid sku_id", http.StatusBadRequest)
		return
	}

	var req struct {
		Count uint16 `json:"count" validate:"required,min=1"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := validation.Validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.AddItemToCart(ctx, userID, skuID, req.Count)
	if err != nil {
		if err.Error() == "product not found" {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]
	skuIDStr := vars["sku_id"]

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}
	skuID, err := strconv.ParseInt(skuIDStr, 10, 64)
	if err != nil || skuID <= 0 {
		http.Error(w, "Invalid sku_id", http.StatusBadRequest)
		return
	}

	err = h.Service.RemoveItemFromCart(userID, skuID)
	if err != nil {
		http.Error(w, "Failed to remove item from cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	err = h.Service.ClearCart(userID)
	if err != nil {
		http.Error(w, "Failed to clear cart", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	cart, err := h.Service.GetCart(userID)
	if err != nil {
		http.Error(w, "Cart not found", http.StatusNotFound)
		return
	}

	items := h.Service.GetSortedItems(cart)

	var totalPrice uint32
	var itemsResponse []CartItemResponse
	for _, item := range items {
		itemsResponse = append(itemsResponse, CartItemResponse{
			SKUID: *item.SKU,
			Name:  item.Name,
			Count: *item.Count,
			Price: *item.Price,
		})
		totalPrice += uint32(*(item.Count)) * (*item.Price)
	}

	res := CartResponse{
		Items:      itemsResponse,
		TotalPrice: totalPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *Server) Checkout(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	reqData, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

}
