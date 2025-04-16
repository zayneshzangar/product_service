package rest

import (
	"product_service/internal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(productHandler *ProductHandler) *mux.Router {
	r := mux.NewRouter()

	// Публичные маршруты
	r.HandleFunc("/products", productHandler.GetAllProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductByIDHandler).Methods("GET")

	// Защищённые маршруты (JWT проверка)
	authRoutes := r.PathPrefix("/").Subrouter()
	authRoutes.Use(middleware.JWTAuthMiddleware)

	// Только для admin-ролей
	adminRoutes := authRoutes.PathPrefix("/products").Subrouter()
	adminRoutes.Use(middleware.AdminOnlyMiddleware)
	adminRoutes.HandleFunc("", productHandler.CreateProductHandler).Methods("POST")
	adminRoutes.HandleFunc("/{id}", productHandler.UpdateProductHandler).Methods("PUT")
	adminRoutes.HandleFunc("/{id}", productHandler.DeleteProductHandler).Methods("DELETE")

	return r
}
