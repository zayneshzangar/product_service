package grpc_service

import (
	"context"
	"errors"
	"testing"

	"product_service/internal/entity"
	"product_service/internal/productpb"
	"product_service/internal/repository/mocks"

	"go.uber.org/mock/gomock" // Обновлённый импорт
)

func TestGrpcService_GetProductStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewGrpcService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		req := &productpb.ProductStockRequest{ProductIds: []int64{1, 2}}
		expectedProduct1 := &entity.Product{ID: 1, Name: "Product 1", Stock: 10}
		expectedProduct2 := &entity.Product{ID: 2, Name: "Product 2", Stock: 20}
		mockRepo.EXPECT().GetByID(int64(1)).Return(expectedProduct1, nil)
		mockRepo.EXPECT().GetByID(int64(2)).Return(expectedProduct2, nil)

		// Выполнение
		res, err := service.GetProductStock(context.Background(), req)

		// Проверка
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(res.StockMap) != 2 {
			t.Errorf("expected 2 products, got %d", len(res.StockMap))
		}
		if res.StockMap[1].Name != "Product 1" || res.StockMap[1].Stock != 10 {
			t.Errorf("expected Product 1 with stock 10, got %v", res.StockMap[1])
		}
		if res.StockMap[2].Name != "Product 2" || res.StockMap[2].Stock != 20 {
			t.Errorf("expected Product 2 with stock 20, got %v", res.StockMap[2])
		}
	})

	t.Run("ProductNotFound", func(t *testing.T) {
		// Подготовка
		req := &productpb.ProductStockRequest{ProductIds: []int64{999}}
		mockRepo.EXPECT().GetByID(int64(999)).Return(nil, errors.New("product not found"))

		// Выполнение
		res, err := service.GetProductStock(context.Background(), req)

		// Проверка
		if err == nil || err.Error() != "product not found" {
			t.Errorf("expected error 'product not found', got %v", err)
		}
		if res != nil {
			t.Errorf("expected nil response, got %v", res)
		}
	})
}

func TestGrpcService_UpdateProductStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := NewGrpcService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Подготовка
		req := &productpb.UpdateProductStockRequest{
			Updates: []*productpb.UpdateProductStockRequest_StockUpdate {
				{ProductId: 1, Quantity: 5},
				{ProductId: 2, Quantity: 10},
			},
		}
		mockRepo.EXPECT().UpdateStock(int64(1), int64(5)).Return(nil)
		mockRepo.EXPECT().UpdateStock(int64(2), int64(10)).Return(nil)

		// Выполнение
		res, err := service.UpdateProductStock(context.Background(), req)

		// Проверка
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if res.Error != "" {
			t.Errorf("expected empty error, got %s", res.Error)
		}
	})

	t.Run("RepositoryError", func(t *testing.T) {
		// Подготовка
		req := &productpb.UpdateProductStockRequest{
			Updates: []*productpb.UpdateProductStockRequest_StockUpdate{
				{ProductId: 1, Quantity: 5},
			},
		}
		mockRepo.EXPECT().UpdateStock(int64(1), int64(5)).Return(errors.New("database error"))

		// Выполнение
		res, err := service.UpdateProductStock(context.Background(), req)

		// Проверка
		if err == nil || err.Error() != "database error" {
			t.Errorf("expected error 'database error', got %v", err)
		}
		if res != nil {
			t.Errorf("expected nil response, got %v", res)
		}
	})
}
