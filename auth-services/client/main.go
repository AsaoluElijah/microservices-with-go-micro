package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
)

type HelloRequest struct {
}

type HelloResponse struct {
}

func main() {
	// Initialize a new service client
	service := micro.NewService(micro.Name("hello.client"))
	service.Init()

	// Create a context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": "valid-token", // Normally, this would be a JWT token
	})

	// Create a request and response
	req := client.NewRequest("hello", "Greeter.Hello", &HelloRequest{})
	rsp := &HelloResponse{}

	// Call the service
	if err := service.Client().Call(ctx, req, rsp); err != nil {
		log.Fatalf("Error calling hello service: %v", err)
	}

	fmt.Println("Successfully called hello service")
}
