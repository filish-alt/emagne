# Emagne - Golang Gin API

A RESTful API built with Go, Gin framework, and PostgreSQL for user authentication and management.

## Features

- User registration and login
- JWT-based authentication
- Password hashing with bcrypt
- PostgreSQL database with SQLC for type-safe queries
- Database migrations
- CORS support
- Graceful shutdown

## Prerequisites

- Go 1.24.0 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd emagne
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**
   - Create a database named `escrow` (or update the connection string in config)
   - Update the database connection string in `internal/config/config.go` or set the `DATABASE_URL` environment variable

4. **Run database migrations**
   ```bash
   make migrate
   # or
   go run cmd/migrate/main.go
   ```

5. **Start the application**
   ```bash
   make run
   # or
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080` by default.

## Environment Variables

You can configure the application using environment variables:

- `DATABASE_URL`: PostgreSQL connection string (default: `postgresql://postgres:password@localhost:5432/escrow?sslmode=disable`)
- `SERVER_ADDRESS`: Server address (default: `:8080`)
- `JWT_SECRET`: JWT secret key (default: `your-secret-key-change-in-production`)

## API Endpoints

### Authentication

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890"
}
```

#### Login User
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Get Profile (Protected)
```http
GET /api/profile
Authorization: Bearer <jwt_token>
```

### Health Check

#### Health Check
```http
GET /health
```

## Project Structure

```
├── cmd/
│   ├── main.go              # Main application entry point
│   └── migrate/
│       └── main.go          # Database migration runner
├── db/
│   ├── migrations/          # Database migration files
│   └── query/               # SQL queries for SQLC
├── internal/
│   ├── auth/                # Authentication service
│   ├── config/              # Configuration management
│   ├── database/            # Database models and queries
│   ├── handler/             # HTTP handlers
│   └── middleware/          # HTTP middleware
├── pkg/
│   └── utils/               # Utility functions
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Development

### Running Tests
```bash
make test
```

### Formatting Code
```bash
make fmt
```

### Building
```bash
make build
```

## Database Schema

The application uses the following main tables:

- `users`: Stores user information including email, password hash, names, and verification status

## Security Features

- Password hashing using bcrypt
- JWT tokens for authentication
- Input validation
- CORS support
- SQL injection protection through SQLC

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.

