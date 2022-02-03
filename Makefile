gen-grpc:
	protoc --go_out=echo --go_opt=paths=source_relative --go-grpc_out=echo --go-grpc_opt=paths=source_relative echo.proto 

format:
	go fmt ./...

server:
	go run server.go

client: 
	go run client.go