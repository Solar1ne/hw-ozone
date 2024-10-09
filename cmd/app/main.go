package main

import (
	"awesomeProject4/internal/infrastructure/clients"
	"log"
	"net/http"

	"awesomeProject4/internal/adapters/http/handlers"
	"awesomeProject4/internal/adapters/http/middlewares"
	"awesomeProject4/internal/adapters/repositories"
	"awesomeProject4/internal/domain/services"

	"github.com/gorilla/mux"
)

func main() {
	cartRepo := repositories.NewInMemoryCartRepository()

	cartService := services.NewCartService(cartRepo, clients.NewProductClient("http://localhost:8082", "testtoken", &http.Client{})) // Передаем nil вместо продукта-клиента
	cartHandler := handlers.NewCartHandler(cartService)

	// Настройка маршрутов
	router := mux.NewRouter()
	router.Use(middlewares.LoggingMiddleware)
	router.HandleFunc("/user/{user_id}/cart/{sku_id}", cartHandler.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/user/{user_id}/cart/{sku_id}", cartHandler.RemoveItem).Methods(http.MethodDelete)
	router.HandleFunc("/user/{user_id}/cart", cartHandler.ClearCart).Methods(http.MethodDelete)
	router.HandleFunc("/user/{user_id}/cart", cartHandler.GetCart).Methods(http.MethodGet)

	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
