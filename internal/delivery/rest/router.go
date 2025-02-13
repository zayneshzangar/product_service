package rest

import (
	"github.com/gorilla/mux"
)

func NewRouter(productHandler *ProductHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.CreateProductHandler).Methods("POST")
	r.HandleFunc("/products", productHandler.GetAllProductsHandler).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.GetProductByIDHandler).Methods("GET")
	r.HandleFunc("/products/{id}", productHandler.UpdateProductHandler).Methods("PUT")    // <-- Добавили обновление
	r.HandleFunc("/products/{id}", productHandler.DeleteProductHandler).Methods("DELETE") // <-- Добавили удаление
	return r
}
