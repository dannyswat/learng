# Frontend Authentication Testing Guide

## 🎯 Testing the Complete Authentication Flow

**Date**: 2025-10-07  
**Backend**: http://localhost:8080  
**Frontend**: http://localhost:5173

## ✅ What's Implemented

### Backend API (Go + Echo)
- ✅ POST `/api/v1/auth/register` - User registration
- ✅ POST `/api/v1/auth/login` - User login
- ✅ GET `/api/v1/auth/me` - Get current user (protected)
- ✅ JWT authentication middleware
- ✅ Password hashing with bcrypt
- ✅ Input validation

### Frontend (React + TypeScript)
- ✅ Login Page (`/login`)
- ✅ Register Page (`/register`)
- ✅ Admin Dashboard (`/admin`)
- ✅ Learner Dashboard (`/learn`)
- ✅ Protected routes with role-based access
- ✅ Auth context with JWT token management
- ✅ Axios interceptors for automatic token injection
- ✅ Auto-redirect on 401 Unauthorized

## 🧪 Manual Testing Steps

### Test 1: User Registration (Admin)

1. **Open** http://localhost:5173
2. **You should be redirected to** `/login` (not authenticated)
3. **Click** "create a new account"
4. **Fill in the form**:
   - Display Name: `Admin User`
   - Email: `admin@learng.com`
   - Password: `admin123`
   - Role: `Admin`
5. **Click** "Create account"
6. **Expected Result**:
   - Success: Redirected to `/admin` (Admin Dashboard)
   - See "Admin Dashboard" with your name
   - See 6 feature cards (Journeys, Scenarios, Words, etc.)
   - See account information section

### Test 2: User Registration (Learner)

1. **Logout** from admin dashboard
2. **Go to** `/register`
3. **Fill in the form**:
   - Display Name: `Test Learner`
   - Email: `learner@learng.com`
   - Password: `learner123`
   - Role: `Learner`
4. **Click** "Create account"
5. **Expected Result**:
   - Success: Redirected to `/learn` (Learner Dashboard)
   - See "My Learning" page
   - See progress stats (0 words, 0 mastered, 0 learning)
   - See "No journeys yet" message

### Test 3: Login with Existing User

1. **Logout** from learner dashboard
2. **Go to** `/login`
3. **Fill in the form**:
   - Email: `admin@learng.com`
   - Password: `admin123`
4. **Click** "Sign in"
5. **Expected Result**:
   - Success: Redirected to `/admin`
   - See admin dashboard with your information

### Test 4: Invalid Login

1. **Go to** `/login`
2. **Fill in wrong credentials**:
   - Email: `admin@learng.com`
   - Password: `wrongpassword`
3. **Click** "Sign in"
4. **Expected Result**:
   - Error message: "invalid email or password"
   - Stay on login page

### Test 5: Registration Validation

**Invalid Email**:
1. Go to `/register`
2. Enter email: `notanemail`
3. Submit form
4. **Expected**: Error "invalid email format"

**Weak Password**:
1. Go to `/register`
2. Enter password: `weak`
3. Submit form
4. **Expected**: Error "Password must be at least 8 characters long"

**Duplicate Email**:
1. Go to `/register`
2. Enter email: `admin@learng.com` (already exists)
3. Submit form
4. **Expected**: Error "user with this email already exists"

### Test 6: Protected Routes

**Without Token**:
1. **Open incognito/private window**
2. **Go directly to** http://localhost:5173/admin
3. **Expected**: Automatically redirected to `/login`

**With Token**:
1. **Login as admin**
2. **Go to** `/admin`
3. **Expected**: See admin dashboard

**Role Mismatch**:
1. **Login as learner**
2. **Try to access** `/admin` (URL bar)
3. **Expected**: Redirected to `/` then to `/learn`

### Test 7: Token Persistence

1. **Login as admin**
2. **Refresh the page** (F5)
3. **Expected**: Still logged in, stay on admin dashboard
4. **Close browser tab**
5. **Open new tab** to http://localhost:5173
6. **Expected**: Still logged in, redirected to `/admin`

### Test 8: Logout

1. **Login as any user**
2. **Click** "Logout" button
3. **Expected**:
   - Redirected to `/login`
   - Token removed from localStorage
4. **Try to access** `/admin` or `/learn`
5. **Expected**: Redirected to `/login`

## 🔍 Browser DevTools Inspection

