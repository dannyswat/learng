# Learng Backend# learng Backend# Learng Backend



Go backend API server for the learng AI-powered language learning platform.



## 🎯 Current StatusGo backend API server for the learng language learning platform.AI-powered language learning platform backend built with Go and Echo framework.



✅ **Sprint 1 Complete**: Authentication system

✅ **Sprint 2 Complete**: Content Management + Media Upload

## Tech Stack## 🎯 Current Status

## Tech Stack



- **Go**: 1.21+

- **Web Framework**: Echo v4.13.4- **Language**: Go 1.21+✅ **Phase 1 - Sprint 1 Complete**

- **ORM**: GORM v1.31.0

- **Database**: SQLite 3.x (development) / PostgreSQL (production)- **Framework**: Echo v4 (lightweight web framework)- Authentication system fully implemented and tested

- **Authentication**: JWT (golang-jwt/jwt v5.1.0)

- **Password Hashing**: bcrypt- **ORM**: GORM (database operations)- User registration and login working



## 🚀 Quick Start- **Database**: SQLite (development) → PostgreSQL (production)- JWT-based authentication



### Prerequisites- **Auth**: JWT-based authentication- Protected endpoints with middleware

- Go 1.21 or higher

- SQLite3- **Testing**: Shell scripts + Go tests- Input validation and error handling



### Installation & Running



```bash## Quick Start## Tech Stack

# Install dependencies

go mod download



# Build the application### Prerequisites- **Go**: 1.21+

make build

- Go 1.21 or higher- **Web Framework**: Echo v4.11.3

# Run the server

make run- SQLite3- **ORM**: GORM v1.25.5

# or

./learng-api- **Database**: SQLite 3.x (MVP) / PostgreSQL (Production)

```

### Setup- **Authentication**: JWT (golang-jwt/jwt v5.1.0)

Server will start on `http://localhost:8080`

- **Password Hashing**: bcrypt

### Environment Variables

```bash

Create a `.env` file:

# Clone and navigate## Quick Start

```env

PORT=8080cd backend

DATABASE_PATH=./learng.db

UPLOAD_DIR=./uploads### Prerequisites

JWT_SECRET=your-secret-key-here-min-32-chars

STATIC_DIR=  # Leave empty for dev mode# Install dependencies

```

go mod download- Go 1.21 or higher

## 📁 Project Structure

- SQLite3

```

backend/# Copy environment template

├── cmd/api/main.go              # Application entry point

├── internal/cp .env.example .env### Installation & Running

│   ├── config/                  # Configuration loading

│   ├── handlers/                # HTTP request handlers

│   │   ├── auth.go             # ✅ Authentication endpoints

│   │   ├── journey.go          # ✅ Journey CRUD# Edit .env with your settings```bash

│   │   ├── scenario.go         # ✅ Scenario CRUD

│   │   ├── word.go             # ✅ Word CRUD# Required: JWT_SECRET (generate a random string)# Install dependencies

│   │   └── media.go            # ✅ Media upload endpoints

│   ├── middleware/go mod download

│   │   └── auth.go             # JWT authentication middleware

│   ├── models/                  # Database models (GORM)# Build

│   │   ├── user.go

│   │   ├── journey.gogo build -o bin/api ./cmd/api# Build the application

│   │   ├── scenario.go

│   │   ├── word.gomake build

│   │   ├── quiz.go

│   │   ├── quiz_attempt.go# Run

│   │   └── progress.go

│   ├── repository/              # Data access layerbin/api# Run the server

│   │   ├── user.repo.go        # ✅ User database operations

│   │   ├── journey.repo.go     # ✅ Journey database operations# Server starts on http://localhost:8080make run

│   │   ├── scenario.repo.go    # ✅ Scenario database operations

│   │   └── word.repo.go        # ✅ Word database operations```# or

│   ├── services/                # Business logic layer

│   │   ├── auth.service.go     # ✅ Auth business logic./bin/api

│   │   ├── journey.service.go  # ✅ Journey business logic

│   │   ├── scenario.service.go # ✅ Scenario business logic### Development```

│   │   └── word.service.go     # ✅ Word business logic

│   └── utils/                   # Helper functions

│       ├── jwt.go

