# Authentication Implementation Summary

## ✅ What Was Built

### 1. User Repository (`internal/repository/user.repo.go`)
- `Create(user)` - Create new user
- `GetByID(id)` - Get user by ID
- `GetByEmail(email)` - Get user by email
- `Update(user)` - Update user
- `Delete(id)` - Soft delete user
- `Exists(email)` - Check if email exists

### 2. Auth Service (`internal/services/auth.service.go`)
- `Register(req)` - Register new user with validation
- `Login(req)` - Authenticate user and return JWT token
- `GetUserByID(id)` - Get user details (without password hash)

**Features**:
- Email validation
- Password strength validation (min 8 chars, must contain number)
- Role validation (admin/learner)
- Password hashing with bcrypt
- JWT token generation (24-hour expiry)
- Duplicate email prevention

### 3. Auth Handler (`internal/handlers/auth.go`)
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `GET /api/v1/auth/me` - Get current user (protected)

### 4. Auth Middleware (`internal/middleware/auth.go`)
- `AuthMiddleware` - JWT token validation
- `RequireRole` - Role-based access control
- Context helpers: `GetUserID()`, `GetUserRole()`

### 5. Updated Utils
- `GenerateToken()` - Create JWT with custom duration
- `ValidateToken()` - Parse and validate JWT
- `GetUserID()` - Extract user ID from context
- `GetUserRole()` - Extract user role from context
- Response helpers updated to return map[string]interface{}

## 🔧 API Endpoints

### Public Endpoints (No Auth Required)

#### Register User
```bash
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "displayName": "John Doe",
  "role": "admin"  # or "learner"
}

Response (201):
{
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "displayName": "John Doe",
    "role": "admin",
    "createdAt": "2025-10-05T..."
  },
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

#### Login
```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response (200):
{
  "user": { ... },
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

### Protected Endpoints (Auth Required)

#### Get Current User
```bash
GET /api/v1/auth/me
Authorization: Bearer <token>

Response (200):
{
  "id": "uuid",
  "email": "user@example.com",
  "displayName": "John Doe",
  "role": "admin",
  "createdAt": "2025-10-05T..."
}
```

## 🧪 Testing

### Manual Testing with curl

1. **Start the server**:
```bash
make run
# or
go run cmd/api/main.go
```

2. **Run the automated test script**:
```bash
./test-auth.sh
```

This script tests:
- ✅ User registration (admin and learner)
- ✅ Duplicate email prevention
- ✅ Login with correct credentials
- ✅ Login with wrong password (should fail)
- ✅ Protected endpoint access with token
- ✅ Protected endpoint without token (should fail)
- ✅ Invalid email format validation
- ✅ Weak password validation

### Expected Test Results

**Successful Registration**:
```json
{
  "user": {
    "id": "...",
    "email": "admin@learng.com",
    "displayName": "Admin User",
    "role": "admin"
  },
  "token": "eyJ..."
}
```

**Failed Registration (duplicate email)**:
```json
{
  "error": "user with this email already exists"
}
```

**Successful Login**:
```json
{
  "user": { ... },
  "token": "eyJ..."
}
```

**Failed Login (wrong password)**:
```json
{
  "error": "invalid email or password"
}
```

**Protected Endpoint (no token)**:
```json
{
  "error": "Missing authorization header"
}
```

## 🔒 Security Features

1. **Password Hashing**: bcrypt with default cost factor (10)
2. **JWT Tokens**: 
   - 24-hour expiration
   - HS256 signing algorithm
   - Contains: userID, role (email excluded for security)
3. **Input Validation**:
   - Email format validation
   - Password minimum 8 characters
   - Password must contain at least 1 number
   - Role must be "admin" or "learner"
4. **Error Messages**: Generic messages to prevent information leakage

## 📁 File Structure

```
backend/
├── cmd/api/main.go                    # Updated with auth routes
├── internal/
│   ├── handlers/
│   │   └── auth.go                    # ✅ NEW: Auth handlers
│   ├── services/
│   │   └── auth.service.go            # ✅ NEW: Auth business logic
│   ├── repository/
│   │   └── user.repo.go               # ✅ NEW: User data access
│   ├── middleware/
│   │   └── auth.go                    # Already existed, unchanged
│   └── utils/
│       ├── jwt.go                     # Updated: GenerateToken signature
│       ├── response.go                # Updated: Return maps instead of errors
│       └── validation.go              # Already existed, unchanged
└── test-auth.sh                       # ✅ NEW: Automated test script
```

## 🚀 Next Steps

### Immediate
- [x] Authentication endpoints working
- [x] JWT middleware protecting routes
- [x] User registration and login
- [x] Password hashing and validation

### Sprint 2 - Journey Management
- [ ] Journey repository
- [ ] Journey service
- [ ] Journey handlers (CRUD)
- [ ] Admin-only routes with RequireRole middleware

### Sprint 3 - Scenario & Word Management
- [ ] Scenario repository and handlers
- [ ] Word repository and handlers
- [ ] Nested resource handling

## 🐛 Known Issues / Notes

1. **Database**: Currently using SQLite in-memory for development. Data is lost on restart.
   - Solution: Set `DB_PATH` in `.env` to persist data (e.g., `./learng.db`)

2. **JWT Secret**: Default secret is "dev-secret-change-in-production"
   - **CRITICAL**: Change `JWT_SECRET` in `.env` before production deployment

3. **Token Expiry**: Tokens expire after 24 hours. No refresh token mechanism yet.
   - Future: Implement refresh token flow

4. **CORS**: Currently allows all origins in development
   - Production: Whitelist specific frontend domain

## 📝 Usage Examples

### Register and Login Flow

```bash
# 1. Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "admin123",
    "displayName": "Admin User",
    "role": "admin"
  }'

# Save the token from response

# 2. Access protected endpoint
TOKEN="<token-from-response>"
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8080/api/v1/auth/me

# 3. Login again later
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@learng.com",
    "password": "admin123"
  }'
```

## ✅ Verification Checklist

- [x] User can register with valid email and password
- [x] Duplicate email registration is prevented
- [x] User can login with correct credentials
- [x] Login fails with wrong password
- [x] JWT token is returned on successful auth
- [x] Protected endpoints require valid token
- [x] Invalid tokens are rejected
- [x] User ID and role are accessible in protected handlers
- [x] Password hash is never returned in responses
- [x] Email and password validation works

---

**Status**: ✅ Authentication Complete  
**Next**: Journey Management (Sprint 2)
