package service

import (
	"context"
	"product_service/internal/entity"
	"product_service/internal/productpb"
)

// ProductUseCase — основной интерфейс (для REST API)
type ProductService interface {
	CreateProduct(name, description string, price float64, stock int64) (*entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id int64) error
}

type GrpcService interface {
	GetProductStock(ctx context.Context, req *productpb.ProductStockRequest) (*productpb.ProductStockResponse, error)
	UpdateProductStock(ctx context.Context, req *productpb.UpdateProductStockRequest) (*productpb.UpdateProductStockResponse, error)
}
