package main

import (
	"coursera_microservice/service"
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

//Server ...
type Server struct {
	service.Event
	service.Stat
	EventCh   chan *service.Event
	IsLogging bool
	QueryCh   chan struct{}
	WG        sync.WaitGroup
}

/*func (s *Server) init() {
	s.QueryCh = make(chan struct{}, 1)

}*/

//Check ...
func (s *Server) Check(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {

	return &service.Nothing{Dummy: true}, nil

}

//Add ...
func (s *Server) Add(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {

	return &service.Nothing{Dummy: true}, nil

}

//Test ...
func (s *Server) Test(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {

	return &service.Nothing{Dummy: true}, nil
}

// Statistics ...
func (s *Server) Statistics(in *service.StatInterval, outstream service.Admin_StatisticsServer) (err error) {
	for {
		time.Sleep(time.Duration(in.IntervalSeconds) * time.Second)
		err = outstream.Send(&s.Stat)

		if err != nil {
			fmt.Println("service closed")
			break
		}

	}

	return err
}

//Logging ...
func (s *Server) Logging(in *service.Nothing, outstream service.Admin_LoggingServer) (err error) {
	go func() {

		for {
			<-s.QueryCh
			fmt.Println(s.Event)
			err = outstream.Send(&s.Event)
			if err != nil {
				fmt.Println("service closed", err)
				return

			}

		}

	}()
	return nil
}
