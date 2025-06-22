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

blog-service/\n
├── cmd/ # Main entry point\n
│ └── main.go\n
├── handler/ # HTTP layer (controller)\n
│ └── blog.go\n
├── service/ # Business logic layer\n
│ └── blog.go\n
├── store/ # Data persistence layer\n
│ └── blog.go\n
├── model/ # Domain models\n
│ └── blog.go\n
├── factory/ # Dependency injection / DB init\n
│ └── store_factory.go\n
├── configs/ # .env configuration\n
│ └── .env\n
├── go.mod\n
└── go.sum\n

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
