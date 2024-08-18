package main

import (
	"context"
	"fmt"
	proto "grpc-calc/proto"
	"net"

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
	// time.Sleep(2 * time.Second)
	result := req.A + req.B
	return &proto.AddResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, req *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	result := req.A - req.B
	return &proto.SubtractResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *proto.MultiplyRequest) (*proto.MultiplyResponse, error) {
	result := req.A * req.B
	return &proto.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, req *proto.DivideRequest) (*proto.DivideResponse, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}
	result := req.A / req.B
	return &proto.DivideResponse{Result: result}, nil
}
