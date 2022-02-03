package main

import (
	"context"
	"log"

	"github.com/churmd/go-grcp-example/echo"
	"google.golang.org/grpc"
)

func main() {
	var grpcCon *grpc.ClientConn
	grpcCon, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to grcp server\n")
	}
	defer grpcCon.Close()

	echoClient := echo.NewEchoServiceClient(grpcCon)

	req := echo.EchoRequest{Body: "Hello"}
	resp, err := echoClient.Echo(context.Background(), &req)
	if err != nil {
		log.Fatalf("received an error when calling Echo on the echo client: %w\n", err)
	}

	log.Printf("%+v\n", resp)
}
