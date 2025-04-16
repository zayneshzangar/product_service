package grpc_service

import (
	"context"
	"product_service/internal/productpb"
	"product_service/internal/repository"
	"product_service/internal/service"
)

type grpcService struct {
	repo repository.ProductRepository
}

func NewGrpcService(repo repository.ProductRepository) service.GrpcService {
	return &grpcService{repo: repo}
}

func (s *grpcService) GetProductStock(ctx context.Context, req *productpb.ProductStockRequest) (*productpb.ProductStockResponse, error) {
	productIds := req.GetProductIds()
	res := &productpb.ProductStockResponse{
		StockMap: make(map[int64]*productpb.ProductStockInfo), // Инициализация карты
	}
	for _, productId := range productIds {
		product, err := s.repo.GetByID(productId)
		if err != nil {
			return nil, err
		}

		res.StockMap[productId] = &productpb.ProductStockInfo{
			Stock: product.Stock,
			Name:  product.Name,
		}
	}

	return res, nil
}

func (s *grpcService) UpdateProductStock(ctx context.Context, req *productpb.UpdateProductStockRequest) (*productpb.UpdateProductStockResponse, error) {
	updates := req.GetUpdates()
	for _, update := range updates {
		productId := update.GetProductId()
		quantity := update.GetQuantity()

		err := s.repo.UpdateStock(productId, quantity)
		if err != nil {
			return nil, err
		}
	}

	return &productpb.UpdateProductStockResponse{
		Error: "",
	}, nil
}