│       ├── password.go```bashServer will start on http://localhost:8080

│       ├── response.go

│       └── validation.go# Run with hot reload (if using air)

├── test-auth.sh                 # Auth test suite

├── test-sprint2.sh              # Content CRUD test suiteair### Test Authentication

├── test-media.sh                # Media upload test suite

└── Makefile

```

# Run tests```bash

## 📡 API Endpoints

./test-sprint2.sh# Run automated test suite

### Health Check

```./test-auth.sh

GET  /health  - Server health check

```# Build```



### Authentication (Public)make build

```

POST /api/v1/auth/register  - Register new userExpected output: ✅ All 10 tests pass

POST /api/v1/auth/login     - Login user

```# Clean build artifacts



### Auth (Protected)make clean## 📁 Project Structure

```

GET  /api/v1/auth/me        - Get current user info```

```

```

### Journeys (Protected)

```## Project Structurebackend/

POST   /api/v1/journeys       - Create journey

GET    /api/v1/journeys       - List journeys (paginated, filterable)├── cmd/api/main.go              # Application entry point

GET    /api/v1/journeys/:id   - Get journey with scenarios and words

PUT    /api/v1/journeys/:id   - Update journey```├── internal/

DELETE /api/v1/journeys/:id   - Delete journey (cascade)

```backend/│   ├── config/                  # Configuration loading



### Scenarios (Protected)├── cmd/api/main.go           # Application entry point│   ├── handlers/

```

POST   /api/v1/scenarios      - Create scenario├── internal/│   │   └── auth.go             # ✅ Authentication endpoints

GET    /api/v1/scenarios/:id  - Get scenario with words

PUT    /api/v1/scenarios/:id  - Update scenario│   ├── config/               # Configuration management│   ├── middleware/

DELETE /api/v1/scenarios/:id  - Delete scenario (cascade)

```│   ├── handlers/             # HTTP request handlers│   │   └── auth.go             # JWT authentication



### Words (Protected)│   │   ├── auth.go          # Authentication endpoints│   ├── models/                  # Database models (GORM)

```

POST   /api/v1/words          - Create word│   │   ├── journey.go       # Journey CRUD│   │   ├── user.go

GET    /api/v1/words/:id      - Get word

PUT    /api/v1/words/:id      - Update word│   │   ├── scenario.go      # Scenario CRUD│   │   ├── journey.go

DELETE /api/v1/words/:id      - Delete word

```│   │   └── word.go          # Word CRUD│   │   ├── scenario.go



### Media Upload (Protected)│   ├── services/             # Business logic│   │   └── ...

```

POST /api/v1/media/upload/image  - Upload image file│   │   ├── auth.service.go│   ├── repository/

POST /api/v1/media/upload/audio  - Upload audio file

```│   │   ├── journey.service.go│   │   └── user.repo.go        # ✅ User data access



**Image Upload:**│   │   ├── scenario.service.go│   ├── services/

- **Supported formats**: JPEG, PNG, WebP

- **Max size**: 5MB│   │   └── word.service.go│   │   └── auth.service.go     # ✅ Auth business logic

- **Content-Type**: multipart/form-data

- **Form field**: `file`│   ├── repository/           # Data access layer│   └── utils/                   # Utilities (JWT, validation, etc.)

- **Response**: `{ "url": "/uploads/images/uuid.jpg", "filename": "uuid.jpg", "size": 12345, "mimeType": "image/jpeg" }`

│   │   ├── user.repo.go├── docs/                        # API documentation

**Audio Upload:**

- **Supported formats**: MP3, WAV, WebM│   │   ├── journey.repo.go│   └── api.md

- **Max size**: 2MB

- **Content-Type**: multipart/form-data│   │   ├── scenario.repo.go├── .env                         # Environment variables

- **Form field**: `file`

- **Response**: `{ "url": "/uploads/audio/uuid.mp3", "filename": "uuid.mp3", "size": 54321, "mimeType": "audio/mpeg" }`│   │   └── word.repo.go├── test-auth.sh                # ✅ Authentication tests



### Static Files│   ├── models/               # GORM models├── Makefile                     # Build commands

```

GET  /uploads/images/:filename  - Serve uploaded images│   ├── middleware/           # Auth, CORS, logging└── learng.db                   # SQLite database (auto-created)

GET  /uploads/audio/:filename   - Serve uploaded audio files

```│   └── utils/                # Helper functions```



