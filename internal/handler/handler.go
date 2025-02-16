package handler

import (
	"avito-shop-service/internal/app"
	"encoding/json"
	"net/http"
)

// Структура для обработки HTTP запросов
type Handler struct {
	ProductService *app.ProductService
}

// Новый обработчик
func NewHandler(productService *app.ProductService) *Handler {
	return &Handler{ProductService: productService}
}

// Обработчик для получения всех продуктов
func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProducts(r.Context()) // передаем контекст из запроса
	if err != nil {
		http.Error(w, "Unable to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Unable to encode products", http.StatusInternalServerError)
	}
}
