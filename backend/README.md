# learng Backend# Learng Backend



Go backend API server for the learng language learning platform.AI-powered language learning platform backend built with Go and Echo framework.



## Tech Stack## 🎯 Current Status



- **Language**: Go 1.21+✅ **Phase 1 - Sprint 1 Complete**

- **Framework**: Echo v4 (lightweight web framework)- Authentication system fully implemented and tested

- **ORM**: GORM (database operations)- User registration and login working

- **Database**: SQLite (development) → PostgreSQL (production)- JWT-based authentication

- **Auth**: JWT-based authentication- Protected endpoints with middleware

- **Testing**: Shell scripts + Go tests- Input validation and error handling



## Quick Start## Tech Stack



### Prerequisites- **Go**: 1.21+

- Go 1.21 or higher- **Web Framework**: Echo v4.11.3

- SQLite3- **ORM**: GORM v1.25.5

- **Database**: SQLite 3.x (MVP) / PostgreSQL (Production)

### Setup- **Authentication**: JWT (golang-jwt/jwt v5.1.0)

- **Password Hashing**: bcrypt

```bash

# Clone and navigate## Quick Start

cd backend

### Prerequisites

# Install dependencies

go mod download- Go 1.21 or higher

- SQLite3

# Copy environment template

cp .env.example .env### Installation & Running



# Edit .env with your settings```bash

# Required: JWT_SECRET (generate a random string)# Install dependencies

go mod download

# Build

go build -o bin/api ./cmd/api# Build the application

make build

# Run

bin/api# Run the server

# Server starts on http://localhost:8080make run

```# or

./bin/api

### Development```



```bashServer will start on http://localhost:8080

# Run with hot reload (if using air)

air### Test Authentication



# Run tests```bash

./test-sprint2.sh# Run automated test suite

./test-auth.sh

# Build```

make build

Expected output: ✅ All 10 tests pass

# Clean build artifacts

make clean## 📁 Project Structure

```

```

## Project Structurebackend/

├── cmd/api/main.go              # Application entry point

```├── internal/

backend/│   ├── config/                  # Configuration loading

├── cmd/api/main.go           # Application entry point│   ├── handlers/

├── internal/│   │   └── auth.go             # ✅ Authentication endpoints

│   ├── config/               # Configuration management│   ├── middleware/

│   ├── handlers/             # HTTP request handlers│   │   └── auth.go             # JWT authentication

│   │   ├── auth.go          # Authentication endpoints│   ├── models/                  # Database models (GORM)

│   │   ├── journey.go       # Journey CRUD│   │   ├── user.go

│   │   ├── scenario.go      # Scenario CRUD│   │   ├── journey.go

│   │   └── word.go          # Word CRUD│   │   ├── scenario.go

│   ├── services/             # Business logic│   │   └── ...

│   │   ├── auth.service.go│   ├── repository/

│   │   ├── journey.service.go│   │   └── user.repo.go        # ✅ User data access

│   │   ├── scenario.service.go│   ├── services/

│   │   └── word.service.go│   │   └── auth.service.go     # ✅ Auth business logic

│   ├── repository/           # Data access layer│   └── utils/                   # Utilities (JWT, validation, etc.)

│   │   ├── user.repo.go├── docs/                        # API documentation

│   │   ├── journey.repo.go│   └── api.md

│   │   ├── scenario.repo.go├── .env                         # Environment variables

│   │   └── word.repo.go├── test-auth.sh                # ✅ Authentication tests

│   ├── models/               # GORM models├── Makefile                     # Build commands

│   ├── middleware/           # Auth, CORS, logging└── learng.db                   # SQLite database (auto-created)

│   └── utils/                # Helper functions```

├── uploads/                  # Media file storage

├── test-sprint2.sh          # API test suite## 🔐 API Endpoints

└── PROJECT_STATUS.md         # Detailed project status

```### Public Endpoints

- `GET /health` - Health check

## API Endpoints- `POST /api/v1/auth/register` - Register new user

- `POST /api/v1/auth/login` - Login user

### Authentication

```### Protected Endpoints (Requires JWT)

POST   /api/v1/auth/register  - Create user account- `GET /api/v1/auth/me` - Get current user

