package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"product_service/internal/entity"
	"product_service/internal/service/mocks"

	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
)

func TestProductHandler_CreateProductHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProductService(ctrl)
	handler := NewProductHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		reqBody := struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
			Stock       int64   `json:"stock"`
		}{
			Name:        "Test Product",
			Description: "Description",
			Price:       99.99,
			Stock:       10,
		}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
		rr := httptest.NewRecorder()

		expectedProduct := &entity.Product{
			ID:          1,
			Name:        "Test Product",
			Description: "Description",
			Price:       99.99,
			Stock:       10,
		}
		mockService.EXPECT().CreateProduct("Test Product", "Description", 99.99, int64(10)).Return(expectedProduct, nil) // Исправлено: int64(10)

		// Выполнение
		handler.CreateProductHandler(rr, req)

		// Проверка
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("expected status %v, got %v", http.StatusCreated, status)
		}
		var response entity.Product
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
		if response.Name != expectedProduct.Name || response.Price != expectedProduct.Price {
			t.Errorf("expected product %v, got %v", expectedProduct, response)
		}
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Подготовка
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader([]byte("invalid json")))
		rr := httptest.NewRecorder()

		// Выполнение
		handler.CreateProductHandler(rr, req)

		// Проверка
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, status)
		}
	})
}

func TestProductHandler_GetProductByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProductService(ctrl)
	handler := NewProductHandler(mockService)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		expectedProduct := &entity.Product{
			ID:          1,
			Name:        "Test Product",
			Description: "Description",
			Price:       99.99,
			Stock:       10,
		}
		mockService.EXPECT().GetProductByID(int64(1)).Return(expectedProduct, nil)

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		rr := httptest.NewRecorder()
		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		// Выполнение
		handler.GetProductByIDHandler(rr, req)

		// Проверка
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected status %v, got %v", http.StatusOK, status)
		}
		var response entity.Product
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}
		if response.Name != expectedProduct.Name || response.Price != expectedProduct.Price {
			t.Errorf("expected product %v, got %v", expectedProduct, response)
		}
	})

	t.Run("InvalidID", func(t *testing.T) {
		// Подготовка
		req := httptest.NewRequest(http.MethodGet, "/products/invalid", nil)
		rr := httptest.NewRecorder()
		vars := map[string]string{"id": "invalid"}
		req = mux.SetURLVars(req, vars)

		// Выполнение
		handler.GetProductByIDHandler(rr, req)

		// Проверка
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("expected status %v, got %v", http.StatusBadRequest, status)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		// Подготовка
		mockService.EXPECT().GetProductByID(int64(999)).Return(nil, errors.New("product not found"))

		req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
		rr := httptest.NewRecorder()
		vars := map[string]string{"id": "999"}
		req = mux.SetURLVars(req, vars)

		// Выполнение
		handler.GetProductByIDHandler(rr, req)

		// Проверка
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("expected status %v, got %v", http.StatusNotFound, status)
		}
	})
}
