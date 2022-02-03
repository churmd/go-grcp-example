package main

import (
	"context"
	"log"
	"net"

	"github.com/churmd/go-grcp-example/echo"
	"google.golang.org/grpc"
)

type Server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *Server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("Received request with body %s\n", req.Body)
	return &echo.EchoResponse{Request: req}, nil
}

func main() {
	log.Println("Starting up grpc echo server")

	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("failed to listen on tcp port")
	}

	grcpServer := grpc.NewServer()
	echoServer := Server{}
	echo.RegisterEchoServiceServer(grcpServer, &echoServer)
	err = grcpServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve grcp")
	}
}
