package main

import (
	"context"
	"coursera_microservice/service"
	"fmt"
	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()

	md := metadata.Pairs("consumer", "logger")
	newCtx := metadata.NewOutgoingContext(ctx, md)

	logClient := service.NewAdminClient(conn)

	client, err := logClient.Logging(newCtx, &service.Nothing{})

	for {

		event, err := client.Recv()
		if err == io.EOF {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(event)

	}

}
