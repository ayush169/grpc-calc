package main

import (
	"context"
	"fmt"
	proto "grpc-calc/proto"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedCalculatorServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &server{})
	reflection.Register(s)

	fmt.Println("Server started at :8080")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	time.Sleep(2 * time.Second)
	result := req.A + req.B
	return &proto.AddResponse{Result: result}, nil
}