## 📝 API Usage Examples├── uploads/                  # Media file storage



### 1. Register User├── test-sprint2.sh          # API test suite## 🔐 API Endpoints

```bash

curl -X POST http://localhost:8080/api/v1/auth/register \└── PROJECT_STATUS.md         # Detailed project status

  -H "Content-Type: application/json" \

  -d '{```### Public Endpoints

    "email": "admin@example.com",

    "password": "securepass123",- `GET /health` - Health check

    "displayName": "Admin User",

    "role": "admin"## API Endpoints- `POST /api/v1/auth/register` - Register new user

  }'

```- `POST /api/v1/auth/login` - Login user



### 2. Login### Authentication

```bash

curl -X POST http://localhost:8080/api/v1/auth/login \```### Protected Endpoints (Requires JWT)

  -H "Content-Type: application/json" \

  -d '{POST   /api/v1/auth/register  - Create user account- `GET /api/v1/auth/me` - Get current user

    "email": "admin@example.com",

    "password": "securepass123"POST   /api/v1/auth/login     - Login and get JWT token

  }'

```GET    /api/v1/auth/me        - Get current user (protected)### Coming Soon



Response:```- Journey CRUD endpoints

```json

{- Scenario management

  "token": "eyJhbGciOiJIUzI1NiIs...",

  "user": {### Journeys (Protected)- Word management

    "id": "uuid",

    "email": "admin@example.com",```- Media upload

    "role": "admin",

    "displayName": "Admin User"POST   /api/v1/journeys       - Create journey- Quiz system

  }

}GET    /api/v1/journeys       - List journeys (paginated, filterable)

```

GET    /api/v1/journeys/:id   - Get journey with scenarios and words## 🧪 Testing

### 3. Create Journey

```bashPUT    /api/v1/journeys/:id   - Update journey

TOKEN="your-jwt-token"

DELETE /api/v1/journeys/:id   - Delete journey (cascade)### Automated Tests

curl -X POST http://localhost:8080/api/v1/journeys \

  -H "Authorization: Bearer $TOKEN" \``````bash

  -H "Content-Type: application/json" \

  -d '{./test-auth.sh

    "title": "At the Park",

    "description": "Learn outdoor vocabulary",### Scenarios (Protected)```

    "sourceLanguage": "en",

    "targetLanguage": "zh-HK"```

  }'

```POST   /api/v1/scenarios      - Create scenarioTests include:



### 4. Upload ImageGET    /api/v1/scenarios/:id  - Get scenario with words- ✅ User registration (admin/learner)

```bash

curl -X POST http://localhost:8080/api/v1/media/upload/image \PUT    /api/v1/scenarios/:id  - Update scenario- ✅ Login flow

  -H "Authorization: Bearer $TOKEN" \

  -F "file=@/path/to/image.jpg"DELETE /api/v1/scenarios/:id  - Delete scenario (cascade)- ✅ Protected endpoint access

```

```- ✅ Duplicate email prevention

### 5. Upload Audio

```bash- ✅ Input validation (email, password)

curl -X POST http://localhost:8080/api/v1/media/upload/audio \

  -H "Authorization: Bearer $TOKEN" \### Words (Protected)- ✅ Error handling

  -F "file=@/path/to/audio.mp3"

``````



### 6. Create Word with MediaPOST   /api/v1/words          - Create word### Manual Testing

```bash

curl -X POST http://localhost:8080/api/v1/words \GET    /api/v1/words/:id      - Get word```bash

  -H "Authorization: Bearer $TOKEN" \

  -H "Content-Type: application/json" \PUT    /api/v1/words/:id      - Update word# Register

  -d '{

    "scenarioId": "scenario-uuid",DELETE /api/v1/words/:id      - Delete wordcurl -X POST http://localhost:8080/api/v1/auth/register \

    "sourceText": "apple",

    "targetText": "蘋果",```  -H "Content-Type: application/json" \

    "imageUrl": "/uploads/images/abc123.jpg",

    "audioUrl": "/uploads/audio/def456.mp3",  -d '{"email":"user@test.com","password":"password123","displayName":"Test User","role":"admin"}'

    "generationMethod": "manual"

  }'## Quick API Examples

```

# Login

## 🧪 Testing

### Register Usercurl -X POST http://localhost:8080/api/v1/auth/login \

