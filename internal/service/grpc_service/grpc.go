package grpc_service

import (
	"context"
	"fmt"

	"product_service/internal/productpb"
	"product_service/internal/repository"
	"product_service/internal/service"
)

type stockUseCase struct {
	repo repository.ProductRepository
}

func NewStockUseCase(repo repository.ProductRepository) service.StockUseCase {
	return &stockUseCase{repo: repo}
}

func (u *stockUseCase) CheckAndReserveStock(ctx context.Context, req *productpb.StockReservationRequest) (*productpb.StockReservationResponse, error) {
	tx, err := u.repo.BeginTransaction()
	if err != nil {
		return &productpb.StockReservationResponse{Success: false, Message: "DB transaction error"}, err
	}
	defer tx.Rollback()

	for _, item := range req.Items {
		product, err := u.repo.GetProductByID(ctx, item.ProductId)
		if err != nil {
			return &productpb.StockReservationResponse{Success: false, Message: fmt.Sprintf("Product %d not found", item.ProductId)}, nil
		}

		if product.Stock < int(item.Quantity) {
			return &productpb.StockReservationResponse{Success: false, Message: fmt.Sprintf("Not enough stock for product %d", item.ProductId)}, nil
		}

		err = u.repo.ReserveStock(ctx, tx, req.OrderId, item.ProductId, int(item.Quantity))
		if err != nil {
			return &productpb.StockReservationResponse{Success: false, Message: "Failed to reserve stock"}, err
		}
	}

	if err := tx.Commit(); err != nil {
		return &productpb.StockReservationResponse{Success: false, Message: "Transaction commit failed"}, err
	}

	return &productpb.StockReservationResponse{Success: true, Message: "Stock reserved successfully"}, nil
}
