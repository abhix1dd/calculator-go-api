# Calculator API

A simple yet powerful **REST API** for performing basic arithmetic operations such as addition, subtraction, multiplication, and division. The API also includes JWT-based authentication and rate limiting to ensure secure and efficient usage.

## Features

- **Basic Arithmetic Operations**:
  - Addition (`/add`)
  - Subtraction (`/subtract`)
  - Multiplication (`/product`)
  - Division (`/divide`)
  
- **JWT Authentication** for secure access.
- **Rate Limiting** to control the frequency of requests (2 requests per 5 seconds).
  
## Technologies Used

- **Go** (Golang)
- **JWT** for authentication
- **Rate Limiting** with `golang.org/x/time/rate`
- **CORS** handling for cross-origin requests
- **JSON** for data exchange

## API Endpoints

### 1. **Login**

Authenticate users and obtain a JWT token for further requests.

- **Endpoint**: `/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "admin",
    "password": "password123"
  }
Response:
**{
  "token": "<JWT token>"
}**

2. Addition
Perform addition of two numbers.

Endpoint: /add
Method: POST
Request Body:
**{
  "num1": 10,
  "num2": 5
}**
Response:
**{
  "result": 15
}**

3. Subtraction
Perform subtraction of two numbers.

Endpoint: /subtract
Method: POST
Request Body:
**{
  "num1": 10,
  "num2": 5
}**
Response:
**{
  "result": 5
}**

4. Multiplication
Perform multiplication of two numbers.

Endpoint: /product
Method: POST
Request Body:
**{
  "num1": 10,
  "num2": 5
}**
Response:
**{
  "result": 50
}**

5. Division
Perform division of two numbers.

Endpoint: /divide
Method: POST
Request Body:
**{
  "num1": 10,
  "num2": 5
}**
Response:
**{
  "result": 2
}**

Rate Limiting
The API is rate-limited to 2 requests per 5 seconds. If the limit is exceeded, you will receive a response with the following error:
**{
  "message": "Rate limit exceeded. Try again later."
}**

Installation
To run this project locally, follow the steps below:

1. Clone the repository
bash

git clone https://github.com/your-username/Calculator-API.git
cd Calculator-API

2. Install Dependencies
Ensure you have Go installed. You can check this with the command go version. If it's not installed, download Go.

3. Run the server
bash

go run main.go
The server will start on http://localhost:8080.

4. Test the API
You can test the API using Postman or cURL by sending POST requests to the endpoints mentioned above. Make sure to include the Authorization header with the Bearer token for authenticated routes.

Example Usage with cURL
1. Get the JWT token:
bash

curl -X POST http://localhost:8080/login -d '{"username": "admin", "password": "password123"}' -H "Content-Type: application/json"

2. Use the token to make requests:
Addition:
curl -X POST http://localhost:8080/add -d '{"num1": 10, "num2": 5}' -H "Authorization: Bearer <JWT token>" -H "Content-Type: application/json"

Subtraction:
curl -X POST http://localhost:8080/subtract -d '{"num1": 10, "num2": 5}' -H "Authorization: Bearer <JWT token>" -H "Content-Type: application/json"

Multiplication:
curl -X POST http://localhost:8080/product -d '{"num1": 10, "num2": 5}' -H "Authorization: Bearer <JWT token>" -H "Content-Type: application/json"

Division:
curl -X POST http://localhost:8080/divide -d '{"num1": 10, "num2": 5}' -H "Authorization: Bearer <JWT token>" -H "Content-Type: application/json"
