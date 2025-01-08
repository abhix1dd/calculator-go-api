package operation

import (
	"encoding/json"
	"log"
	"time"
	"net/http"
	"os"
	"github.com/golang-jwt/jwt/v4"

)
var logger = initLogger()

// Types
type CalculationRequest struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

var validUsername = "admin"
var validPassword = "password123"

var jwtSecret = []byte("abc@123")


// Handlers
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username != validUsername || user.Password != validPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(user.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println("Processing request...")

	// Parse the JSON request body
	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		logger.Printf("Invalid Input: %v\n", err)
		return
	}

	// Perform the calculation
	result := req.Num1 + req.Num2
	resp := CalculationResponse{Result: result}
	username := r.Context().Value("username").(string)
	logger.Printf("Request processed by user: %s", username)
	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Printf("Request processed successfully: %v + %v = %v\n", req.Num1, req.Num2, result)
}

//Handle for subtract
func SubtractHandler(w http.ResponseWriter, r*http.Request){
	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		logger.Printf("Invalid Input: %v\n", err)
		return
	}
	result := req.Num1 - req.Num2;
	resp := CalculationResponse{Result:result}
	username := r.Context().Value("username").(string)
	logger.Printf("Request processed by user: %s", username)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Printf("Request processed successfully: %v - %v = %v\n", req.Num1, req.Num2, result)
}


//Handle for Product
func ProductHandler(w http.ResponseWriter, r*http.Request){
	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		logger.Printf("Invalid Input: %v\n", err)
		return
	}
	result := req.Num1 * req.Num2;
	resp := CalculationResponse{Result:result}
	username := r.Context().Value("username").(string)
	logger.Printf("Request processed by user: %s", username)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Printf("Request processed successfully: %v * %v = %v\n", req.Num1, req.Num2, result)
}

//Handle for Divison
func DivisonHandler(w http.ResponseWriter, r*http.Request){
	var req CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		logger.Printf("Invalid Input: %v\n", err)
		return
	}
	result := req.Num1 / req.Num2;
	resp := CalculationResponse{Result:result}
	username := r.Context().Value("username").(string)
	logger.Printf("Request processed by user: %s", username)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logger.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Printf("Request processed successfully: %v / %v = %v\n", req.Num1, req.Num2, result)
}

// Generate JWT Token
func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(jwtSecret)
}

// Initialize a logger
func initLogger() *log.Logger {
	return log.New(os.Stdout, "LOG: ", log.LstdFlags|log.Lshortfile)
}
