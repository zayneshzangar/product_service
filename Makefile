proto:
	protoc --go_out=. --go-grpc_out=. proto/product.proto

run:
	go run cmd/main.go

goGet:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
