// httpService.go
package main

import (
	"fmt"
	"net/http"

	"go-micro.dev/v4/web"
)

func main() {
	// Create a new web service
	service := web.NewService(
		web.Name("http.service"),
		web.Address(":8080"),
	)

	// Initialize the service
	service.Init()

	// Define the protected handler
	protectedHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Access to protected resource granted!")
	}

	// Wrap the protected handler with the authentication middleware
	protectedEndpoint := AuthMiddleware(protectedHandler)

	// Register the handler with the Go-micro web service
	service.HandleFunc("/protected", protectedEndpoint)

	// Start the service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
