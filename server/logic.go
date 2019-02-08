package main

import (
	"coursera_microservice/service"
	"fmt"
	"sync"

	"golang.org/x/net/context"
)

//Client ...
type Client struct {
	name  string
	token string
	m     sync.RWMutex
}

//NewClient ...
func NewClient() *Client {
	return &Client{}

}

//Check ...
func (c *Client) Check(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {
	fmt.Println("Call Check ")
	return &service.Nothing{Dummy: true}, nil

}

//Add ...
func (c *Client) Add(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {
	fmt.Println("Call Add ")
	return &service.Nothing{Dummy: true}, nil

}

//Test ...
func (c *Client) Test(ctx context.Context, in *service.Nothing) (out *service.Nothing, err error) {
	fmt.Println("Call Test ")
	return &service.Nothing{Dummy: true}, nil

}