### localStorage Check

Open Browser DevTools > Application > Local Storage > http://localhost:5173

**After Successful Login**:
```
authToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
user: {"id":"...","email":"admin@learng.com","role":"admin","displayName":"Admin User",...}
```

**After Logout**:
```
(Both items should be removed)
```

### Network Requests

Open DevTools > Network > XHR

**Registration Request**:
```
POST http://localhost:8080/api/v1/auth/register
Request Payload:
{
  "email": "admin@learng.com",
  "password": "admin123",
  "displayName": "Admin User",
  "role": "admin"
}

Response (201):
{
  "user": {...},
  "token": "eyJ..."
}
```

**Login Request**:
```
POST http://localhost:8080/api/v1/auth/login
Request Payload:
{
  "email": "admin@learng.com",
  "password": "admin123"
}

Response (200):
{
  "user": {...},
  "token": "eyJ..."
}
```

**Protected Request**:
```
GET http://localhost:8080/api/v1/auth/me
Request Headers:
Authorization: Bearer eyJ...

Response (200):
{
  "id": "...",
  "email": "admin@learng.com",
  ...
}
```

## 🎨 UI/UX Features to Verify

### Login Page
- ✅ Centered layout with logo area
- ✅ Email and password fields
- ✅ "Sign in" button with loading state
- ✅ Link to registration page
- ✅ Error messages displayed in red box

### Register Page
- ✅ Display name, email, password fields
- ✅ Role dropdown (Admin/Learner)
- ✅ Password helper text
- ✅ Role description text
- ✅ "Create account" button with loading state
- ✅ Link to login page

### Admin Dashboard
- ✅ Header with "Admin Dashboard" title
- ✅ Welcome message with user name
- ✅ Logout button in header
- ✅ 6 feature cards with icons
- ✅ "Coming Soon" disabled buttons
- ✅ Account information section
- ✅ User ID, email, role displayed

### Learner Dashboard
- ✅ Header with "My Learning" title
- ✅ Welcome message
- ✅ 3 progress stat cards (words, mastered, learning)
- ✅ "My Journeys" empty state
- ✅ "Recent Activity" empty state
- ✅ Account information section

## 🐛 Known Issues / Limitations

1. **No email verification** - Users can register with any email
2. **No password reset** - If user forgets password, admin must manually reset
3. **No refresh token** - Tokens expire after 24 hours, user must re-login
4. **No rate limiting** - Registration and login can be brute-forced
5. **Placeholder content** - All feature buttons show "Coming Soon"

## 🔄 CORS & Proxy

The frontend uses Vite proxy to avoid CORS issues in development:

**vite.config.ts**:
```typescript
proxy: {
  '/api': 'http://localhost:8080',
  '/uploads': 'http://localhost:8080',
}
```

Requests to `/api/*` are automatically forwarded to the backend.

## 🎯 Next Steps After Validation

1. **Frontend**:
   - [ ] Add journey management UI
   - [ ] Add scenario/word editors
   - [ ] Add media upload component
   - [ ] Add quiz interface

2. **Backend**:
   - [ ] Journey CRUD endpoints
   - [ ] Scenario CRUD endpoints
   - [ ] Word CRUD endpoints
   - [ ] Media upload handlers

3. **Testing**:
   - [ ] Add Vitest unit tests for components
   - [ ] Add E2E tests with Playwright
   - [ ] Add API integration tests

## 📸 Expected Screenshots

### Login Page
- Clean, centered form
- Blue "Sign in" button
- Link to register

### Register Page
- 4 form fields + dropdown
- Blue "Create account" button
- Helper text under password

### Admin Dashboard
- Grid of 6 feature cards
- Each card has icon + title + description
- Account info at bottom

### Learner Dashboard
- 3 stat cards at top
- Empty state graphics
- Friendly messages

## ✅ Success Criteria

- [x] User can register as admin
- [x] User can register as learner
- [x] User can login with credentials
- [x] Invalid credentials rejected
- [x] JWT token stored in localStorage
- [x] Token sent with protected requests
- [x] Admin sees admin dashboard
- [x] Learner sees learner dashboard
- [x] User can logout
- [x] Protected routes redirect to login
- [x] Page refresh preserves login state

---

**Test Completed**: _____________  
**Tester**: _____________  
**All Tests Passing**: ☐ Yes ☐ No  
**Issues Found**: _____________