### Automated Test Suites

```bash  -H "Content-Type: application/json" \

```bash

# Test authenticationcurl -X POST http://localhost:8080/api/v1/auth/register \  -d '{"email":"user@test.com","password":"password123"}'

./test-auth.sh

  -H "Content-Type: application/json" \

# Test content CRUD (journeys, scenarios, words)

./test-sprint2.sh  -d '{"email":"admin@test.com","password":"Test1234","displayName":"Admin","role":"admin"}'# Get current user (use token from login)



# Test media upload```curl -H "Authorization: Bearer YOUR_TOKEN" \

./test-media.sh

```  http://localhost:8080/api/v1/auth/me



**Test Coverage:**### Login```

- ✅ User registration (admin/learner)

- ✅ Login flow```bash

- ✅ Protected endpoint access

- ✅ Journey CRUD (5 endpoints, 17 scenarios)curl -X POST http://localhost:8080/api/v1/auth/login \## 📚 Documentation

- ✅ Scenario CRUD with cascade deletion

- ✅ Word CRUD with media URLs  -H "Content-Type: application/json" \

- ✅ Media upload (image: PNG/JPEG, audio: WAV/MP3/WebM)

- ✅ File validation (size, type, extension)  -d '{"email":"admin@test.com","password":"Test1234"}'- [DEVELOPER_GUIDE.md](DEVELOPER_GUIDE.md) - Quick start and debugging

- ✅ Authentication required for uploads

- ✅ Uploaded files accessible via /uploads```- [AUTH_IMPLEMENTATION.md](AUTH_IMPLEMENTATION.md) - Authentication details



All test suites should show: **✅ All tests passed!**- [TEST_RESULTS.md](TEST_RESULTS.md) - Latest test results



## 🗄️ Database### Create Journey- [docs/api.md](docs/api.md) - API documentation



### Schema```bash

- **users**: User accounts with roles (admin/learner)

- **journeys**: Learning journeys with language pairsTOKEN="your-jwt-token"## 🏗️ Architecture

- **scenarios**: Themed word collections within journeys

- **words**: Individual vocabulary items with media URLscurl -X POST http://localhost:8080/api/v1/journeys \

- **quizzes**: Quiz templates for scenarios (placeholder)

- **quiz_attempts**: User quiz results (placeholder)  -H "Authorization: Bearer $TOKEN" \**Three-Layer Architecture**:

- **learner_progress**: User word mastery tracking (placeholder)

  -H "Content-Type: application/json" \1. **Handler Layer** - HTTP request/response handling

### SQLite Commands

```bash  -d '{"title":"At the Park","sourceLanguage":"en","targetLanguage":"zh-HK"}'2. **Service Layer** - Business logic and validation

# Open database

sqlite3 learng.db```3. **Repository Layer** - Database operations



# View tables

.tables

## Testing**Key Features**:

# Query users

SELECT * FROM users;- Dependency injection in `main.go`



# View journey with scenarios### Run All Tests- Middleware for authentication and CORS

SELECT j.title as journey, s.title as scenario 

FROM journeys j ```bash- Structured error responses

LEFT JOIN scenarios s ON s.journey_id = j.id;

```./test-sprint2.sh- Input validation utilities



## 🛠️ Development```- JWT token management



### Makefile Commands



```bashThis tests:## 🔧 Development

make build       # Build the application

make run         # Run the server- Authentication flow

make clean       # Remove build artifacts

make test        # Run all test suites- Journey CRUD operations### Make Commands

```

- Scenario CRUD operations```bash

### Manual Development

- Word CRUD operationsmake build    # Build the application

```bash

# Build- Nested data loadingmake run      # Run the application

go build -o learng-api ./cmd/api

- Cascade deletionsmake test     # Run tests

# Run

./learng-apimake clean    # Remove build artifacts



# Run with air for hot reload (if installed)### Expected Output```

air

``````



## 🚨 Troubleshooting✅ All Sprint 2 API tests completed!### Environment Variables



### Port 8080 already in use

```bash

# Kill existing processSummary:Create `.env` file:

lsof -ti:8080 | xargs kill -9

- Journey CRUD: ✅```env

# Or use different port in .env

PORT=8081- Scenario CRUD: ✅APP_ENV=development

