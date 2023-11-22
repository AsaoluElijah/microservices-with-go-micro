package main

import (
	"context"
	"fmt"

	"go-micro.dev/v4"
)

type Greeter struct{}

type HelloRequest struct {
	// Define the request structure
}

type HelloResponse struct {
	// Define the response structure
}

func (g *Greeter) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	fmt.Println("Hello service was called")
	// Business logic goes here...
	return nil
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("hello"),
		micro.WrapHandler(AuthMiddleware()),
	)

	// Initialize the service
	service.Init()

	// Register handler
	if err := micro.RegisterHandler(service.Server(), new(Greeter)); err != nil {
		fmt.Println(err)
		return
	}

	// Start the service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
