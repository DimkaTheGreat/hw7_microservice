package main

import (
	"coursera_microservice/service"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

//Stats ...
var Stats = &service.Stat{
	Timestamp:  time.Now().Unix(),
	ByMethod:   make(map[string]uint64),
	ByConsumer: make(map[string]uint64),
}

func methodCountInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	reply, err := handler(ctx, req)
	Stats.Timestamp = time.Now().Unix()

	md, ok := metadata.FromIncomingContext(ctx)
	cons := md["consumer"]
	consumer := strings.Join(cons, "")

	if !ok {
		fmt.Println("There is no any context")
	}

	if _, ok := Stats.ByConsumer[consumer]; ok {
		Stats.ByConsumer[consumer]++
	} else {
		Stats.ByConsumer[consumer] = 1
	}

	if _, ok := Stats.ByMethod[info.FullMethod]; ok {
		Stats.ByMethod[info.FullMethod]++

	} else {
		Stats.ByMethod[info.FullMethod] = 1
	}
	fmt.Println(Stats)

	return reply, err
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	checkError(err)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(methodCountInterceptor))
	service.RegisterBizServer(grpcServer, NewClient())
	fmt.Println("starting server at :8081")

	/*go func() {
		for {
			time.Sleep(time.Second * 10)
			fmt.Println(Stats)
		}
	}()*/

	grpcServer.Serve(listener)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
