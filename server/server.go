package main

import (
	"coursera_microservice/service"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

func (s *Server) adminInterceptor() grpc.StreamServerInterceptor {

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		md, ok := metadata.FromIncomingContext(ss.Context())
		cons := md["consumer"]
		consumer := strings.Join(cons, "")

		if info.FullMethod == "/service.Admin/Logging" {
			s.IsLogging = true

		}

		if !ok {
			fmt.Println("There is no any context for user authentication")
			return errors.New("There is no any context for user authentication")
		}

		if consumer != "stat" && strings.Contains(info.FullMethod, "Statistics") || consumer != "logger" && strings.Contains(info.FullMethod, "Logging") {
			fmt.Println("This user cant call method ", info.FullMethod)
			return errors.New("Authentication error")
		}

		return handler(srv, ss)

	}
}
func (s *Server) methodCountInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (some interface{}, err error) {
		fmt.Println("Client calling method : ", info.FullMethod)

		md, ok := metadata.FromIncomingContext(ctx)
		cons := md["consumer"]
		consumer := strings.Join(cons, "")

		if !ok {
			fmt.Println("There is no any context for user authentication")
			return nil, errors.New("There is no any context for user authentication")
		}

		if consumer == "biz_user" && strings.Contains(info.FullMethod, "Test") {
			fmt.Println("This user cant call method Test")
			return nil, errors.New("Authentication error")
		}

		if _, ok := s.Stat.ByConsumer[consumer]; ok {
			s.Stat.ByConsumer[consumer]++
		} else {
			s.Stat.ByConsumer[consumer] = 1
		}

		if _, ok := s.Stat.ByMethod[info.FullMethod]; ok {
			s.Stat.ByMethod[info.FullMethod]++

		} else {
			s.Stat.ByMethod[info.FullMethod] = 1
		}

		switch s.IsLogging {
		case false:
			return handler(ctx, req)
		case true:
			s.Event.Timestamp = time.Now().Unix()
			s.Event.Consumer = consumer
			s.Event.Method = info.FullMethod
			md, ok = metadata.FromIncomingContext(ctx)
			if !ok {
				fmt.Println("Cant find client host data")
			}
			s.Event.Host = strings.Join(md[":authority"], "")

			s.EventCh <- &s.Event

			return handler(ctx, req)
		}
		return handler(ctx, req)
	}
}

func main() {
	s := &Server{
		Event: service.Event{Timestamp: 0, Consumer: "", Method: "", Host: ""},
		Stat:  service.Stat{Timestamp: 0, ByConsumer: map[string]uint64{}, ByMethod: map[string]uint64{}},
	}
	s.init()
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	checkError(err)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(s.methodCountInterceptor()), grpc.StreamInterceptor(s.adminInterceptor()))
	service.RegisterBizServer(grpcServer, s)
	service.RegisterAdminServer(grpcServer, s)
	fmt.Println("starting server at :8081")

	grpcServer.Serve(listener)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
