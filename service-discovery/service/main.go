package main

import (
	"context"
	"fmt"

	"go-micro.dev/v4"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *interface{}, rsp *interface{}) error {
	fmt.Println("Service A was called")
	return nil
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("serviceA"),
	)

	// Initialize the service
	service.Init()

	// Register handler
	micro.RegisterHandler(service.Server(), new(Greeter))

	// Start the service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
