package rest

import (
	"encoding/json"
	"net/http"
	"product_service/internal/entity"
	"product_service/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductHandler отвечает за обработку REST-запросов
type ProductHandler struct {
	productUseCase service.ProductUseCase
}

// NewProductHandler создаёт новый обработчик
func NewProductHandler(productUseCase service.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: productUseCase}
}

// CreateProductHandler — обработчик для создания продукта
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	product, err := h.productUseCase.CreateProduct(req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductByIDHandler — обработчик для получения продукта по ID
func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing product ID", http.StatusBadRequest)
		return
	}

	// Конвертируем строку в int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	// Получаем продукт из usecase
	product, err := h.productUseCase.GetProductByID(id)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	// Отправляем ответ в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetAllProductsHandler — обработчик для получения всех продуктов
func (h *ProductHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.productUseCase.GetAllProducts()
	if err != nil {
		http.Error(w, "failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// UpdateProductHandler — обновление товара по ID
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	product.ID = id // Устанавливаем ID из URL

	err = h.productUseCase.UpdateProduct(&product)
	if err != nil {
		http.Error(w, "failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "product updated successfully"})
}

// DeleteProductHandler — удаление товара по ID
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing product ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.productUseCase.DeleteProduct(id)
	if err != nil {
		http.Error(w, "failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "product deleted successfully"})
}
