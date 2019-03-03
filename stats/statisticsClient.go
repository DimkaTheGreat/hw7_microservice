package main

import (
	"coursera_microservice/service"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	checkError1(err)
	defer conn.Close()
	ctx := context.Background()

	md := metadata.Pairs("consumer", "stat")
	newCtx := metadata.NewOutgoingContext(ctx, md)

	statClient := service.NewAdminClient(conn)
	interval := &service.StatInterval{IntervalSeconds: 5}

	client, err := statClient.Statistics(newCtx, interval)

	for {

		statistic, err := client.Recv()
		if err == io.EOF {

		}
		if err != nil {
			fmt.Println(err)
			break
		}
		statistic.Timestamp = time.Now().Unix()
		fmt.Println(statistic)

	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
