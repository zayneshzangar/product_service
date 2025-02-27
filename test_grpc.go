package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "product_service/internal/productpb"

	"google.golang.org/grpc"
)

func main() {
	// Устанавливаем соединение с gRPC-сервером
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Формируем запрос
	req := &pb.StockReservationRequest{
		OrderId: 123,
		Items: []*pb.OrderItem{
			{ProductId: 1, Quantity: 2},
			{ProductId: 2, Quantity: 1},
		},
	}

	// Отправляем запрос
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.CheckAndReserveStock(ctx, req)
	if err != nil {
		log.Fatalf("Ошибка вызова gRPC: %v", err)
	}

	fmt.Printf("Ответ от сервера: success=%v, message=%s\n", res.Success, res.Message)
}
