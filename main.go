package main
import "calculator-api/operation"

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "time"
	"context"
	// "CALCULATOR-API/operation"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/cors"
	"golang.org/x/time/rate"
)

var jwtSecret = []byte("abc@123")

var validUsername = "admin"
var validPassword = "password123"

var limiter = rate.NewLimiter(2, 5)
var logger = initLogger()

// Middleware: Rate Limiting
func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded. Try again later.", http.StatusTooManyRequests)
			fmt.Println("Failed limit exceed")
			logger.Printf("Rate limit exceeded. Try again later.")
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware: JWT Authentication
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		var tokenString string
		fmt.Sscanf(authHeader, "Bearer %s", &tokenString)
		if tokenString == "" {
			http.Error(w, "Bearer token missing", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)
			ctx := r.Context()
			ctx = context.WithValue(ctx, "username", username)
			r = r.WithContext(ctx)
		}
		fmt.Println("successs")

		next.ServeHTTP(w, r)
	})
}

func main() {
	logger.Println("Starting server...")

	// Handlers
	mux := http.NewServeMux()
	mux.Handle("/login", http.HandlerFunc(operation.LoginHandler))
	mux.Handle("/add", authMiddleware(rateLimitMiddleware(http.HandlerFunc(operation.AddHandler))))
	mux.Handle("/subtract", authMiddleware(rateLimitMiddleware(http.HandlerFunc(operation.SubtractHandler))))
	mux.Handle("/product", authMiddleware(rateLimitMiddleware(http.HandlerFunc(operation.ProductHandler))))
	mux.Handle("/divide", authMiddleware(rateLimitMiddleware(http.HandlerFunc(operation.DivisonHandler))))

	// Rate Limiter and CORS
	rateLimitedMux := rateLimitMiddleware(mux)
	handler := cors.Default().Handler(rateLimitedMux)

	logger.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

// Initialize a logger
func initLogger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.LstdFlags|log.Lshortfile)
}
