package main

import (
	"context"
	"fmt"
	proto "grpc-calc/proto"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.Add(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Result)
}
