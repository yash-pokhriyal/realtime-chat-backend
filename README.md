# Realtime Chat Backend (Go)

A production-inspired realtime chat backend built with Go, Gin, PostgreSQL, GORM, JWT Authentication, and Gorilla WebSocket.

## Features

- User Registration
- User Login
- Password Hashing (bcrypt)
- JWT Authentication
- Protected Routes
- PostgreSQL Database
- GORM ORM
- Repository Pattern
- WebSocket Integration
- Real-Time Message Broadcasting
- Concurrent Client Management using Goroutines & Channels

---

## Tech Stack

- Go
- Gin
- PostgreSQL
- GORM
- Gorilla WebSocket
- JWT
- bcrypt

---

## Project Structure

```
realtime-chat-backend/
│
├── cmd/
│   └── server/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── utils/
│   └── websocket/
│
├── migrations/
├── scripts/
├── .env.example
├── go.mod
└── README.md
```

---

## Architecture

```
                Client

                   │

          HTTP / WebSocket

                   │

                Gin Router

        ┌──────────┴──────────┐
        │                     │

 Authentication         WebSocket

        │                     │

     Handlers             Hub Pattern

        │                     │

 Repository Layer      Read / Write Pumps

        │                     │

     PostgreSQL      Concurrent Clients
```

---

## Authentication Flow

```
Register

↓

Hash Password

↓

Store User

↓

Login

↓

Verify Password

↓

Generate JWT

↓

Protected APIs
```

---

## WebSocket Flow

```
Client

↓

GET /ws

↓

Upgrade HTTP → WebSocket

↓

Register Client

↓

Hub

↓

Broadcast

↓

All Connected Clients
```

---

## API Endpoints

### Authentication

| Method | Endpoint |
|---------|----------|
| POST | /register |
| POST | /login |

### Health

| Method | Endpoint |
|---------|----------|
| GET | /health |

### WebSocket

| Method | Endpoint |
|---------|----------|
| GET | /ws |

---

## Environment Variables

```
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=realtime_chat

JWT_SECRET=your-secret-key

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```

---

## Running Locally

### Clone Repository

```bash
git clone <repository-url>
cd realtime-chat-backend
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment

Create a `.env` file.

### Run Server

```bash
go run cmd/server/main.go
```

---

## Example Login Response

```json
{
  "message": "Login successful",
  "token": "<jwt-token>",
  "user": {
    "id": 1,
    "name": "Yash",
    "email": "yash@example.com"
  }
}
```

---

## Implemented Concepts

- REST APIs
- Repository Pattern
- Password Hashing
- JWT Authentication
- Authentication Middleware
- Environment Configuration
- PostgreSQL
- GORM
- WebSockets
- Goroutines
- Channels
- Hub Pattern
- Concurrent Client Management

---

## Upcoming Features

- JWT Protected WebSocket
- Chat Rooms
- Private Messaging
- Redis Pub/Sub
- Docker
- Docker Compose
- Graceful Shutdown
- Unit Testing
- Swagger Documentation
- Deployment

---

## Learning Outcomes

This project demonstrates practical backend development concepts including:

- Backend API Development
- Authentication & Authorization
- Concurrent Programming in Go
- Real-Time Communication
- Database Integration
- Clean Project Structure
- Repository Pattern
- Production-inspired Backend Design