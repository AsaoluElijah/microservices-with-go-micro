package main

import (
	"context"
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("serviceB"))
	service.Init()

	// Request message
	req := service.Client().NewRequest("serviceA", "Greeter.Hello", &map[string]interface{}{}, client.WithContentType("application/json"))
	rsp := &map[string]interface{}{}

	// Call the service
	if err := service.Client().Call(context.Background(), req, rsp); err != nil {
		log.Fatalf("Error calling service A: %v", err)
	}

	fmt.Println("Successfully called service A")
}