POST   /api/v1/auth/login     - Login and get JWT token

GET    /api/v1/auth/me        - Get current user (protected)### Coming Soon

```- Journey CRUD endpoints

- Scenario management

### Journeys (Protected)- Word management

```- Media upload

POST   /api/v1/journeys       - Create journey- Quiz system

GET    /api/v1/journeys       - List journeys (paginated, filterable)

GET    /api/v1/journeys/:id   - Get journey with scenarios and words## 🧪 Testing

PUT    /api/v1/journeys/:id   - Update journey

DELETE /api/v1/journeys/:id   - Delete journey (cascade)### Automated Tests

``````bash

./test-auth.sh

### Scenarios (Protected)```

```

POST   /api/v1/scenarios      - Create scenarioTests include:

GET    /api/v1/scenarios/:id  - Get scenario with words- ✅ User registration (admin/learner)

PUT    /api/v1/scenarios/:id  - Update scenario- ✅ Login flow

DELETE /api/v1/scenarios/:id  - Delete scenario (cascade)- ✅ Protected endpoint access

```- ✅ Duplicate email prevention

- ✅ Input validation (email, password)

### Words (Protected)- ✅ Error handling

```

POST   /api/v1/words          - Create word### Manual Testing

GET    /api/v1/words/:id      - Get word```bash

PUT    /api/v1/words/:id      - Update word# Register

DELETE /api/v1/words/:id      - Delete wordcurl -X POST http://localhost:8080/api/v1/auth/register \

```  -H "Content-Type: application/json" \

  -d '{"email":"user@test.com","password":"password123","displayName":"Test User","role":"admin"}'

## Quick API Examples

# Login

### Register Usercurl -X POST http://localhost:8080/api/v1/auth/login \

```bash  -H "Content-Type: application/json" \

curl -X POST http://localhost:8080/api/v1/auth/register \  -d '{"email":"user@test.com","password":"password123"}'

  -H "Content-Type: application/json" \

  -d '{"email":"admin@test.com","password":"Test1234","displayName":"Admin","role":"admin"}'# Get current user (use token from login)

```curl -H "Authorization: Bearer YOUR_TOKEN" \

  http://localhost:8080/api/v1/auth/me

### Login```

```bash

curl -X POST http://localhost:8080/api/v1/auth/login \## 📚 Documentation

  -H "Content-Type: application/json" \

  -d '{"email":"admin@test.com","password":"Test1234"}'- [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md) - Quick start and debugging

```- [AUTH_IMPLEMENTATION.md](AUTH_IMPLEMENTATION.md) - Authentication details

- [TEST_RESULTS.md](TEST_RESULTS.md) - Latest test results

### Create Journey- [docs/api.md](docs/api.md) - API documentation

```bash

TOKEN="your-jwt-token"## 🏗️ Architecture

curl -X POST http://localhost:8080/api/v1/journeys \

  -H "Authorization: Bearer $TOKEN" \**Three-Layer Architecture**:

  -H "Content-Type: application/json" \1. **Handler Layer** - HTTP request/response handling

  -d '{"title":"At the Park","sourceLanguage":"en","targetLanguage":"zh-HK"}'2. **Service Layer** - Business logic and validation

```3. **Repository Layer** - Database operations



## Testing**Key Features**:

- Dependency injection in `main.go`

### Run All Tests- Middleware for authentication and CORS

```bash- Structured error responses

./test-sprint2.sh- Input validation utilities

```- JWT token management



This tests:## 🔧 Development

- Authentication flow

- Journey CRUD operations### Make Commands

- Scenario CRUD operations```bash

- Word CRUD operationsmake build    # Build the application

- Nested data loadingmake run      # Run the application

- Cascade deletionsmake test     # Run tests

make clean    # Remove build artifacts

### Expected Output```

```

✅ All Sprint 2 API tests completed!### Environment Variables



Summary:Create `.env` file:

- Journey CRUD: ✅```env

- Scenario CRUD: ✅APP_ENV=development

- Word CRUD: ✅APP_PORT=8080

- Nested data retrieval: ✅DB_DRIVER=sqlite

- Cascade deletions: ✅DB_PATH=./learng.db

