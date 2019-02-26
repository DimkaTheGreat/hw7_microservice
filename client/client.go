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

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	BizClient := service.NewBizClient(conn)

	consumerName := "biz_user"

	ctx := metadata.AppendToOutgoingContext(context.Background(), "consumer", consumerName)

	_, err = BizClient.Check(ctx, &service.Nothing{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Call Check from client")
	}

	_, err = BizClient.Add(ctx, &service.Nothing{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Call Add from client")
	}

	/*_, err = BizClient.Test(ctx, &service.Nothing{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Call Test from client")
	}*/

}
