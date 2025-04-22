package product_service

import (
	"errors"
	"testing"
	"time"

	"go.uber.org/mock/gomock" // Обновлённый импорт
	"product_service/internal/entity"
	"product_service/internal/repository/mocks"
)

func TestProductService_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		expectedProduct := &entity.Product{
			Name:        "Test Product",
			Description: "Description",
			Price:       99.99,
			Stock:       10,
			CreatedAt:   time.Now().Truncate(time.Second),
			UpdatedAt:   time.Now().Truncate(time.Second),
		}
		mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(p *entity.Product) error {
			p.ID = 1 // Симулируем присвоение ID
			return nil
		})

		// Выполнение
		product, err := service.CreateProduct("Test Product", "Description", 99.99, 10)

		// Проверка
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if product.Name != expectedProduct.Name || product.Price != expectedProduct.Price || product.Stock != expectedProduct.Stock {
			t.Errorf("expected product %v, got %v", expectedProduct, product)
		}
	})

	t.Run("InvalidData_NameEmpty", func(t *testing.T) {
		// Выполнение
		product, err := service.CreateProduct("", "Description", 99.99, 10)

		// Проверка
		if err == nil || err.Error() != "invalid product data" {
			t.Errorf("expected error 'invalid product data', got %v", err)
		}
		if product != nil {
			t.Errorf("expected nil product, got %v", product)
		}
	})

	t.Run("InvalidData_NegativePrice", func(t *testing.T) {
		// Выполнение
		product, err := service.CreateProduct("Test Product", "Description", -1, 10)

		// Проверка
		if err == nil || err.Error() != "invalid product data" {
			t.Errorf("expected error 'invalid product data', got %v", err)
		}
		if product != nil {
			t.Errorf("expected nil product, got %v", product)
		}
	})

	t.Run("RepositoryError", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().Create(gomock.Any()).Return(errors.New("database error"))

		// Выполнение
		product, err := service.CreateProduct("Test Product", "Description", 99.99, 10)

		// Проверка
		if err == nil || err.Error() != "database error" {
			t.Errorf("expected error 'database error', got %v", err)
		}
		if product != nil {
			t.Errorf("expected nil product, got %v", product)
		}
	})
}

func TestProductService_GetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		expectedProduct := &entity.Product{
			ID:          1,
			Name:        "Test Product",
			Description: "Description",
			Price:       99.99,
			Stock:       10,
		}
		mockRepo.EXPECT().GetByID(int64(1)).Return(expectedProduct, nil)

		// Выполнение
		product, err := service.GetProductByID(1)

		// Проверка
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if product != expectedProduct {
			t.Errorf("expected product %v, got %v", expectedProduct, product)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		// Подготовка
		mockRepo.EXPECT().GetByID(int64(999)).Return(nil, errors.New("product not found"))

		// Выполнение
		product, err := service.GetProductByID(999)

		// Проверка
		if err == nil || err.Error() != "product not found" {
			t.Errorf("expected error 'product not found', got %v", err)
		}
		if product != nil {
			t.Errorf("expected nil product, got %v", product)
		}
	})
}

func TestProductService_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewProductService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		product := &entity.Product{
			ID:          1,
			Name:        "Updated Product",
			Description: "Updated Description",
			Price:       149.99,
			Stock:       20,
		}
		mockRepo.EXPECT().Update(gomock.Any()).Return(nil)

		// Выполнение
		err := service.UpdateProduct(product)

		// Проверка
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("InvalidID", func(t *testing.T) {
		// Подготовка
		product := &entity.Product{
			ID:          0,
			Name:        "Updated Product",
			Description: "Updated Description",
			Price:       149.99,
			Stock:       20,
		}

		// Выполнение
		err := service.UpdateProduct(product)

		// Проверка
		if err == nil || err.Error() != "invalid product ID" {
			t.Errorf("expected error 'invalid product ID', got %v", err)
		}
	})

	t.Run("RepositoryError", func(t *testing.T) {
		// Подготовка
		product := &entity.Product{
			ID:          1,
			Name:        "Updated Product",
			Description: "Updated Description",
			Price:       149.99,
			Stock:       20,
		}
		mockRepo.EXPECT().Update(gomock.Any()).Return(errors.New("database error"))

		// Выполнение
		err := service.UpdateProduct(product)

		// Проверка
		if err == nil || err.Error() != "database error" {
			t.Errorf("expected error 'database error', got %v", err)
		}
	})
}