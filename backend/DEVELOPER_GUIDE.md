# Developer Quick Start Guide

## 🚀 Running the Backend

### First Time Setup

1. **Install Dependencies**:
```bash
cd backend
go mod download
```

2. **Create Environment File**:
```bash
cp .env.example .env
# Edit .env and set JWT_SECRET to a secure value
```

3. **Build the Application**:
```bash
make build
# or
go build -o bin/api ./cmd/api
```

4. **Run the Server**:
```bash
make run
# or
./bin/api
```

The server will start on http://localhost:8080

### Environment Variables

Required variables in `.env`:
```env
APP_ENV=development
APP_PORT=8080
DB_DRIVER=sqlite
DB_PATH=./learng.db
JWT_SECRET=your-secret-key-change-in-production
UPLOAD_DIR=./uploads
MAX_FILE_SIZE=10485760
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

## 🧪 Testing

### Run Authentication Tests
```bash
./test-auth.sh
```

This will test:
- User registration (admin and learner)
- Login flow
- Protected endpoints
- Input validation
- Error handling

### Manual API Testing

**Register a new user**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "displayName": "John Doe",
    "role": "admin"
  }'
```

**Login**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**Get current user (protected)**:
```bash
TOKEN="your-token-from-login"
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8080/api/v1/auth/me
```

## 📁 Project Structure

```
backend/
├── cmd/api/main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration loading
│   ├── handlers/                # HTTP request handlers
│   │   └── auth.go             # ✅ Authentication endpoints
│   ├── middleware/              # HTTP middleware
│   │   └── auth.go             # JWT authentication
│   ├── models/                  # Database models (GORM)
│   ├── repository/              # Data access layer
│   │   └── user.repo.go        # ✅ User database operations
│   ├── services/                # Business logic layer
│   │   └── auth.service.go     # ✅ Auth business logic
│   └── utils/                   # Utility functions
│       ├── jwt.go              # JWT token generation/validation
│       ├── password.go         # Password hashing
│       ├── response.go         # HTTP response helpers
│       └── validation.go       # Input validation
├── .env                         # Environment variables
├── go.mod                       # Go dependencies
├── Makefile                     # Build commands
├── test-auth.sh                # ✅ Authentication test suite
└── learng.db                   # SQLite database (auto-created)
```

## 🔐 API Endpoints

### Public Endpoints (No Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/v1/auth/register` | Register new user |
| POST | `/api/v1/auth/login` | Login user |

### Protected Endpoints (Requires JWT Token)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/auth/me` | Get current user info |

## 🏗️ Architecture Layers

### 1. Handler Layer (`internal/handlers/`)
- Handles HTTP requests/responses
- Validates request format
- Calls service layer
- Returns JSON responses

### 2. Service Layer (`internal/services/`)
- Contains business logic
- Validates business rules
- Calls repository layer
- Handles errors

### 3. Repository Layer (`internal/repository/`)
- Database operations (GORM)
- CRUD operations
- Query building
- No business logic

## 🛠️ Make Commands

```bash
make build    # Build the application
make run      # Run the application
make test     # Run tests
make clean    # Remove build artifacts
make watch    # Run with hot reload (requires air)
```

## 📊 Database

### View Database Contents

```bash
sqlite3 learng.db
```

Useful queries:
```sql
-- List all tables
.tables

-- View users
SELECT id, email, role, display_name FROM users;

-- View schema
.schema users

-- Exit
.quit
```

### Reset Database

```bash
rm learng.db
# Restart server - it will auto-migrate and create new database
```

## 🔧 Debugging

### View Server Logs

If running with nohup:
```bash
tail -f server.log
```

### Common Issues

**"JWT_SECRET environment variable is required"**
- Solution: Create `.env` file with JWT_SECRET

**"Port 8080 already in use"**
- Solution: Kill existing process or change APP_PORT in .env
```bash
lsof -ti:8080 | xargs kill -9
```

**Database locked errors**
- Solution: Only one process can write to SQLite at a time
- Make sure no other instances are running

## 📝 Adding New Endpoints

1. **Create Handler** in `internal/handlers/`
2. **Create Service** in `internal/services/`
3. **Create Repository** in `internal/repository/` (if needed)
4. **Wire up in** `cmd/api/main.go`
5. **Add tests** in `test-*.sh` or Go tests

Example pattern:
```go
// 1. Repository (internal/repository/example.repo.go)
type ExampleRepository struct {
    db *gorm.DB
}

// 2. Service (internal/services/example.service.go)
type ExampleService struct {
    exampleRepo *repository.ExampleRepository
}

// 3. Handler (internal/handlers/example.go)
type ExampleHandler struct {
    exampleService *services.ExampleService
}

// 4. Wire up (cmd/api/main.go)
exampleRepo := repository.NewExampleRepository(db)
exampleService := services.NewExampleService(exampleRepo)
exampleHandler := handlers.NewExampleHandler(exampleService)

v1.GET("/examples", exampleHandler.List)
v1.POST("/examples", exampleHandler.Create, customMiddleware.AuthMiddleware)
```

## 🔐 Security Checklist

- [x] Passwords hashed with bcrypt
- [x] JWT tokens expire after 24 hours
- [x] Input validation on all endpoints
- [x] CORS configured for allowed origins
- [x] SQL injection prevention (GORM prepared statements)
- [x] Password hash never returned in responses
- [ ] Rate limiting (TODO)
- [ ] Refresh token mechanism (TODO)

## 🚀 Deployment

### Build Production Binary

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/api-linux ./cmd/api
```

### Docker (Coming Soon)

```bash
docker build -t learng-api .
docker run -p 8080:8080 --env-file .env learng-api
```

## 📚 Additional Documentation

- [REQUIREMENT.md](../REQUIREMENT.md) - Business requirements
- [CORE.md](../design/CORE.md) - Functional specification
- [AUTH_IMPLEMENTATION.md](AUTH_IMPLEMENTATION.md) - Authentication details
- [TEST_RESULTS.md](TEST_RESULTS.md) - Latest test results

---

**Last Updated**: 2025-10-07  
**Status**: ✅ Authentication Complete  
**Next**: Journey Management Endpoints
