package main

import (
	"log"
	"net/http"
	"os"

	"product_service/internal/delivery/rest"
	"product_service/internal/repository"
	"product_service/internal/usecase"
)

func main() {
	// Загружаем .env
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Выбираем базу данных
	dbType := repository.DatabaseType(os.Getenv("DB_TYPE"))
	repo, err := repository.NewDatabaseConnection(dbType)
	if err != nil {
		log.Fatal("Error creating repository: ", err)
	}

	// Создаём usecase
	productUseCase := usecase.NewProductUseCase(repo)

	// Создаём REST handler
	productHandler := rest.NewProductHandler(productUseCase)

	// Создаём роутер и запускаем сервер
	router := rest.NewRouter(productHandler)
	port := ":8080"
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
