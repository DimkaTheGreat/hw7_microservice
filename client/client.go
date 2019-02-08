package main

import (
	"coursera_microservice/service"
	"fmt"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())

	checkError(err)
	defer conn.Close()

	BizClient := service.NewBizClient(conn)

	userName := "Pisya"

	md := metadata.New(map[string]string{
		"consumer": userName})

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err = BizClient.Check(ctx, &service.Nothing{})
	checkError(err)
	fmt.Println("Call Check from client")

	_, err = BizClient.Add(ctx, &service.Nothing{})
	checkError(err)
	fmt.Println("Call Add from client ")

	_, err = BizClient.Test(ctx, &service.Nothing{})
	checkError(err)
	fmt.Println("Call Test from client ")

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
