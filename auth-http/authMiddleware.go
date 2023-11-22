// authMiddleware.go
package main

import (
	"net/http"
	"strings"
)

// AuthMiddleware checks for the presence of an Authorization header and validates its content
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Split the token type and the token itself
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != "valid-token" {
			http.Error(w, "Invalid or missing auth token", http.StatusForbidden)
			return
		}

		// Token is valid, proceed with the request
		next(w, r)
	}
}
