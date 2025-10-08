# learng - Sprint 2 Complete! 🎉

## What's New

Sprint 2 delivers the complete backend CRUD APIs for content management. Admins can now create, read, update, and delete journeys, scenarios, and words through RESTful APIs.

## Sprint 2 Highlights

### ✅ Journey Management APIs
- Create learning journeys with source/target language pairs
- List journeys with pagination and filtering
- Get complete journey data with nested scenarios and words
- Update journey details and status
- Delete journeys (cascades to scenarios and words)

### ✅ Scenario Management APIs
- Create themed scenarios within journeys
- Organize scenarios with display order
- Get scenarios with all associated words
- Update scenario details
- Delete scenarios (cascades to words)

### ✅ Word Management APIs
- Create vocabulary words with translations
- Support for manual media URLs (images, audio)
- Update word details and media references
- Delete individual words
- Track generation method (manual/AI)

### 🏗️ Architecture Improvements
- **3-Layer Architecture**: Handlers → Services → Repositories
- **Data Validation**: Input validation at service layer
- **Ownership Checks**: Only creators can modify/delete content
- **Nested Loading**: Optimized queries with GORM preloading
- **Cascade Deletes**: Automatic cleanup of related data

## Quick Start

### Backend Server
```bash
cd backend
go build -o bin/api ./cmd/api
bin/api
# Server: http://localhost:8080
```

### Run Tests
```bash
cd backend
./test-sprint2.sh
# Tests all CRUD operations
```

### Example API Usage
```bash
# 1. Register admin user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Test1234","role":"admin"}'

# 2. Create journey
curl -X POST http://localhost:8080/api/v1/journeys \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"title":"At the Park","sourceLanguage":"en","targetLanguage":"zh-HK"}'

# 3. Create scenario
curl -X POST http://localhost:8080/api/v1/scenarios \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"journeyId":"<journey-id>","title":"Colors","displayOrder":1}'

# 4. Create word
curl -X POST http://localhost:8080/api/v1/words \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"scenarioId":"<scenario-id>","targetText":"紅色","sourceText":"Red"}'
```

## Test Results

All Sprint 2 tests passing ✅:

```
✅ User registration and authentication
✅ Journey CRUD operations
✅ Scenario CRUD operations
✅ Word CRUD operations
✅ Nested data retrieval
✅ Pagination and filtering
✅ Cascade deletions
✅ Ownership validation
```

## Documentation

📚 **Backend Documentation:**
- `backend/SPRINT2_SUMMARY.md` - Detailed feature documentation
- `backend/SPRINT2_QUICK_REF.md` - Developer quick reference
- `backend/PROJECT_STATUS.md` - Project status and roadmap
- `backend/test-sprint2.sh` - Comprehensive API tests

📋 **Design Docs:**
- `design/CORE.md` - Complete technical specification
- `.github/copilot-instructions.md` - AI agent context

## API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register user
- `POST /api/v1/auth/login` - Login & get JWT token
- `GET /api/v1/auth/me` - Get current user

### Journeys (Protected)
- `POST /api/v1/journeys` - Create
- `GET /api/v1/journeys` - List (paginated)
- `GET /api/v1/journeys/:id` - Get with nested data
- `PUT /api/v1/journeys/:id` - Update
- `DELETE /api/v1/journeys/:id` - Delete

### Scenarios (Protected)
- `POST /api/v1/scenarios` - Create
- `GET /api/v1/scenarios/:id` - Get with words
- `PUT /api/v1/scenarios/:id` - Update
- `DELETE /api/v1/scenarios/:id` - Delete

### Words (Protected)
- `POST /api/v1/words` - Create
- `GET /api/v1/words/:id` - Get
- `PUT /api/v1/words/:id` - Update
- `DELETE /api/v1/words/:id` - Delete

## Database Schema

```
journeys (1) → scenarios (*) → words (*)
    ↓
  users (creator)
```

All entities use:
- UUID primary keys
- Soft deletes (GORM DeletedAt)
- Timestamps (CreatedAt, UpdatedAt)
- Foreign key constraints with CASCADE DELETE

## Project Structure

```
learng/
├── backend/              ✅ Sprint 1 & 2 Complete
│   ├── cmd/api/          # Entry point
│   ├── internal/
│   │   ├── handlers/     # HTTP handlers
│   │   ├── services/     # Business logic
│   │   ├── repository/   # Data access
│   │   ├── models/       # GORM entities
│   │   ├── middleware/   # Auth, logging
│   │   └── utils/        # Helpers
│   └── test-sprint2.sh   # API tests
├── frontend/             ⏳ Sprint 2 (Frontend) - Next
│   ├── src/
│   │   ├── pages/
│   │   ├── components/
│   │   └── services/
├── design/               ✅ Complete
│   ├── CORE.md           # Technical spec
│   └── SUMMARY.md
└── .github/
    └── copilot-instructions.md
```

## Technology Stack

**Backend:**
- Go 1.21+
- Echo (web framework)
- GORM (ORM)
- SQLite (dev) → PostgreSQL (production path)
- JWT authentication

**Frontend (Next Sprint):**
- React 18 + TypeScript
- Vite
- Tailwind CSS
- Axios

## Next Steps: Sprint 3

**Media Handling (Week 5-6):**
- [ ] Media upload endpoints (image, audio)
- [ ] File validation (size, type, format)
- [ ] Audio recorder component (MediaRecorder API)
- [ ] Image uploader component (drag-drop)
- [ ] Admin UI for journey/scenario/word management

**Timeline:**
- Sprint 3 Backend: Media upload APIs
- Sprint 3 Frontend: Media recording/upload components
- Sprint 4: Learner experience (card view, navigation)
- Sprint 5: Quiz system

## Development Commands

```bash
# Backend
cd backend
go build -o bin/api ./cmd/api     # Build
bin/api                            # Run server
./test-sprint2.sh                  # Test APIs

# Database
sqlite3 backend/learng.db          # View database
rm backend/learng.db               # Reset (auto-recreates)

# Health check
curl http://localhost:8080/health
```

## Key Features

✅ **RESTful APIs** - Clean, predictable endpoints  
✅ **JWT Auth** - Secure token-based authentication  
✅ **Pagination** - Efficient data loading  
✅ **Filtering** - Query by status, language, etc.  
✅ **Nested Loading** - Journey → Scenarios → Words in one call  
✅ **Cascade Delete** - Automatic cleanup of related data  
✅ **Ownership Validation** - Only creators can modify content  
✅ **Error Handling** - Consistent error responses  
✅ **Test Coverage** - Comprehensive automated tests  

## Contributing

When working on new features:
1. Check `design/CORE.md` for specifications
2. Follow 3-layer architecture (handler → service → repository)
3. Add tests to `test-sprint2.sh` (or create new test file)
4. Update documentation

## Questions?

- Technical Spec: `design/CORE.md`
- Sprint Details: `backend/SPRINT2_SUMMARY.md`
- Quick Reference: `backend/SPRINT2_QUICK_REF.md`
- API Examples: `backend/test-sprint2.sh`

---

**Sprint 2 Status: ✅ COMPLETE**  
**Ready for Sprint 3: Media Handling**
