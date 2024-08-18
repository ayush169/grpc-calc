package main

import (
	"context"
	"fmt"
	proto "grpc-calc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCalculatorClient(conn)

	req := &proto.AddRequest{A: 10, B: 40}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	res, err := client.Add(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Addition result:", res.Result)

	req2 := &proto.SubtractRequest{A: -50, B: 20}
	res2, err := client.Subtract(context.Background(), req2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Subtraction result:", res2.Result)

	req3 := &proto.MultiplyRequest{A: 0, B: 8}
	res3, err := client.Multiply(context.Background(), req3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Multiplication result:", res3.Result)

	req4 := &proto.DivideRequest{A: 100, B: 11}
	res4, err := client.Divide(context.Background(), req4)
	if err != nil {
		panic(err)
	}
	fmt.Println("Division result:", res4.Result)
}
