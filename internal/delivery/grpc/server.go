package grpc

import (
	"context"
	"log"
	"net"

	pb "product_service/internal/productpb"
	"product_service/internal/service"
	// "google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedProductServiceServer
	service service.StockUseCase
}

func NewGRPCServer(service service.StockUseCase) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) CheckAndReserveStock(ctx context.Context, req *pb.StockReservationRequest) (*pb.StockReservationResponse, error) {
	return s.service.CheckAndReserveStock(ctx, req)
}

func StartGRPCServer(stockService service.StockUseCase) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, NewGRPCServer(stockService))

	// // Включаем reflection API
	// reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