```JWT_SECRET=your-secret-key-here

UPLOAD_DIR=./uploads

## DatabaseMAX_FILE_SIZE=10485760

ALLOWED_ORIGINS=http://localhost:5173

### Schema Overview```

```

users (auth)### Database

  ↓

journeys (1) → scenarios (*) → words (*)View database contents:

  ↓```bash

quizzes → quiz_questionssqlite3 learng.db

  ↓> .tables

learner_progress, quiz_attempts> SELECT * FROM users;

```> .quit

```

### Reset Database

```bash## 🚀 Next Steps

rm learng.db

# Restart server - database auto-recreates with migrations### Sprint 2 (Weeks 3-4)

```- [ ] Journey CRUD endpoints

- [ ] Scenario management

### View Database- [ ] Word management

```bash- [ ] Admin-only routes with role middleware

sqlite3 learng.db

sqlite> .tables### Sprint 3 (Weeks 5-6)

sqlite> SELECT * FROM journeys;- [ ] Media upload handling

sqlite> .quit- [ ] AI service integration stubs

```- [ ] Quiz system



## Environment Variables### Future

- [ ] PostgreSQL migration

```bash- [ ] Docker deployment

# Server- [ ] API rate limiting

PORT=8080- [ ] Refresh token mechanism



# Database## 📝 Contributing

DB_PATH=./learng.db

1. Follow the three-layer architecture pattern

# Security (REQUIRED - generate random string)2. Add tests for new endpoints

JWT_SECRET=your-secret-key-here3. Update documentation

4. Use the existing error handling patterns

# File Storage

UPLOAD_DIR=./uploads## 🐛 Troubleshooting



# Static Files (empty in development)**Port already in use**:

STATIC_DIR=```bash

```lsof -ti:8080 | xargs kill -9

```

## Architecture

**Database locked**:

### 3-Layer Design```bash

```# Kill all instances and restart

HTTP Requestpkill api

    ↓./bin/api

Handler (HTTP layer)```

    ↓

Service (Business logic)**Missing JWT_SECRET**:

    ↓```bash

Repository (Data access)# Create .env file with JWT_SECRET

    ↓echo "JWT_SECRET=dev-secret-change-in-production" >> .env

Database```

```

---

### Key Features

- **JWT Authentication**: Secure token-based auth**Last Updated**: 2025-10-07  

- **Role-Based Access**: Admin vs learner permissions**Current Version**: v0.1.0-alpha  

- **Nested Loading**: Efficient data retrieval with GORM preloading**Status**: ✅ Authentication Complete | 🚧 Journey Management In Progress

- **Cascade Deletes**: Automatic cleanup of related records
- **Ownership Validation**: Only creators can modify their content
- **Pagination**: Efficient list endpoints
- **Soft Deletes**: GORM DeletedAt for data recovery

## Common Tasks

### Add New Endpoint
1. Create/update model in `internal/models/`
2. Add repository methods in `internal/repository/`
3. Add service logic in `internal/services/`
4. Add handler in `internal/handlers/`
5. Register route in `cmd/api/main.go`
6. Add tests to test script

### Troubleshooting

**Port 8080 already in use:**
```bash
lsof -ti:8080 | xargs kill -9
```

**Database locked:**
```bash
# Stop server, then:
rm learng.db  # Reset if corrupted
```

**Token expired:**
```bash
# Re-login to get fresh token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Test1234"}'
```

## Performance

Typical response times (local SQLite):
- Journey creation: ~5ms
- Nested data retrieval: ~15ms
- Word creation: ~3ms
- Updates: ~5ms
- All operations: <20ms

## Deployment

See `design/CORE.md` for production deployment with Docker.

## Documentation

- `README.md` - This file (setup and API reference)
- `PROJECT_STATUS.md` - Detailed project status and sprint progress
- `design/CORE.md` - Complete technical specification
- `test-sprint2.sh` - Executable API tests (self-documenting)

## Contributing

1. Follow 3-layer architecture (handler → service → repository)
2. Add tests for new endpoints
3. Update PROJECT_STATUS.md with progress
4. Keep code comments for complex logic

---

**Status:** Sprint 2 Complete ✅  
**Next:** Sprint 3 - Media upload endpoints
