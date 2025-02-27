package main

import (
	"log"
	"net/http"
	"os"

	"product_service/internal/delivery/grpc"
	"product_service/internal/delivery/rest"
	"product_service/internal/repository"
	"product_service/internal/service/grpc_service"
	"product_service/internal/service/product_service"
	"product_service/internal/service/reservation_cleaner"

	_ "github.com/lib/pq"
)

func main() {
	// Загружаем .env
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dbType := repository.DatabaseType(os.Getenv("DB_TYPE"))
	repo, err := repository.NewDatabaseConnection(dbType)
	if err != nil {
		log.Fatal("Error creating repository: ", err)
	}

	productUseCase := product_service.NewProductUseCase(repo)
	stockUseCase := grpc_service.NewStockUseCase(repo)

	// Создаём и запускаем сервис очистки резервов
	reservationCleaner := reservation_cleaner.NewReservationCleaner(repo)
	go reservationCleaner.Start()
	// Запускаем REST API
	productHandler := rest.NewProductHandler(productUseCase)
	router := rest.NewRouter(productHandler)

	// Запускаем gRPC сервер в отдельной горутине
	go grpc.StartGRPCServer(stockUseCase)

	port := ":8080"
	log.Println("Starting REST API server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
