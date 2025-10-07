# Learng Backend

AI-powered language learning platform backend built with Go and Echo framework.

## 🎯 Current Status

✅ **Phase 1 - Sprint 1 Complete**
- Authentication system fully implemented and tested
- User registration and login working
- JWT-based authentication
- Protected endpoints with middleware
- Input validation and error handling

## Tech Stack

- **Go**: 1.21+
- **Web Framework**: Echo v4.11.3
- **ORM**: GORM v1.25.5
- **Database**: SQLite 3.x (MVP) / PostgreSQL (Production)
- **Authentication**: JWT (golang-jwt/jwt v5.1.0)
- **Password Hashing**: bcrypt

## Quick Start

### Prerequisites

- Go 1.21 or higher
- SQLite3

### Installation & Running

```bash
# Install dependencies
go mod download

# Build the application
make build

# Run the server
make run
# or
./bin/api
```

Server will start on http://localhost:8080

### Test Authentication

```bash
# Run automated test suite
./test-auth.sh
```

Expected output: ✅ All 10 tests pass

## 📁 Project Structure

```
backend/
├── cmd/api/main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration loading
│   ├── handlers/
│   │   └── auth.go             # ✅ Authentication endpoints
│   ├── middleware/
│   │   └── auth.go             # JWT authentication
│   ├── models/                  # Database models (GORM)
│   │   ├── user.go
│   │   ├── journey.go
│   │   ├── scenario.go
│   │   └── ...
│   ├── repository/
│   │   └── user.repo.go        # ✅ User data access
│   ├── services/
│   │   └── auth.service.go     # ✅ Auth business logic
│   └── utils/                   # Utilities (JWT, validation, etc.)
├── docs/                        # API documentation
│   └── api.md
├── .env                         # Environment variables
├── test-auth.sh                # ✅ Authentication tests
├── Makefile                     # Build commands
└── learng.db                   # SQLite database (auto-created)
```

## 🔐 API Endpoints

### Public Endpoints
- `GET /health` - Health check
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user

### Protected Endpoints (Requires JWT)
- `GET /api/v1/auth/me` - Get current user

### Coming Soon
- Journey CRUD endpoints
- Scenario management
- Word management
- Media upload
- Quiz system

## 🧪 Testing

### Automated Tests
```bash
./test-auth.sh
```

Tests include:
- ✅ User registration (admin/learner)
- ✅ Login flow
- ✅ Protected endpoint access
- ✅ Duplicate email prevention
- ✅ Input validation (email, password)
- ✅ Error handling

### Manual Testing
```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@test.com","password":"password123","displayName":"Test User","role":"admin"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@test.com","password":"password123"}'

# Get current user (use token from login)
curl -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/api/v1/auth/me
```

## 📚 Documentation

- [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md) - Quick start and debugging
- [AUTH_IMPLEMENTATION.md](AUTH_IMPLEMENTATION.md) - Authentication details
- [TEST_RESULTS.md](TEST_RESULTS.md) - Latest test results
- [docs/api.md](docs/api.md) - API documentation

## 🏗️ Architecture

**Three-Layer Architecture**:
1. **Handler Layer** - HTTP request/response handling
2. **Service Layer** - Business logic and validation
3. **Repository Layer** - Database operations

**Key Features**:
- Dependency injection in `main.go`
- Middleware for authentication and CORS
- Structured error responses
- Input validation utilities
- JWT token management

## 🔧 Development

### Make Commands
```bash
make build    # Build the application
make run      # Run the application
make test     # Run tests
make clean    # Remove build artifacts
```

### Environment Variables

Create `.env` file:
```env
APP_ENV=development
APP_PORT=8080
DB_DRIVER=sqlite
DB_PATH=./learng.db
JWT_SECRET=your-secret-key-here
UPLOAD_DIR=./uploads
MAX_FILE_SIZE=10485760
ALLOWED_ORIGINS=http://localhost:5173
```

### Database

View database contents:
```bash
sqlite3 learng.db
> .tables
> SELECT * FROM users;
> .quit
```

## 🚀 Next Steps

### Sprint 2 (Weeks 3-4)
- [ ] Journey CRUD endpoints
- [ ] Scenario management
- [ ] Word management
- [ ] Admin-only routes with role middleware

### Sprint 3 (Weeks 5-6)
- [ ] Media upload handling
- [ ] AI service integration stubs
- [ ] Quiz system

### Future
- [ ] PostgreSQL migration
- [ ] Docker deployment
- [ ] API rate limiting
- [ ] Refresh token mechanism

## 📝 Contributing

1. Follow the three-layer architecture pattern
2. Add tests for new endpoints
3. Update documentation
4. Use the existing error handling patterns

## 🐛 Troubleshooting

**Port already in use**:
```bash
lsof -ti:8080 | xargs kill -9
```

**Database locked**:
```bash
# Kill all instances and restart
pkill api
./bin/api
```

**Missing JWT_SECRET**:
```bash
# Create .env file with JWT_SECRET
echo "JWT_SECRET=dev-secret-change-in-production" >> .env
```

---

**Last Updated**: 2025-10-07  
**Current Version**: v0.1.0-alpha  
**Status**: ✅ Authentication Complete | 🚧 Journey Management In Progress
