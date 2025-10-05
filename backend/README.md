# learng Backend

Go backend for the learng language learning platform.

## Technology Stack

- **Go**: 1.21+
- **Framework**: Echo v4
- **ORM**: GORM
- **Database**: SQLite (MVP) → PostgreSQL (production)
- **Authentication**: JWT

## Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/               # Configuration management
│   ├── models/               # GORM models
│   ├── handlers/             # HTTP handlers (controllers)
│   ├── middleware/           # Custom middleware
│   ├── services/             # Business logic
│   ├── repository/           # Database access layer
│   └── utils/                # Utility functions
├── migrations/               # SQL migration files (future)
├── uploads/                  # Local file storage (dev)
│   ├── images/
│   └── audio/
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository-url>
   cd learng/backend
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Setup environment variables**:
   ```bash
   cp .env.example .env
   # Edit .env and set JWT_SECRET
   ```

4. **Run the application**:
   ```bash
   go run cmd/api/main.go
   ```

   Or use the Makefile:
   ```bash
   make run
   ```

5. **Access the API**:
   ```
   http://localhost:8080
   Health check: http://localhost:8080/health
   ```

## Development

### Available Commands

```bash
# Run the application
make run

# Run with hot reload (using air)
make dev

# Run tests
make test

# Build binary
make build

# Clean build artifacts
make clean

# Format code
make fmt

# Run linter
make lint

# Initialize database
make db-init
```

### Environment Variables

See `.env.example` for all available configuration options:

- `PORT`: Server port (default: 8080)
- `DB_PATH`: SQLite database file path
- `JWT_SECRET`: Secret key for JWT tokens (REQUIRED)
- `UPLOAD_DIR`: Directory for uploaded files
- `STATIC_DIR`: Frontend build directory (production only)
- `MAX_IMAGE_SIZE`: Maximum image upload size in bytes
- `MAX_AUDIO_SIZE`: Maximum audio upload size in bytes

### Database Migrations

The application uses GORM's AutoMigrate feature for development:

```go
db.AutoMigrate(
    &models.User{},
    &models.Journey{},
    // ... other models
)
```

For production, proper migration files will be added.

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user

### Journeys (Admin)

- `GET /api/v1/journeys` - List journeys
- `POST /api/v1/journeys` - Create journey
- `GET /api/v1/journeys/:id` - Get journey details
- `PUT /api/v1/journeys/:id` - Update journey
- `DELETE /api/v1/journeys/:id` - Delete journey

### Scenarios (Admin)

- `POST /api/v1/scenarios` - Create scenario
- `GET /api/v1/scenarios/:id` - Get scenario details
- `PUT /api/v1/scenarios/:id` - Update scenario
- `DELETE /api/v1/scenarios/:id` - Delete scenario

### Words (Admin)

- `POST /api/v1/words` - Create word
- `GET /api/v1/words/:id` - Get word details
- `PUT /api/v1/words/:id` - Update word
- `DELETE /api/v1/words/:id` - Delete word

### Media Upload (Admin)

- `POST /api/v1/media/upload/image` - Upload image
- `POST /api/v1/media/upload/audio` - Upload audio

### Learner Endpoints

- `GET /api/v1/learner/journeys` - List available journeys
- `GET /api/v1/learner/scenarios/:id` - Get scenario with words
- `POST /api/v1/learner/progress` - Track word view

### Quizzes

- `GET /api/v1/quizzes/:id` - Get quiz questions
- `POST /api/v1/quizzes/:id/submit` - Submit quiz answers

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/handlers
```

## Building for Production

### Build Binary

```bash
make build
# or
go build -o learng-api ./cmd/api
```

### Docker Build

See the root Dockerfile for multi-stage build that includes frontend.

For backend-only testing:

```bash
docker build -t learng-backend:latest -f Dockerfile.backend .
docker run -p 8080:8080 learng-backend:latest
```

## Project Status

### Completed ✅

- [x] Project structure setup
- [x] GORM models defined
- [x] Configuration management
- [x] JWT authentication utilities
- [x] Middleware (auth, CORS, logging)
- [x] Database initialization

### In Progress 🚧

- [ ] Handler implementation (auth, journeys, etc.)
- [ ] Repository layer
- [ ] Service layer
- [ ] File upload handling
- [ ] Comprehensive tests

### Planned 📋

- [ ] AI integration (Phase 2)
- [ ] Advanced analytics
- [ ] Rate limiting
- [ ] API documentation (Swagger)

## Contributing

1. Create a feature branch
2. Make your changes
3. Run tests and linting
4. Submit a pull request

## License

[To be determined]

## Support

For questions or issues, please contact the development team.
