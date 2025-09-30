# Practice 2 – Tiny HTTP Service in Go

This project is a simple HTTP service built with Go’s `net/http` package.
It demonstrates routes, JSON parsing, validation, and middleware for authentication.

---

## Features

* **GET /user**

  * Query parameter: `id` (integer)
  * Success: `200 OK {"user_id": <id>}`
  * Error: `400 Bad Request {"error":"invalid id"}`

* **POST /user**

  * JSON body: `{"name": "Alice"}`
  * Success: `201 Created {"created":"Alice"}`
  * Error: `400 Bad Request {"error":"invalid name"}`

* **Middleware**

  * Requires header: `X-API-Key: secret123`
  * Denies missing or wrong key with `401 Unauthorized {"error":"unauthorized"}`
  * Logs every request method + path (e.g., `GET /user`)

---

## Project Structure

```
practice2/
  go.mod
  main.go
```

---

## How to Run

1. Clone the repository:

   ```bash
   git clone https://github.com/nurgaliev-d/golang/tree/practice2
   cd practice2
   ```

2. Initialize module (if not already done):

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

4. The server will start on:

   ```
   http://localhost:8080
   ```

---

## How to Test

Use `curl` commands to test endpoints.

### 1. Valid GET

```bash
curl -i -H "X-API-Key: secret123" "http://localhost:8080/user?id=42"
# -> 200 {"user_id":42}
```

### 2. GET without API key

```bash
curl -i "http://localhost:8080/user?id=42"
# -> 401 {"error":"unauthorized"}
```

### 3. GET with invalid id

```bash
curl -i -H "X-API-Key: secret123" "http://localhost:8080/user?id=abc"
# -> 400 {"error":"invalid id"}
```

### 4. Valid POST

```bash
curl -i -X POST -H "X-API-Key: secret123" -H "Content-Type: application/json" \
     -d '{"name":"Alice"}' "http://localhost:8080/user"
# -> 201 {"created":"Alice"}
```

### 5. POST with invalid name

```bash
curl -i -X POST -H "X-API-Key: secret123" -H "Content-Type: application/json" \
     -d '{"name":""}' "http://localhost:8080/user"
# -> 400 {"error":"invalid name"}
```

---

## Logs

Every request is logged to the console, for example:

```
GET /user
POST /user
```

---

## Definition of Done

* Server starts with `go run main.go`
* Both routes work with correct status codes and JSON
* Middleware checks API key and logs requests
* All success and error cases tested with `curl`

---

## License

This project is for educational purposes only.
Feel free to fork and modify for your own learning.
