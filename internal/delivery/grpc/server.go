package grpc

import (
	"context"
	"log"
	"net"

	"product_service/internal/productpb"
	pb "product_service/internal/productpb"
	"product_service/internal/service"

	// "google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	productpb.UnimplementedProductServiceServer
	service service.GrpcService
}

func NewGRPCServer(service service.GrpcService) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) GetProductStock(ctx context.Context, req *productpb.ProductStockRequest) (*productpb.ProductStockResponse, error) {
	return s.service.GetProductStock(ctx, req)
}

func (s *GRPCServer) UpdateProductStock(ctx context.Context, req *productpb.UpdateProductStockRequest) (*productpb.UpdateProductStockResponse, error) {
	return s.service.UpdateProductStock(ctx, req)
}

func StartGRPCServer(grpcService service.GrpcService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, NewGRPCServer(grpcService))

	// // Включаем reflection API
	// reflection.Register(grpcServer)

	log.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
