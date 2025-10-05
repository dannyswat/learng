# Backend Quick Reference

## Directory Structure
```
backend/
├── cmd/api/main.go              # Server entry point
├── internal/
│   ├── config/                  # Configuration
│   ├── models/                  # GORM models (7 files)
│   ├── handlers/                # HTTP handlers (TO DO)
│   ├── middleware/              # Auth middleware
│   ├── services/                # Business logic (TO DO)
│   ├── repository/              # Data access (TO DO)
│   └── utils/                   # Helpers
└── uploads/                     # Media storage
```

## Setup (First Time)
```bash
cd backend
./setup.sh          # Auto-setup everything
# or manually:
cp .env.example .env && go mod download
```

## Daily Development
```bash
make dev            # Start with hot reload
make run            # Start normally
make test           # Run tests
make fmt            # Format code
```

## Database Models

| Model | Purpose | Key Fields |
|-------|---------|-----------|
| User | Authentication | email, role, passwordHash |
| Journey | Content container | title, status, sourceLanguage |
| Scenario | Themed word group | journeyId, title, displayOrder |
| Word | Vocabulary item | targetText, imageUrl, audioUrl |
| Quiz | Assessment | scenarioId, passThreshold |
| QuizQuestion | Quiz item | wordId, questionType, options |
| LearnerProgress | Tracking | userId, wordId, masteryLevel |
| QuizAttempt | Results | userId, quizId, score, answers |

## API Structure (Planned)

### Auth
- POST /api/v1/auth/register
- POST /api/v1/auth/login

### Admin
- GET/POST/PUT/DELETE /api/v1/journeys
- GET/POST/PUT/DELETE /api/v1/scenarios
- GET/POST/PUT/DELETE /api/v1/words
- POST /api/v1/media/upload/image
- POST /api/v1/media/upload/audio

### Learner
- GET /api/v1/learner/journeys
- GET /api/v1/learner/scenarios/:id
- POST /api/v1/learner/progress

### Quiz
- GET /api/v1/quizzes/:id
- POST /api/v1/quizzes/:id/submit

## Environment Variables
```bash
PORT=8080                    # Server port
DB_PATH=./learng.db         # SQLite file
JWT_SECRET=xxx              # REQUIRED
UPLOAD_DIR=./uploads
STATIC_DIR=                 # Empty in dev
MAX_IMAGE_SIZE=5242880      # 5MB
MAX_AUDIO_SIZE=2097152      # 2MB
```

## Common Code Patterns

### Creating a Handler
```go
// internal/handlers/example.go
package handlers

import (
    "github.com/labstack/echo/v4"
    "github.com/learng/backend/internal/middleware"
)

type ExampleHandler struct {
    db *gorm.DB
}

func (h *ExampleHandler) GetExample(c echo.Context) error {
    userID := middleware.GetUserID(c)
    // ... logic
    return c.JSON(200, result)
}
```

### Creating a Repository
```go
// internal/repository/example.repo.go
package repository

import "gorm.io/gorm"

type ExampleRepository struct {
    db *gorm.DB
}

func (r *ExampleRepository) Find(id string) (*Model, error) {
    var model Model
    err := r.db.First(&model, "id = ?", id).Error
    return &model, err
}
```

### Creating a Service
```go
// internal/services/example.service.go
package services

type ExampleService struct {
    repo *repository.ExampleRepository
}

func (s *ExampleService) DoSomething(data Input) (*Output, error) {
    // Business logic here
    return result, nil
}
```

## Utilities

### JWT
```go
utils.GenerateToken(userID, email, role, secret)
utils.ValidateToken(tokenString, secret)
```

### Password
```go
utils.HashPassword(password)
utils.CheckPassword(password, hash)
```

### Validation
```go
utils.ValidateEmail(email)
utils.ValidatePassword(password)
utils.ValidateLanguageCode(code)
utils.ValidateRole(role)
```

### Response
```go
utils.SuccessResponse(c, 200, data)
utils.ErrorResponse(c, 400, "error message")
utils.ValidationErrorResponse(c, errors)
```

## Middleware Usage

### Protect Route
```go
api.POST("/admin/journeys", 
    handler.CreateJourney,
    middleware.AuthMiddleware(cfg.JWTSecret),
    middleware.RequireRole("admin"))
```

### Get User Context
```go
func (h *Handler) MyHandler(c echo.Context) error {
    userID := middleware.GetUserID(c)
    role := middleware.GetUserRole(c)
    // ... use userID and role
}
```

## Testing

### Run Tests
```bash
go test ./...                    # All tests
go test ./internal/handlers      # Specific package
go test -cover ./...             # With coverage
```

### Write Tests
```go
func TestExample(t *testing.T) {
    // Setup
    db := setupTestDB()
    handler := NewHandler(db)
    
    // Test
    result, err := handler.DoSomething()
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expected, result)
}
```

## Debugging

### Enable Debug Mode
```go
e.Debug = true  // in main.go
```

### Log Database Queries
```go
db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
```

## File Upload Example
```go
func (h *Handler) UploadImage(c echo.Context) error {
    file, err := c.FormFile("file")
    // Validate size
    if file.Size > h.cfg.MaxImageSize {
        return utils.ErrorResponse(c, 400, "File too large")
    }
    // Save file
    src, _ := file.Open()
    defer src.Close()
    dst, _ := os.Create(savePath)
    defer dst.Close()
    io.Copy(dst, src)
    
    return utils.SuccessResponse(c, 201, map[string]string{
        "url": "/uploads/images/" + filename,
    })
}
```

## Status Codes
- 200 OK - Success
- 201 Created - Resource created
- 400 Bad Request - Validation error
- 401 Unauthorized - Not authenticated
- 403 Forbidden - Not authorized
- 404 Not Found - Resource not found
- 500 Internal Server Error - Server error

## Helpful Commands
```bash
# View logs
make run | tee server.log

# Test a single endpoint
curl http://localhost:8080/health

# Check database
sqlite3 learng.db "SELECT * FROM users;"

# Reset database
rm learng.db && make run

# Build for production
CGO_ENABLED=1 GOOS=linux go build -o learng-api ./cmd/api
```

---
**Tip**: Keep this file open while developing!