```

- Word CRUD: ✅APP_PORT=8080

### Database locked

```bash- Nested data retrieval: ✅DB_DRIVER=sqlite

# Stop all running instances

pkill -9 learng-api- Cascade deletions: ✅DB_PATH=./learng.db



# Remove database and restart```JWT_SECRET=your-secret-key-here

rm learng.db

./learng-apiUPLOAD_DIR=./uploads

```

## DatabaseMAX_FILE_SIZE=10485760

### Upload directory permission denied

```bashALLOWED_ORIGINS=http://localhost:5173

# Create upload directories with correct permissions

mkdir -p uploads/images uploads/audio### Schema Overview```

chmod 755 uploads uploads/images uploads/audio

``````



### JWT token invalidusers (auth)### Database

- Ensure JWT_SECRET in .env is at least 32 characters

- Token expires after 24 hours - re-login to get new token  ↓



## 📦 Deploymentjourneys (1) → scenarios (*) → words (*)View database contents:



### Production Build  ↓```bash

```bash

# Build optimized binaryquizzes → quiz_questionssqlite3 learng.db

CGO_ENABLED=1 GOOS=linux go build -o learng-api ./cmd/api

  ↓> .tables

# Run in production mode

export STATIC_DIR=/path/to/frontend/distlearner_progress, quiz_attempts> SELECT * FROM users;

./learng-api

``````> .quit



### Environment Variables (Production)```

```env

PORT=8080### Reset Database

DATABASE_PATH=/data/learng.db

UPLOAD_DIR=/data/uploads```bash## 🚀 Next Steps

JWT_SECRET=<strong-random-secret-min-32-chars>

STATIC_DIR=/app/frontend/distrm learng.db

```

# Restart server - database auto-recreates with migrations### Sprint 2 (Weeks 3-4)

## 🎯 Sprint Progress

```- [ ] Journey CRUD endpoints

### ✅ Sprint 1: Authentication & Foundation

- [x] User registration and login- [ ] Scenario management

- [x] JWT authentication

- [x] Protected routes middleware### View Database- [ ] Word management

- [x] Password hashing

- [x] Input validation```bash- [ ] Admin-only routes with role middleware

- [x] Error handling

- [x] Test suite (10 scenarios)sqlite3 learng.db



### ✅ Sprint 2: Content Management + Media Uploadsqlite> .tables### Sprint 3 (Weeks 5-6)

- [x] Journey CRUD (5 endpoints)

- [x] Scenario CRUD (4 endpoints)sqlite> SELECT * FROM journeys;- [ ] Media upload handling

- [x] Word CRUD (4 endpoints)

- [x] Media upload endpoints (2 endpoints)sqlite> .quit- [ ] AI service integration stubs

- [x] File validation (size, type, extension)

- [x] Static file serving```- [ ] Quiz system

- [x] 3-layer architecture (handlers → services → repositories)

- [x] Nested data loading

- [x] Cascade deletions

- [x] Ownership validation## Environment Variables### Future

- [x] Pagination & filtering

- [x] Test suites (28 total scenarios)- [ ] PostgreSQL migration



### 🔜 Sprint 3: Learner Features (Next)```bash- [ ] Docker deployment

- [ ] Learner journey listing

- [ ] Word card viewing# Server- [ ] API rate limiting

- [ ] Progress tracking

- [ ] Quiz endpointsPORT=8080- [ ] Refresh token mechanism

- [ ] Quiz attempt tracking



## 📝 Notes

# Database## 📝 Contributing

- **Media Storage**: Currently using local filesystem (`./uploads`). Production should use cloud storage (Azure Blob/S3).

- **Database**: SQLite for development. Migrate to PostgreSQL for production.DB_PATH=./learng.db

- **CORS**: Currently allows all origins. Configure for production domains.

- **Rate Limiting**: Not implemented yet. Add for production.1. Follow the three-layer architecture pattern

- **AI Generation**: Not implemented yet. Current upload is manual only.

# Security (REQUIRED - generate random string)2. Add tests for new endpoints

## 🔗 Related Docs

JWT_SECRET=your-secret-key-here3. Update documentation

- [Project Status](../PROJECT_STATUS.md)

- [Design Specification](../design/CORE.md)4. Use the existing error handling patterns

- [Requirements](../REQUIREMENT.md)

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
