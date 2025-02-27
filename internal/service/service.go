package service

import (
	"context"

	"product_service/internal/entity"
	pb "product_service/internal/productpb"
)

// ProductUseCase — основной интерфейс (для REST API)
type ProductUseCase interface {
	CreateProduct(name, description string, price float64, stock int) (*entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int64) error
}

// StockUseCase — интерфейс для работы со складом (используется в gRPC)
type StockUseCase interface {
	CheckAndReserveStock(ctx context.Context, req *pb.StockReservationRequest) (*pb.StockReservationResponse, error)
}
