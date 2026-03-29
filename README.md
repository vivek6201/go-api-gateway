# 🚀 Go API Gateway (Mini AWS API Gateway)

A lightweight API Gateway built in Go to understand core backend and infrastructure concepts like reverse proxying, middleware, authentication, and rate limiting.

This project is inspired by real-world systems like AWS API Gateway, but focuses on learning by building a simplified, production-style version from scratch.

---

## 🧠 What This Project Does

This gateway acts as a **single entry point** for multiple backend services.

Instead of clients calling services directly:

```
Client → Users Service
Client → Orders Service
```

You route everything through the gateway:

```
Client → API Gateway → Backend Services
```

---

## ⚙️ Features (MVP)

### ✅ Implemented

- Reverse Proxy (request forwarding)
- Static Routing
- Logging Middleware
- API Key Authentication
- Basic Rate Limiting
- Graceful Error Handling (502 when service is down)

### ❌ Not Included Yet (Planned)

- Dynamic routing (Admin API)
- Load balancing
- Circuit breaker
- WebSocket support
- Persistent storage (DB)

---

## 🏗️ Project Structure

```
go-api-gateway/
│
├── cmd/
│   └── gateway/
│       └── main.go          # Entry point
│
├── internal/
│   ├── router/              # Routing logic
│   ├── proxy/               # Reverse proxy logic
│   ├── middleware/          # Logging, Auth, Rate limiting
│   └── config/              # Route configuration
│
└── go.mod
```

---

## 🔁 How It Works

### Request Flow

```
Client Request
      ↓
Gateway Server
      ↓
Middleware Chain
 (Logging → Auth → RateLimit)
      ↓
Router
      ↓
Reverse Proxy
      ↓
Backend Service
      ↓
Response → Client
```

---


## 🧪 Running the Project

### 1. Start Dummy Services

**Users Service (port 9001)**

```go
http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Users Service"))
})
http.ListenAndServe(":9001", nil)
```

**Orders Service (port 9002)**

```go
http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Orders Service"))
})
http.ListenAndServe(":9002", nil)
```

### 2. Run Gateway

```bash
go run cmd/gateway/main.go
```

### 3. Test Endpoints

```bash
curl localhost:8080/users
curl localhost:8080/orders
```

## 🧠 Key Learnings

This project helps you understand:

- Go HTTP internals
- Reverse proxy architecture
- Middleware design pattern
- Request/response lifecycle
- Basic system design
- Fault tolerance basics

---

## 🚀 Future Improvements

- Dynamic route management (Admin API)
- Redis-based rate limiting
- JWT authentication
- Load balancing (round robin)
- Circuit breaker
- Metrics and monitoring
- WebSocket support for real-time features

---

## 💡 Why This Project Matters

This is not just a CRUD app. It demonstrates:

- Backend engineering skills
- System design understanding
- Real-world infrastructure concepts

---

## 👨‍💻 Author

Built as a learning project to deeply understand Go and backend systems.

---

## ⭐ Final Note

This project is a **learning-focused implementation** of an API Gateway. It is intentionally simple, but designed in a way that can be extended into a production-grade system.