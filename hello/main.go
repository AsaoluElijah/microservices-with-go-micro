package main

import (
	"context"
	"log"

	"go-micro.dev/v4"
)

// Define the request and response structure
type HelloRequest struct{}
type HelloResponse struct {
	Message string
}

// Define a handler
type Greeter struct{}

// Define a method on the handler
func (g *Greeter) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	rsp.Message = "Hello, World!"
	return nil
}

func main2() {
	// Create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
	)

	// Initialize the service
	service.Init()

	// Register handler
	if err := service.Server().Handle(
		service.Server().NewHandler(&Greeter{}),
	); err != nil {
		log.Fatal(err)
	}

	// Start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
