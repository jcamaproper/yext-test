# Sorting Service

This is a RESTful API service built in **Golang (1.22)** using **Gin Gonic** to handle sorting of arrays within a JSON payload. The service supports **parallel processing** for large arrays and **efficient error handling**.

## Features
- ✅ Sorts arrays of strings and numbers within a JSON payload.
- ✅ Parallel sorting for large arrays using Goroutines and `errgroup`.
- ✅ Validates input payload and handles errors gracefully.
- ✅ Unit tests covering various scenarios and edge cases.
- ✅ Lightweight and fast, designed with simplicity and performance in mind.

---

## Project Structure
```
myapp/
│
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
├── handler/
│   └── sort_handler.go
├── service/
│   └── service_sort.go
|   └── service_sort_test.go
├── routes/
│   └── router.go
├── model/
│   └── request.go
├── error/
    └── custom_error.go


```
---

## Installation & Setup
### 1. Clone the Repository
```bash
git clone https://github.com/jcamalpz/yext-test
cd app
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Run the Service
```bash
go run main.go
```
The server will start at `http://localhost:8080`

---

## Usage
### Sample Request
```bash
curl -X POST http://localhost:8080/sort \
-H "Content-Type: application/json" \
-d '{
  "sortKeys": ["fruits", "numbers"],
  "payload": {
    "fruits": ["watermelon", "apple", "pineapple"],
    "numbers": [1333, 4, 2431, 7],
    "colors": ["green", "blue", "yellow"]
  }
}'
```
### Expected Response
```json
{
  "fruits": ["apple", "pineapple", "watermelon"],
  "numbers": [4, 7, 1333, 2431],
  "colors": ["green", "blue", "yellow"]
}
```

---

## Running Tests
Unit tests are provided to cover various cases, including parallel processing and edge cases.
```bash
go test ./...
```

---

## Docker Setup
### 1. Build Docker Image
```bash
docker build -t sorting-service .
```

### 2. Run Container
```bash
docker run -p 8080:8080 sorting-service
```

---

## Key Points to Remember
- **Parallel Processing:** Parallelism is applied at the key-sorting level, leveraging Goroutines and `errgroup`.
- **Built-in Go Sorting:** Internal array sorting is handled using Go’s optimized `sort` package.
- **Error Handling:** Custom errors are used to validate the payload and handle edge cases.
- **Deep Copy:** Ensures that the original payload is not modified during sorting.

---

## Assumptions & Limitations
- The service handles **string, float64, and interger types** only within arrays. Other types will trigger an error.
- Sorting is case-sensitive for strings.
- Arrays with mixed data types are not supported.

---

## Example Edge Cases Covered in Tests
- Empty `payload`
- Empty `sortKeys`
- Unsupported data types
- Large arrays for parallel processing (numbers and strings)

---

## Author
[Juan Camacho]
[jcamacholpz@gmail.com]  
[https://github.com/jcamalpz]

---

## License
This project is licensed under the MIT License.

