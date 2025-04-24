make proto --always-make

go get github.com/golang-jwt/jwt/v5
go install go.uber.org/mock/mockgen@latest
mockgen --version

mockgen -source=internal/repository/product_repository.go -destination=internal/repository/mocks/mock_product_repository.go -package=mocks
mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock_product_service.go -package=mocks


// TEST
go clean -testcache
go test ./internal/service/product_service
go test ./internal/service/grpc_service
go test ./internal/delivery/rest

go test -cover ./internal/service/product_service
go test -cover ./internal/service/grpc_service
go test -cover ./internal/delivery/rest

// C выводом
go test -v ./internal/service/product_service

// Формирование отчета покрытии
go test -coverprofile=cover.out ./internal/service/product_service
go tool cover -html=cover.out -o cover.html
