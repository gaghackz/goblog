# 📝 Blog Service — Built with GoFr (Golang Framework)

A minimal, clean, and production-ready RESTful blog service built using [GoFr](https://github.com/tothenew/gofr) — an opinionated Golang framework for building scalable backend applications.

This project demonstrates:

- Three-layer clean architecture
- SOLID design principles
- Interface-driven development
- MySQL integration via Docker
- Factory pattern for dependency management

---

## 🚀 Features

- 🆕 Create blog post (`POST /blogs`)
- 📖 Retrieve all blog posts (`GET /blogs`)
- 🔍 Retrieve a blog post by ID (`GET /blogs/{id}`)
- 🧱 Modular architecture for easy testing and extension
- 🐳 Docker-based MySQL integration

---

## 🧱 Project Structure

blog-service/
├── cmd/ # Main entry point
│ └── main.go
├── handler/ # HTTP layer (controller)
│ └── blog.go
├── service/ # Business logic layer
│ └── blog.go
├── store/ # Data persistence layer
│ └── blog.go
├── model/ # Domain models
│ └── blog.go
├── factory/ # Dependency injection / DB init
│ └── store_factory.go
├── configs/ # .env configuration
│ └── .env
├── go.mod
└── go.sum

---

## 🌐 API Endpoints

| Method | Endpoint      | Description            |
| ------ | ------------- | ---------------------- |
| POST   | `/blogs`      | Create a new blog post |
| GET    | `/blogs`      | Get all blog posts     |
| GET    | `/blogs/{id}` | Get blog by ID         |

---

## 🔁 Architecture Explained

This service follows a **Three-Layer Architecture** pattern:

### 1️⃣ Handler Layer (Presentation)

- Accepts incoming HTTP requests
- Parses path/body parameters
- Calls service layer
- Example: `handler/blog.go`

### 2️⃣ Service Layer (Application)

- Contains business logic
- Implements the `BlogService` interface
- Delegates data access to the store layer
- Example: `service/blog.go`

### 3️⃣ Store Layer (Data)

- Talks directly to MySQL using SQL queries
- Implements the `BlogStore` interface
- Returns domain models to the service layer
- Example: `store/blog.go`

✅ This separation keeps the application loosely coupled, testable, and extensible.

---

## 🧠 SOLID Principles Applied

| Principle                     | How it’s Applied                                                        |
| ----------------------------- | ----------------------------------------------------------------------- |
| **S** - Single Responsibility | Handler, Service, and Store have distinct, single responsibilities      |
| **O** - Open/Closed           | New features can be added without altering existing code                |
| **L** - Liskov Substitution   | Concrete types can replace interfaces seamlessly                        |
| **I** - Interface Segregation | Each layer depends only on interfaces it uses                           |
| **D** - Dependency Inversion  | High-level modules (Handler) depend on abstractions, not concrete Store |
