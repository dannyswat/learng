# Sprint 2 Quick Reference - Journey/Scenario/Word APIs

## Quick Start

### Start Backend Server
```bash
cd backend
go build -o bin/api ./cmd/api
bin/api
# Server runs on http://localhost:8080
```

### Run Tests
```bash
cd backend
./test-sprint2.sh
```

## API Endpoints Quick Reference

### Authentication (Required for all endpoints below)
```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Test1234","displayName":"Admin","role":"admin"}'

# Login & Get Token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Test1234"}'

# Set token for subsequent requests
TOKEN="your-jwt-token-here"
```

### Journeys

```bash
# Create Journey
curl -X POST http://localhost:8080/api/v1/journeys \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"At the Park","description":"Learn outdoor vocabulary","sourceLanguage":"en","targetLanguage":"zh-HK"}'

# List Journeys (with filtering)
curl -X GET "http://localhost:8080/api/v1/journeys?status=draft&page=1&limit=10" \
  -H "Authorization: Bearer $TOKEN"

# Get Journey (with scenarios and words)
curl -X GET http://localhost:8080/api/v1/journeys/{id} \
  -H "Authorization: Bearer $TOKEN"

# Update Journey
curl -X PUT http://localhost:8080/api/v1/journeys/{id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","status":"published"}'

# Delete Journey
curl -X DELETE http://localhost:8080/api/v1/journeys/{id} \
  -H "Authorization: Bearer $TOKEN"
```

### Scenarios

```bash
# Create Scenario
curl -X POST http://localhost:8080/api/v1/scenarios \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"journeyId":"journey-uuid","title":"Colors","description":"Learn colors","displayOrder":1}'

# Get Scenario (with words)
curl -X GET http://localhost:8080/api/v1/scenarios/{id} \
  -H "Authorization: Bearer $TOKEN"

# Update Scenario
curl -X PUT http://localhost:8080/api/v1/scenarios/{id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","displayOrder":2}'

# Delete Scenario
curl -X DELETE http://localhost:8080/api/v1/scenarios/{id} \
  -H "Authorization: Bearer $TOKEN"
```

### Words

```bash
# Create Word
curl -X POST http://localhost:8080/api/v1/words \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"scenarioId":"scenario-uuid","targetText":"紅色","sourceText":"Red","displayOrder":1}'

# Get Word
curl -X GET http://localhost:8080/api/v1/words/{id} \
  -H "Authorization: Bearer $TOKEN"

# Update Word (add media URLs)
curl -X PUT http://localhost:8080/api/v1/words/{id} \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"imageUrl":"/uploads/images/red.jpg","audioUrl":"/uploads/audio/red.mp3"}'

# Delete Word
curl -X DELETE http://localhost:8080/api/v1/words/{id} \
  -H "Authorization: Bearer $TOKEN"
```

## Response Format

### Success Response
```json
{
  "success": true,
  "data": { ... }
}
```

### Error Response
```json
{
  "error": "Error message here"
}
```

## HTTP Status Codes

- `200 OK` - Successful GET/PUT request
- `201 Created` - Successful POST request
- `204 No Content` - Successful DELETE request
- `400 Bad Request` - Invalid input data
- `401 Unauthorized` - Missing or invalid token
- `403 Forbidden` - User doesn't have permission
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## Database Schema Overview

```
journeys
├── id (UUID, PK)
├── title
├── description
├── sourceLanguage
├── targetLanguage
├── status (draft/published/archived)
├── createdBy (FK -> users.id)
└── timestamps

scenarios
├── id (UUID, PK)
├── journeyId (FK -> journeys.id, CASCADE DELETE)
├── title
├── description
├── displayOrder
└── timestamps

words
├── id (UUID, PK)
├── scenarioId (FK -> scenarios.id, CASCADE DELETE)
├── targetText
├── sourceText
├── displayOrder
├── imageUrl
├── audioUrl
├── generationMethod
└── timestamps
```

## Common Workflows

### Create Complete Journey
1. Create journey → Get journey ID
2. Create scenario(s) with journey ID → Get scenario ID
3. Create word(s) with scenario ID
4. Update words with media URLs (Sprint 3)
5. Update journey status to "published"

### Update Content
1. GET journey to view current state
2. PUT journey/scenario/word to update
3. Changes are immediate (no separate save action)

### Delete Content
- Deleting journey → Deletes all scenarios and words
- Deleting scenario → Deletes all words
- Deleting word → Only deletes that word

## Development Tips

### Testing with jq (pretty JSON)
```bash
curl http://localhost:8080/api/v1/journeys \
  -H "Authorization: Bearer $TOKEN" | jq '.'
```

### Check Server Health
```bash
curl http://localhost:8080/health
# Response: {"status":"healthy","version":"1.0.0"}
```

### View Database
```bash
sqlite3 backend/learng.db
sqlite> .tables
sqlite> SELECT * FROM journeys;
sqlite> .quit
```

### Reset Database
```bash
rm backend/learng.db
# Restart server to auto-create new database
```

## Files to Know

### Backend Structure
```
backend/
├── cmd/api/main.go           # Entry point, route setup
├── internal/
│   ├── models/               # Database entities
│   ├── repository/           # Database access
│   ├── services/             # Business logic
│   ├── handlers/             # HTTP handlers
│   ├── middleware/           # Auth, logging
│   └── utils/                # Helpers
├── test-sprint2.sh           # Automated tests
└── SPRINT2_SUMMARY.md        # Detailed documentation
```

## Troubleshooting

### Port 8080 already in use
```bash
lsof -ti:8080 | xargs kill -9
```

### Database locked
```bash
# Stop server, restart
rm backend/learng.db  # If corrupted
```

### Token expired
```bash
# Re-login to get new token
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Test1234"}'
```

## Next Sprint Preview

Sprint 3 will add:
- Media upload endpoints (image, audio)
- File validation and storage
- Frontend components for media recording/upload

---

**Questions?** Check `SPRINT2_SUMMARY.md` for detailed documentation.
