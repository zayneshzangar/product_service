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

	// добавляем пакет для CORS

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
	defer repo.Close()

	productService := product_service.NewProductService(repo)
	stockService := grpc_service.NewGrpcService(repo)

	// Запускаем REST API
	productHandler := rest.NewProductHandler(productService)
	router := rest.NewRouter(productHandler)

	// Обертываем router в CORS
	corsHandler := rest.UserCors(router)

	// Запускаем gRPC сервер в отдельной горутине
	go grpc.StartGRPCServer(stockService)

	port := ":8080"
	log.Println("Starting REST API server on", port)
	log.Fatal(http.ListenAndServe(port, corsHandler)) // Используем CORS обработчик
}	
