# learng – Phase 1 MVP Functional Specification

## Document Overview
**Version**: 1.0  
**Last Updated**: 2025-10-05  
**Status**: Draft for Review  
**Phase**: MVP (Manual Content Creation → AI Automation Later)

## 1. Technology Stack

### 1.1 Backend
- **Language**: Go 1.21+
- **Web Framework**: Echo (lightweight, production-ready)
- **ORM**: GORM (robust, feature-complete)
- **Database**: SQLite (MVP) → PostgreSQL (production migration path)
- **File Storage**: Local filesystem (MVP) → Cloud object storage (production)
- **Authentication**: JWT-based sessions
- **API Style**: RESTful JSON APIs

### 1.2 Frontend
- **Framework**: React 18+ with TypeScript
- **Build Tool**: Vite (fast dev server, optimized builds)
- **Routing**: React Router v6
- **State Management**: React Context + hooks (MVP) → Zustand if needed
- **Styling**: Tailwind CSS + shadcn/ui components
- **HTTP Client**: Axios or fetch API with custom hooks
- **Media Recording**: MediaRecorder API (audio), file upload inputs

### 1.3 Development & Deployment
- **Package Manager**: npm (frontend), Go modules (backend)
- **Version Control**: Git + GitHub
- **Local Development**: 
  - Frontend: Vite dev server with proxy to backend API (avoids CORS)
  - Backend: Go binary serving API endpoints only
  - Single domain: `http://localhost:5173` (Vite proxies `/api/*` to Go backend)
- **Deployment**: 
  - Docker container with Go binary serving both API and static frontend
  - Single-domain deployment (no CORS issues)
- **Environment Config**: `.env` files (never committed)

### 1.4 Future Migration Path
- SQLite → PostgreSQL (schema compatible via GORM)
- Local storage → Azure Blob Storage / S3
- Manual media → AI generation endpoints (plug-in architecture)

---

## 2. Phase 1 Delivery Strategy

### 2.1 MVP Feature Scope
**First Delivery (Manual Content Creation)**:
- Admin creates journeys, scenarios, and words
- Admin **manually** records audio (via browser)
- Admin **manually** uploads images
- Learner views cards with uploaded media
- Learner takes quizzes
- Basic progress tracking

**Second Delivery (AI Automation)**:
- Add AI generation endpoints (image, audio)
- Admin can choose: manual upload OR AI generation
- Queue-based async generation
- Safety moderation pipeline

### 2.2 Design Philosophy
- **API-First**: Backend exposes clean REST APIs; frontend is pure consumer
- **Separation of Concerns**: Media handling abstracted for easy AI swap-in
- **Progressive Enhancement**: Manual → AI is additive, not replacement
- **Admin Choice**: Even after AI, allow manual overrides for quality control

---

## 3. Data Model (SQLite Schema)

### 3.1 Core Entities

#### 3.1.1 Users Table
```sql
CREATE TABLE users (
    id TEXT PRIMARY KEY,              -- UUID
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,      -- bcrypt
    role TEXT NOT NULL,               -- 'admin' | 'learner'
    display_name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
```

#### 3.1.2 Journeys Table
```sql
CREATE TABLE journeys (
    id TEXT PRIMARY KEY,              -- UUID
    title TEXT NOT NULL,
    description TEXT,
    source_language TEXT NOT NULL,    -- ISO 639-1 code (e.g., 'en')
    target_language TEXT NOT NULL,    -- ISO 639-1 code (e.g., 'zh-HK')
    status TEXT NOT NULL DEFAULT 'draft', -- 'draft' | 'published' | 'archived'
    created_by TEXT NOT NULL,         -- FK to users.id
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX idx_journeys_status ON journeys(status);
CREATE INDEX idx_journeys_created_by ON journeys(created_by);
```

#### 3.1.3 Scenarios Table
```sql
CREATE TABLE scenarios (
    id TEXT PRIMARY KEY,              -- UUID
    journey_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    display_order INTEGER NOT NULL,   -- For sequencing within journey
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (journey_id) REFERENCES journeys(id) ON DELETE CASCADE
);
CREATE INDEX idx_scenarios_journey ON scenarios(journey_id);
CREATE INDEX idx_scenarios_order ON scenarios(journey_id, display_order);
```

#### 3.1.4 Words Table
```sql
CREATE TABLE words (
    id TEXT PRIMARY KEY,              -- UUID
    scenario_id TEXT NOT NULL,
    target_text TEXT NOT NULL,        -- Word/phrase in target language
    source_text TEXT,                 -- Optional translation
    display_order INTEGER NOT NULL,
    image_url TEXT,                   -- Path to uploaded/generated image
    audio_url TEXT,                   -- Path to uploaded/generated audio
    generation_method TEXT DEFAULT 'manual', -- 'manual' | 'ai_image' | 'ai_audio' | 'ai_both'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (scenario_id) REFERENCES scenarios(id) ON DELETE CASCADE
);
CREATE INDEX idx_words_scenario ON words(scenario_id);
CREATE INDEX idx_words_order ON words(scenario_id, display_order);
```

#### 3.1.5 Quizzes Table
```sql
CREATE TABLE quizzes (
    id TEXT PRIMARY KEY,              -- UUID
    scenario_id TEXT NOT NULL,
    title TEXT NOT NULL,
    pass_threshold REAL DEFAULT 70.0, -- Percentage required to pass
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (scenario_id) REFERENCES scenarios(id) ON DELETE CASCADE
);
CREATE INDEX idx_quizzes_scenario ON quizzes(scenario_id);
```

#### 3.1.6 Quiz Questions Table
```sql
CREATE TABLE quiz_questions (
    id TEXT PRIMARY KEY,              -- UUID
    quiz_id TEXT NOT NULL,
    word_id TEXT NOT NULL,            -- Word being tested
    question_type TEXT NOT NULL,      -- 'multiple_choice' | 'audio_match' | 'image_match'
    question_text TEXT,               -- Optional explicit question
    correct_answer TEXT NOT NULL,     -- JSON or simple text
    options TEXT,                     -- JSON array of choices (for MCQ)
    display_order INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE,
    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE
);
CREATE INDEX idx_quiz_questions_quiz ON quiz_questions(quiz_id);
```

#### 3.1.7 Learner Progress Table
```sql
CREATE TABLE learner_progress (
    id TEXT PRIMARY KEY,              -- UUID
    user_id TEXT NOT NULL,
    word_id TEXT NOT NULL,
    mastery_level TEXT DEFAULT 'new', -- 'new' | 'learning' | 'review' | 'mastered'
    view_count INTEGER DEFAULT 0,
    last_viewed_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE,
    UNIQUE(user_id, word_id)
);
CREATE INDEX idx_progress_user ON learner_progress(user_id);
CREATE INDEX idx_progress_word ON learner_progress(word_id);
CREATE INDEX idx_progress_mastery ON learner_progress(user_id, mastery_level);
```

#### 3.1.8 Quiz Attempts Table
```sql
CREATE TABLE quiz_attempts (
    id TEXT PRIMARY KEY,              -- UUID
    user_id TEXT NOT NULL,
    quiz_id TEXT NOT NULL,
    score REAL NOT NULL,              -- Percentage (0-100)
    total_questions INTEGER NOT NULL,
    correct_answers INTEGER NOT NULL,
    answers TEXT NOT NULL,            -- JSON array of {question_id, answer, is_correct}
    completed_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE
);
CREATE INDEX idx_attempts_user ON quiz_attempts(user_id);
CREATE INDEX idx_attempts_quiz ON quiz_attempts(quiz_id);
CREATE INDEX idx_attempts_completed ON quiz_attempts(completed_at);
```

### 3.2 Future Extensions (Post-MVP)
- `generation_jobs` table (for async AI processing)
- `audit_logs` table (for compliance tracking)
- `media_cache` table (for AI prompt → asset reuse)

---

## 4. API Specification

### 4.1 Authentication Endpoints

#### POST /api/v1/auth/register
**Request**:
```json
{
  "email": "admin@example.com",
  "password": "securepass123",
  "displayName": "Teacher Jane",
  "role": "admin"
}
```
**Response** (201):
```json
{
  "id": "uuid-here",
  "email": "admin@example.com",
  "displayName": "Teacher Jane",
  "role": "admin",
  "token": "jwt-token-here"
}
```

#### POST /api/v1/auth/login
**Request**:
```json
{
  "email": "admin@example.com",
  "password": "securepass123"
}
```
**Response** (200):
```json
{
  "user": {
    "id": "uuid",
    "email": "admin@example.com",
    "displayName": "Teacher Jane",
    "role": "admin"
  },
  "token": "jwt-token-here"
}
```

### 4.2 Journey Management (Admin)

#### GET /api/v1/journeys
**Query Params**: `status=published`, `page=1`, `limit=20`  
**Response** (200):
```json
{
  "journeys": [
    {
      "id": "uuid",
      "title": "At the Park",
      "description": "Learn outdoor vocabulary",
      "sourceLanguage": "en",
      "targetLanguage": "zh-HK",
      "status": "published",
      "createdBy": "uuid",
      "createdAt": "2025-10-01T10:00:00Z",
      "scenarioCount": 3,
      "wordCount": 45
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 20
}
```

#### POST /api/v1/journeys
**Request**:
```json
{
  "title": "At the Park",
  "description": "Learn outdoor vocabulary",
  "sourceLanguage": "en",
  "targetLanguage": "zh-HK"
}
```
**Response** (201):
```json
{
  "id": "uuid",
  "title": "At the Park",
  "description": "Learn outdoor vocabulary",
  "sourceLanguage": "en",
  "targetLanguage": "zh-HK",
  "status": "draft",
  "createdBy": "uuid",
  "createdAt": "2025-10-05T12:00:00Z"
}
```

#### GET /api/v1/journeys/:id
**Response** (200):
```json
{
  "id": "uuid",
  "title": "At the Park",
  "description": "Learn outdoor vocabulary",
  "sourceLanguage": "en",
  "targetLanguage": "zh-HK",
  "status": "draft",
  "createdBy": "uuid",
  "scenarios": [
    {
      "id": "scenario-uuid",
      "title": "Nature Words",
      "description": "Trees, flowers, grass",
      "displayOrder": 1,
      "wordCount": 15
    }
  ]
}
```

#### PUT /api/v1/journeys/:id
**Request**:
```json
{
  "title": "At the Park (Updated)",
  "status": "published"
}
```
**Response** (200): Updated journey object

#### DELETE /api/v1/journeys/:id
**Response** (204): No content

### 4.3 Scenario Management (Admin)

#### POST /api/v1/scenarios
**Request**:
```json
{
  "journeyId": "uuid",
  "title": "Colors",
  "description": "Basic color vocabulary",
  "displayOrder": 1
}
```
**Response** (201): Scenario object

#### GET /api/v1/scenarios/:id
**Response** (200):
```json
{
  "id": "uuid",
  "journeyId": "uuid",
  "title": "Colors",
  "description": "Basic color vocabulary",
  "displayOrder": 1,
  "words": [
    {
      "id": "word-uuid",
      "targetText": "紅色",
      "sourceText": "Red",
      "displayOrder": 1,
      "imageUrl": "/uploads/images/red.jpg",
      "audioUrl": "/uploads/audio/red.mp3",
      "generationMethod": "manual"
    }
  ]
}
```

#### PUT /api/v1/scenarios/:id
#### DELETE /api/v1/scenarios/:id

### 4.4 Word Management (Admin)

#### POST /api/v1/words
**Request**:
```json
{
  "scenarioId": "uuid",
  "targetText": "紅色",
  "sourceText": "Red",
  "displayOrder": 1
}
```
**Response** (201):
```json
{
  "id": "uuid",
  "scenarioId": "uuid",
  "targetText": "紅色",
  "sourceText": "Red",
  "displayOrder": 1,
  "imageUrl": null,
  "audioUrl": null,
  "generationMethod": "manual"
}
```

#### PUT /api/v1/words/:id
**Request** (partial update):
```json
{
  "imageUrl": "/uploads/images/red-updated.jpg"
}
```

#### DELETE /api/v1/words/:id

### 4.5 Media Upload (Admin)

#### POST /api/v1/media/upload/image
**Content-Type**: `multipart/form-data`  
**Form Fields**:
- `file`: Image file (JPEG/PNG/WebP, max 5MB)
- `wordId`: UUID of associated word (optional for orphaned uploads)

**Response** (201):
```json
{
  "url": "/uploads/images/abc123.jpg",
  "filename": "abc123.jpg",
  "size": 245678,
  "mimeType": "image/jpeg"
}
```

#### POST /api/v1/media/upload/audio
**Content-Type**: `multipart/form-data`  
**Form Fields**:
- `file`: Audio file (MP3/WAV/WebM, max 2MB)
- `wordId`: UUID of associated word

**Response** (201):
```json
{
  "url": "/uploads/audio/def456.mp3",
  "filename": "def456.mp3",
  "size": 98765,
  "mimeType": "audio/mpeg",
  "duration": 1.8
}
```

**Error Handling**:
- 400: File too large, invalid format
- 413: Payload too large
- 500: Storage failure

### 4.6 Learner Endpoints

#### GET /api/v1/learner/journeys
**Response** (200): List of published journeys with progress indicators
```json
{
  "journeys": [
    {
      "id": "uuid",
      "title": "At the Park",
      "description": "Learn outdoor vocabulary",
      "progress": {
        "totalWords": 45,
        "viewedWords": 12,
        "masteredWords": 3,
        "completionPercentage": 26.7
      }
    }
  ]
}
```

#### GET /api/v1/learner/scenarios/:id
**Response** (200): Scenario with words and progress
```json
{
  "id": "uuid",
  "title": "Colors",
  "words": [
    {
      "id": "uuid",
      "targetText": "紅色",
      "sourceText": "Red",
      "imageUrl": "/uploads/images/red.jpg",
      "audioUrl": "/uploads/audio/red.mp3",
      "masteryLevel": "learning"
    }
  ],
  "quiz": {
    "id": "quiz-uuid",
    "title": "Colors Quiz",
    "questionCount": 5
  }
}
```

#### POST /api/v1/learner/progress
**Request**:
```json
{
  "wordId": "uuid"
}
```
**Response** (200): Updates view count, adjusts mastery if needed
```json
{
  "wordId": "uuid",
  "masteryLevel": "learning",
  "viewCount": 3,
  "lastViewedAt": "2025-10-05T14:30:00Z"
}
```

### 4.7 Quiz Endpoints

#### GET /api/v1/quizzes/:id
**Response** (200):
```json
{
  "id": "uuid",
  "scenarioId": "uuid",
  "title": "Colors Quiz",
  "questions": [
    {
      "id": "question-uuid",
      "type": "multiple_choice",
      "questionText": "What is this color?",
      "imageUrl": "/uploads/images/red.jpg",
      "options": ["紅色", "藍色", "綠色", "黃色"]
    }
  ]
}
```

#### POST /api/v1/quizzes/:id/submit
**Request**:
```json
{
  "answers": [
    {
      "questionId": "uuid",
      "answer": "紅色"
    }
  ]
}
```
**Response** (200):
```json
{
  "attemptId": "uuid",
  "score": 80.0,
  "totalQuestions": 5,
  "correctAnswers": 4,
  "passed": true,
  "feedback": [
    {
      "questionId": "uuid",
      "correct": true,
      "correctAnswer": "紅色"
    }
  ]
}
```

---

## 5. Frontend Architecture

### 5.1 Application Structure
```
frontend/
├── src/
│   ├── components/
│   │   ├── admin/
│   │   │   ├── JourneyForm.tsx
│   │   │   ├── ScenarioForm.tsx
│   │   │   ├── WordForm.tsx
│   │   │   ├── MediaUploader.tsx
│   │   │   └── AudioRecorder.tsx
│   │   ├── learner/
│   │   │   ├── JourneyCard.tsx
│   │   │   ├── WordCard.tsx
│   │   │   ├── QuizView.tsx
│   │   │   └── ProgressBar.tsx
│   │   ├── shared/
│   │   │   ├── Navigation.tsx
│   │   │   ├── Button.tsx
│   │   │   └── Modal.tsx
│   ├── pages/
│   │   ├── admin/
│   │   │   ├── DashboardPage.tsx
│   │   │   ├── JourneyListPage.tsx
│   │   │   ├── JourneyEditPage.tsx
│   │   │   └── AnalyticsPage.tsx
│   │   ├── learner/
│   │   │   ├── JourneysPage.tsx
│   │   │   ├── ScenarioViewPage.tsx
│   │   │   └── QuizPage.tsx
│   │   ├── LoginPage.tsx
│   │   └── RegisterPage.tsx
│   ├── hooks/
│   │   ├── useAuth.ts
│   │   ├── useJourneys.ts
│   │   ├── useMediaUpload.ts
│   │   └── useAudioRecorder.ts
│   ├── services/
│   │   ├── api.ts          // Axios instance + interceptors
│   │   ├── auth.service.ts
│   │   ├── journey.service.ts
│   │   └── media.service.ts
│   ├── contexts/
│   │   └── AuthContext.tsx
│   ├── types/
│   │   ├── api.types.ts
│   │   └── models.types.ts
│   ├── utils/
│   │   ├── validation.ts
│   │   └── format.ts
│   ├── App.tsx
│   └── main.tsx
├── public/
│   └── assets/
├── index.html
├── package.json
├── tsconfig.json
└── vite.config.ts
```

### 5.2 Vite Configuration (Development Proxy)

**Configuration** (`frontend/vite.config.ts`):
```typescript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

export default defineConfig({
  plugins: [react()],
  
  server: {
    port: 5173,
    // Proxy API and uploads to backend
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/uploads': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
})
```

**Benefits of this approach**:
- ✅ **No CORS configuration needed** - frontend and backend appear to be on the same domain
- ✅ **Same URLs in dev and production** - `/api/v1/auth/login` works identically
- ✅ **Simplified development** - Just run both servers, no CORS headers or preflight complexity
- ✅ **Realistic testing** - Development setup mirrors production (single origin)

**API Service Configuration** (`frontend/src/services/api.ts`):
```typescript
import axios from 'axios';

// No need for different base URLs - proxy handles it
const api = axios.create({
  baseURL: '', // Empty = same origin (works in dev via proxy, production via same server)
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add auth token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('authToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
```

### 5.3 Key Components

#### 5.3.1 AudioRecorder Component
**Functionality**:
- Uses `MediaRecorder API` to capture audio from user's microphone
- Visual waveform feedback during recording (optional: use WaveSurfer.js)
- Play back recorded audio before saving
- Upload to backend via `/api/v1/media/upload/audio`

**UI Flow**:
1. Click "Record Audio" button
2. Browser prompts for microphone permission
3. Recording indicator (red dot, timer)
4. Click "Stop" to finish
5. Playback controls appear
6. Click "Use This" to upload, or "Re-record" to discard

**Tech Stack**:
```typescript
interface AudioRecorderProps {
  onUploadComplete: (audioUrl: string) => void;
  wordId?: string;
}

// Implementation uses:
// - navigator.mediaDevices.getUserMedia()
// - MediaRecorder API
// - Blob handling for file upload
```

#### 5.3.2 MediaUploader Component
**Functionality**:
- Drag-and-drop or click-to-browse file selection
- Image preview before upload
- Progress bar during upload
- Supports replacing existing media

**Validation**:
- Image: JPEG/PNG/WebP, max 5MB, min 400x400px
- Audio: MP3/WAV/WebM, max 2MB, min 0.5s duration

#### 5.3.3 WordCard Component (Learner)
**Layout**:
```
┌──────────────────────────────┐
│                              │
│      [Large Image]           │
│                              │
├──────────────────────────────┤
│  紅色 (large, centered)       │
│  Red (smaller, gray)         │
├──────────────────────────────┤
│  [🔊 Play Audio Button]      │
└──────────────────────────────┘
```

**Behavior**:
- Auto-play audio on card load (optional setting)
- Tap anywhere to replay audio
- Swipe left/right or arrow keys to navigate
- Track view event → POST to `/api/v1/learner/progress`

### 5.4 Routing
```typescript
// React Router v6 setup
<Route path="/" element={<LandingPage />} />
<Route path="/login" element={<LoginPage />} />
<Route path="/register" element={<RegisterPage />} />

{/* Admin Routes */}
<Route path="/admin" element={<AdminLayout />}>
  <Route index element={<DashboardPage />} />
  <Route path="journeys" element={<JourneyListPage />} />
  <Route path="journeys/new" element={<JourneyEditPage />} />
  <Route path="journeys/:id" element={<JourneyEditPage />} />
  <Route path="analytics" element={<AnalyticsPage />} />
</Route>

{/* Learner Routes */}
<Route path="/learn" element={<LearnerLayout />}>
  <Route index element={<JourneysPage />} />
  <Route path="scenarios/:id" element={<ScenarioViewPage />} />
  <Route path="quizzes/:id" element={<QuizPage />} />
</Route>
```

### 5.5 State Management
**MVP Approach**: React Context + hooks (no Redux)

**AuthContext**:
- Stores: `user`, `token`, `isAuthenticated`
- Methods: `login()`, `logout()`, `register()`

**Local Component State**:
- Forms: `useState` for field values
- Data fetching: `useEffect` + `useState` for loading/error states
- Media uploads: Custom `useMediaUpload` hook with progress tracking

**Future**: If state complexity grows, migrate to Zustand or Redux Toolkit

---

## 6. Backend Architecture (Go)

### 6.1 Project Structure
```
backend/
├── cmd/
│   └── api/
│       └── main.go           // Entry point
├── internal/
│   ├── models/               // GORM models
│   │   ├── user.go
│   │   ├── journey.go
│   │   ├── scenario.go
│   │   ├── word.go
│   │   ├── quiz.go
│   │   └── progress.go
│   ├── handlers/             // HTTP handlers (controllers)
│   │   ├── auth.go
│   │   ├── journey.go
│   │   ├── scenario.go
│   │   ├── word.go
│   │   ├── media.go
│   │   ├── learner.go
│   │   └── quiz.go
│   ├── middleware/
│   │   ├── auth.go           // JWT validation
│   │   ├── cors.go
│   │   └── logger.go
│   ├── services/             // Business logic
│   │   ├── auth.service.go
│   │   ├── journey.service.go
│   │   ├── media.service.go
│   │   └── progress.service.go
│   ├── repository/           // Database access layer
│   │   ├── user.repo.go
│   │   ├── journey.repo.go
│   │   └── word.repo.go
│   ├── utils/
│   │   ├── jwt.go
│   │   ├── validation.go
│   │   └── response.go
│   └── config/
│       └── config.go         // Environment config
├── migrations/               // SQL migration files
│   ├── 001_create_users.sql
│   ├── 002_create_journeys.sql
│   └── ...
├── uploads/                  // Local file storage (MVP)
│   ├── images/
│   └── audio/
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

### 6.2 Key Packages
```go
// go.mod
module github.com/learng/backend

go 1.21

require (
    github.com/labstack/echo/v4 v4.11.3      // Web framework
    gorm.io/gorm v1.25.5                     // ORM
    gorm.io/driver/sqlite v1.5.4             // SQLite driver
    github.com/golang-jwt/jwt/v5 v5.1.0      // JWT
    github.com/google/uuid v1.4.0            // UUID generation
    golang.org/x/crypto v0.15.0              // bcrypt
    github.com/joho/godotenv v1.5.1          // .env loading
)
```

### 6.3 Configuration Management
```go
// internal/config/config.go
package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Port           string
    DatabasePath   string
    JWTSecret      string
    UploadDir      string
    StaticDir      string  // Frontend build directory (empty in dev)
    MaxImageSize   int64   // bytes
    MaxAudioSize   int64   // bytes
}

func Load() (*Config, error) {
    godotenv.Load()
    
    return &Config{
        Port:         getEnv("PORT", "8080"),
        DatabasePath: getEnv("DB_PATH", "./learng.db"),
        JWTSecret:    getEnv("JWT_SECRET", ""),
        UploadDir:    getEnv("UPLOAD_DIR", "./uploads"),
        StaticDir:    getEnv("STATIC_DIR", ""),  // Empty in dev, set in production
        MaxImageSize: 5 * 1024 * 1024,  // 5MB
        MaxAudioSize: 2 * 1024 * 1024,  // 2MB
    }, nil
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
```

### 6.4 GORM Models Example
```go
// internal/models/word.go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Word struct {
    ID               string         `gorm:"primaryKey" json:"id"`
    ScenarioID       string         `gorm:"index;not null" json:"scenarioId"`
    TargetText       string         `gorm:"not null" json:"targetText"`
    SourceText       string         `json:"sourceText"`
    DisplayOrder     int            `gorm:"not null" json:"displayOrder"`
    ImageURL         *string        `json:"imageUrl"`
    AudioURL         *string        `json:"audioUrl"`
    GenerationMethod string         `gorm:"default:manual" json:"generationMethod"`
    CreatedAt        time.Time      `json:"createdAt"`
    UpdatedAt        time.Time      `json:"updatedAt"`
    DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
    
    // Associations
    Scenario         Scenario       `gorm:"foreignKey:ScenarioID" json:"-"`
}

func (w *Word) BeforeCreate(tx *gorm.DB) error {
    if w.ID == "" {
        w.ID = uuid.New().String()
    }
    return nil
}
```

### 6.5 Media Upload Handler
```go
// internal/handlers/media.go
package handlers

import (
    "net/http"
    "path/filepath"
    "io"
    "os"
    "github.com/labstack/echo/v4"
    "github.com/google/uuid"
)

func (h *Handler) UploadImage(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file uploaded"})
    }
    
    // Validate file size
    if file.Size > h.config.MaxImageSize {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "File too large"})
    }
    
    // Validate MIME type
    if !isValidImageType(file.Header.Get("Content-Type")) {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file type"})
    }
    
    // Generate unique filename
    ext := filepath.Ext(file.Filename)
    filename := uuid.New().String() + ext
    savePath := filepath.Join(h.config.UploadDir, "images", filename)
    
    // Save file
    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open file"})
    }
    defer src.Close()
    
    dst, err := os.Create(savePath)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
    }
    defer dst.Close()
    
    if _, err = io.Copy(dst, src); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
    }
    
    // Return URL
    url := "/uploads/images/" + filename
    return c.JSON(http.StatusCreated, map[string]interface{}{
        "url":      url,
        "filename": filename,
        "size":     file.Size,
        "mimeType": file.Header.Get("Content-Type"),
    })
}

func isValidImageType(mimeType string) bool {
    validTypes := []string{"image/jpeg", "image/png", "image/webp"}
    for _, t := range validTypes {
        if t == mimeType {
            return true
        }
    }
    return false
}
```

### 6.6 Authentication Middleware
```go
// internal/middleware/auth.go
package middleware

import (
    "net/http"
    "strings"
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            authHeader := c.Request().Header.Get("Authorization")
            if authHeader == "" {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization header"})
            }
            
            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte(jwtSecret), nil
            })
            
            if err != nil || !token.Valid {
                return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
            }
            
            claims := token.Claims.(jwt.MapClaims)
            c.Set("userId", claims["userId"])
            c.Set("role", claims["role"])
            return next(c)
        }
    }
}

func RequireRole(role string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            userRole := c.Get("role")
            if userRole != role {
                return c.JSON(http.StatusForbidden, map[string]string{"error": "Insufficient permissions"})
            }
            return next(c)
        }
    }
}
```

---

## 7. User Workflows

### 7.1 Admin: Create Complete Journey (Manual)

**Step-by-Step Flow**:

1. **Login** → Navigate to Admin Dashboard
2. **Create Journey**:
   - Click "New Journey"
   - Fill form: Title, Description, Source Lang, Target Lang
   - Click "Create" → Status: `draft`
3. **Add Scenario**:
   - Inside journey editor, click "Add Scenario"
   - Fill: Title ("Colors"), Description, Order
   - Click "Save Scenario"
4. **Add Words**:
   - Click "Add Word" under "Colors" scenario
   - Enter: Target Text ("紅色"), Source Text ("Red"), Order (1)
   - Click "Save Word" → Word created but no media yet
5. **Upload Image**:
   - Click "Upload Image" button for word
   - Drag & drop or browse → `red.jpg`
   - Preview appears → Click "Confirm Upload"
   - Server returns `/uploads/images/abc123.jpg`
   - Word updated with `imageUrl`
6. **Record Audio**:
   - Click "Record Audio" button
   - Browser asks for mic permission → Allow
   - Click "Start Recording"
   - Speak word pronunciation
   - Click "Stop" → Playback controls appear
   - Listen to verify quality
   - Click "Use This Audio" → Uploads to server
   - Word updated with `audioUrl`
7. **Repeat for all words** (admin can bulk-add words, then batch upload media)
8. **Create Quiz**:
   - Click "Add Quiz" under scenario
   - Auto-generates 5 questions from scenario words
   - Admin can edit question type, add distractors
9. **Publish Journey**:
   - Click "Publish" → Status changes to `published`
   - Learners can now see it

**Time Estimate**: ~30-60 minutes for 15-word scenario (manual media)

### 7.2 Learner: Complete Scenario & Quiz

**Step-by-Step Flow**:

1. **Login** → Navigate to "My Journeys"
2. **Select Journey**: Click "At the Park" card
3. **Choose Scenario**: Click "Colors" (shows 15 words, 0% complete)
4. **View Cards**:
   - First card loads: large image of red apple, text "紅色 / Red"
   - Audio auto-plays (pronunciation)
   - Click "Next" or swipe → Second card (blue sky, "藍色 / Blue")
   - Progress updates in backend (view count increments)
5. **Navigate Through All Words**: Forward/back as needed
6. **Take Quiz**:
   - After viewing all words, "Take Quiz" button appears
   - Quiz loads: 5 multiple-choice questions
   - Question 1: Shows image of red apple, "Select the correct word"
     - Options: 紅色, 藍色, 綠色, 黃色
     - Learner taps "紅色"
   - Continue through all 5 questions
7. **Submit Quiz**:
   - Click "Submit" → Server grades answers
   - Results page: "You scored 80% (4/5 correct)"
   - Shows which question was wrong + correct answer
   - Mastery levels update (words answered correctly → `mastered`)
8. **Return to Journeys**: Click "Back to Journeys"
   - "Colors" scenario now shows 100% complete
   - Quiz badge with score displayed

---

## 8. Non-Functional Requirements

### 8.1 Performance
- **API Response Time**: <200ms for CRUD operations (local SQLite)
- **Media Upload**: Progress feedback every 100ms during upload
- **Card Navigation**: <100ms transition between cards
- **Quiz Grading**: <500ms for 10-question quiz

### 8.2 Security
- **Password Hashing**: bcrypt with cost factor 12
- **JWT Expiry**: 24 hours (refresh token in future phase)
- **File Upload Validation**: MIME type + file extension + magic number check
- **SQL Injection Protection**: GORM parameterized queries
- **CORS**: Whitelist frontend origin only
- **Rate Limiting**: (Future) 100 requests/minute per IP

### 8.3 Accessibility
- **WCAG 2.1 Level AA** compliance target:
  - Alt text for all images (required field for admin)
  - Keyboard navigation (tab order, arrow keys)
  - Focus indicators on interactive elements
  - Minimum contrast ratio 4.5:1
  - Screen reader compatible labels

### 8.4 Browser Support
- **Desktop**: Chrome 100+, Firefox 100+, Safari 15+, Edge 100+
- **Mobile**: iOS Safari 15+, Chrome Mobile 100+
- **Audio Recording**: Requires browser with MediaRecorder API support

### 8.5 Data Validation
**Backend (Go)**:
- Email format, password strength (min 8 chars, 1 number)
- Language codes (ISO 639-1 validation)
- File size limits enforced before save
- UUID format validation for all IDs

**Frontend (TypeScript)**:
- Form validation with immediate feedback
- Zod or Yup schemas for type-safe validation
- Disable submit button until valid

---

## 9. Migration Path to AI (Phase 2)

### 9.1 Architecture Preparation (MVP)
**Design decisions that enable future AI integration**:

1. **Abstract Media Source**:
   - Current: `generationMethod = 'manual'`
   - Future: `'ai_image'`, `'ai_audio'`, `'ai_both'`
   - Backend already tracks source method

2. **Endpoint Design**:
   - Keep manual upload endpoints: `/api/v1/media/upload/*`
   - Add parallel AI endpoints: `/api/v1/media/generate/*`
   - Admin UI offers choice: "Upload" or "Generate with AI"

3. **Data Model Extensions** (future):
   ```sql
   -- New table for AI generation jobs
   CREATE TABLE generation_jobs (
       id TEXT PRIMARY KEY,
       word_id TEXT NOT NULL,
       job_type TEXT NOT NULL,  -- 'image' | 'audio'
       status TEXT NOT NULL,     -- 'pending' | 'processing' | 'completed' | 'failed'
       provider TEXT,            -- 'azure_openai' | 'elevenlabs' | etc.
       prompt TEXT,              -- Generated prompt for audit
       result_url TEXT,
       error_message TEXT,
       created_at DATETIME,
       completed_at DATETIME,
       FOREIGN KEY (word_id) REFERENCES words(id)
   );
   ```

4. **Service Layer Abstraction**:
   ```go
   // Future interface (not implemented in MVP)
   type MediaGenerator interface {
       GenerateImage(ctx context.Context, req ImageRequest) (string, error)
       GenerateAudio(ctx context.Context, req AudioRequest) (string, error)
   }
   
   // Implementations: AzureImageGenerator, StabilityImageGenerator, etc.
   ```

### 9.2 Admin UI Flow (Post-AI)
**Hybrid workflow**:

1. Admin creates word with text
2. Two buttons appear:
   - **"Upload Media Manually"** → Opens file picker (current flow)
   - **"Generate with AI"** → Calls `/api/v1/media/generate`
3. If "Generate with AI":
   - Backend creates `generation_jobs` record
   - Returns job ID immediately
   - Frontend polls job status (or WebSocket in advanced version)
   - When complete, updates word with `imageUrl` and `audioUrl`
4. Admin can always regenerate or replace with manual upload

**No code changes to learner frontend** → They just see `imageUrl` and `audioUrl` regardless of source

### 9.3 Cost Tracking (Future)
```go
// Add to generation_jobs table
cost_estimate REAL,  -- USD
tokens_used INTEGER, -- For LLM calls
characters_used INTEGER -- For TTS calls

// Monthly aggregation query
SELECT 
    SUM(cost_estimate) as total_cost,
    COUNT(*) as total_jobs,
    AVG(cost_estimate) as avg_cost_per_job
FROM generation_jobs
WHERE created_at >= date('now', '-30 days');
```

---

## 10. Development Milestones

### Sprint 1 (Week 1-2): Foundation
- [ ] Backend: Project setup, GORM models, database migrations
- [ ] Backend: Auth endpoints (register, login, JWT middleware)
- [ ] Frontend: Project setup, routing, login/register pages
- [ ] Frontend: AuthContext and protected routes

### Sprint 2 (Week 3-4): Admin Core
- [ ] Backend: Journey/Scenario/Word CRUD APIs
- [ ] Backend: Media upload endpoints (image, audio)
- [ ] Frontend: Admin dashboard, journey list/create
- [ ] Frontend: Scenario and word management UI

### Sprint 3 (Week 5-6): Media Handling
- [ ] Frontend: Image upload component (drag-drop)
- [ ] Frontend: Audio recorder component (MediaRecorder API)
- [ ] Frontend: Media preview and playback
- [ ] Backend: File validation and storage logic

### Sprint 4 (Week 7-8): Learner Experience
- [ ] Backend: Learner endpoints (journeys, progress tracking)
- [ ] Frontend: Learner journey browsing
- [ ] Frontend: WordCard component with image/audio
- [ ] Frontend: Card navigation (swipe, keyboard)
- [ ] Backend: Progress tracking logic

### Sprint 5 (Week 9-10): Quiz System
- [ ] Backend: Quiz CRUD, question generation
- [ ] Backend: Quiz submission and grading
- [ ] Frontend: Quiz UI (question display, answer selection)
- [ ] Frontend: Quiz results and feedback
- [ ] Backend: Mastery level calculation

### Sprint 6 (Week 11-12): Polish & Testing
- [ ] End-to-end testing (manual + automated)
- [ ] UI/UX refinements (responsive design, animations)
- [ ] Performance optimization (lazy loading, caching)
- [ ] Basic analytics dashboard (admin)
- [ ] Deployment setup (production build, Docker)

### Sprint 7+ (Post-MVP): AI Integration
- [ ] Backend: AI adapter interfaces and implementations
- [ ] Backend: Generation job queue (BullMQ or Asynq)
- [ ] Backend: Safety moderation pipeline
- [ ] Frontend: "Generate with AI" UI option
- [ ] Cost tracking and monitoring

---

## 11. Testing Strategy

### 11.1 Backend Testing (Go)
```go
// Example: Test word creation
func TestCreateWord(t *testing.T) {
    db := setupTestDB()
    repo := repository.NewWordRepository(db)
    
    word := &models.Word{
        ScenarioID:   "scenario-123",
        TargetText:   "紅色",
        SourceText:   "Red",
        DisplayOrder: 1,
    }
    
    err := repo.Create(word)
    assert.NoError(t, err)
    assert.NotEmpty(t, word.ID)
    
    // Verify retrieval
    fetched, err := repo.GetByID(word.ID)
    assert.NoError(t, err)
    assert.Equal(t, "紅色", fetched.TargetText)
}
```

**Coverage Goals**:
- Unit tests: All service layer functions
- Integration tests: API endpoints with in-memory SQLite
- Manual testing: Media upload/playback

### 11.2 Frontend Testing (React)
```typescript
// Example: Test WordCard component
import { render, screen, fireEvent } from '@testing-library/react';
import { WordCard } from './WordCard';

test('plays audio when button clicked', async () => {
  const mockAudio = { play: jest.fn() };
  global.Audio = jest.fn(() => mockAudio);
  
  render(<WordCard 
    targetText="紅色" 
    sourceText="Red" 
    audioUrl="/test.mp3" 
  />);
  
  const playButton = screen.getByRole('button', { name: /play/i });
  fireEvent.click(playButton);
  
  expect(mockAudio.play).toHaveBeenCalled();
});
```

**Coverage Goals**:
- Component tests: All major UI components
- Integration tests: User flows (login → create journey → view cards)
- E2E tests (Playwright): Critical paths only

### 11.3 Manual QA Checklist
- [ ] Admin can create full journey with 10+ words
- [ ] Image upload works (JPEG, PNG, WebP)
- [ ] Audio recording works in Chrome, Firefox, Safari
- [ ] Learner can navigate cards with keyboard arrows
- [ ] Quiz grading is correct (test edge cases: all correct, all wrong)
- [ ] Progress persists after logout/login
- [ ] Responsive design works on tablet (iPad size)

---

## 12. Deployment

### 12.1 Development Setup (Vite Proxy)

**Architecture (Development)**:
```
┌─────────────────────────────────────┐
│  Vite Dev Server (:5173)            │
│  - Serves React app with HMR       │
│  - Proxies /api/* → :8080          │
│  - Proxies /uploads/* → :8080      │
└─────────────────────────────────────┘
              ↓ (proxy)
┌─────────────────────────────────────┐
│  Go Backend (:8080)                 │
│  - API endpoints (/api/*)           │
│  - Static uploads (/uploads/*)      │
│  - SQLite DB                        │
└─────────────────────────────────────┘
```

**Benefits**:
- ✅ Single domain (`http://localhost:5173`) - no CORS issues
- ✅ Hot module replacement (HMR) for fast frontend development
- ✅ Backend changes require restart, but frontend stays live

**Vite Configuration** (`frontend/vite.config.ts`):
```typescript
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/uploads': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

**Development Workflow**:
```bash
# Terminal 1: Start backend
cd backend
go run cmd/api/main.go

# Terminal 2: Start frontend with proxy
cd frontend
npm run dev

# Access app at http://localhost:5173
# All /api and /uploads requests auto-proxy to backend
```

### 12.2 Production Deployment (Docker)

**Architecture (Production)**:
```
┌─────────────────────────────────────┐
│  Docker Container                   │
│  ┌───────────────────────────────┐  │
│  │  Go Binary (Echo Server)      │  │
│  │  - Port 8080                  │  │
│  │  - Serves /api/* (API)        │  │
│  │  - Serves /* (React build)    │  │
│  │  - Serves /uploads/* (media)  │  │
│  │  - SQLite DB file             │  │
│  └───────────────────────────────┘  │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│  Reverse Proxy (Nginx/Caddy)       │
│  - SSL termination                  │
│  - Optional CDN for static assets   │
└─────────────────────────────────────┘
```

**Benefits**:
- ✅ Single binary deployment (simplicity)
- ✅ No CORS configuration needed (same origin)
- ✅ Easy to scale horizontally
- ✅ Portable across cloud providers

**Dockerfile**:
```dockerfile
# Build stage for frontend
FROM node:20-alpine AS frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ ./
RUN npm run build

# Build stage for backend
FROM golang:1.21-alpine AS backend-build
WORKDIR /app/backend
COPY backend/go.* ./
RUN go mod download
COPY backend/ ./
RUN CGO_ENABLED=1 GOOS=linux go build -o learng-api ./cmd/api

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /app

# Copy backend binary
COPY --from=backend-build /app/backend/learng-api .

# Copy frontend build
COPY --from=frontend-build /app/frontend/dist ./static

# Create uploads directory
RUN mkdir -p ./uploads/images ./uploads/audio

# Environment
ENV PORT=8080
ENV STATIC_DIR=/app/static
ENV UPLOAD_DIR=/app/uploads
ENV DB_PATH=/app/data/learng.db

# Expose port
EXPOSE 8080

# Run
CMD ["./learng-api"]
```

**Go Backend Configuration** (Serving Static Files):
```go
// cmd/api/main.go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())
    
    // API routes
    api := e.Group("/api/v1")
    // ... register API handlers
    
    // Serve uploaded media files
    e.Static("/uploads", config.UploadDir)
    
    // Serve frontend static files (production only)
    if config.StaticDir != "" {
        e.Static("/assets", config.StaticDir+"/assets")
        e.File("/favicon.ico", config.StaticDir+"/favicon.ico")
        
        // SPA fallback: serve index.html for all non-API routes
        e.File("/*", config.StaticDir+"/index.html")
    }
    
    e.Logger.Fatal(e.Start(":" + config.Port))
}
```

**Docker Compose (for local testing)**:
```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./data:/app/data
      - ./uploads:/app/uploads
    environment:
      - PORT=8080
      - DB_PATH=/app/data/learng.db
      - JWT_SECRET=${JWT_SECRET}
      - UPLOAD_DIR=/app/uploads
      - STATIC_DIR=/app/static
```

**Deployment Steps**:

1. **Build Docker image**:
   ```bash
   docker build -t learng:latest .
   ```

2. **Run container**:
   ```bash
   docker run -d \
     --name learng \
     -p 8080:8080 \
     -v $(pwd)/data:/app/data \
     -v $(pwd)/uploads:/app/uploads \
     -e JWT_SECRET=your-secret-key \
     learng:latest
   ```

3. **With Nginx reverse proxy** (optional, for SSL):
   ```nginx
   server {
       listen 80;
       server_name learng.example.com;
       
       # Redirect to HTTPS
       return 301 https://$server_name$request_uri;
   }
   
   server {
       listen 443 ssl http2;
       server_name learng.example.com;
       
       ssl_certificate /etc/letsencrypt/live/learng.example.com/fullchain.pem;
       ssl_certificate_key /etc/letsencrypt/live/learng.example.com/privkey.pem;
       
       # Proxy all requests to Go app
       location / {
           proxy_pass http://localhost:8080;
           proxy_set_header Host $host;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
           proxy_set_header X-Forwarded-Proto $scheme;
       }
   }
   ```

4. **Or use Caddy** (auto-SSL, simpler):
   ```
   learng.example.com {
       reverse_proxy localhost:8080
   }
   ```

### 12.3 Environment Variables (.env)

**Development** (`.env` in backend directory):
```bash
PORT=8080
DB_PATH=./learng.db
JWT_SECRET=dev-secret-change-in-production
UPLOAD_DIR=./uploads
STATIC_DIR=  # Empty in dev (Vite serves frontend)

# Future (AI Phase)
AZURE_OPENAI_KEY=xxx
AZURE_OPENAI_ENDPOINT=https://xxx.openai.azure.com
ELEVENLABS_API_KEY=xxx
```

**Production** (Docker environment or `.env` file):
```bash
PORT=8080
DB_PATH=/app/data/learng.db
JWT_SECRET=your-production-secret-key-here
UPLOAD_DIR=/app/uploads
STATIC_DIR=/app/static

# Future (AI Phase)
AZURE_OPENAI_KEY=xxx
AZURE_OPENAI_ENDPOINT=https://xxx.openai.azure.com
ELEVENLABS_API_KEY=xxx
```

### 12.4 Database Backups (SQLite)

**Using Docker volumes**:
```bash
# Backup
docker exec learng sqlite3 /app/data/learng.db ".backup '/app/data/backup-$(date +%Y%m%d).db'"

# Copy backup out of container
docker cp learng:/app/data/backup-20251005.db ./backups/

# Restore
docker cp ./backups/backup-20251005.db learng:/app/data/learng.db
docker restart learng
```

**Automated daily backups** (cron on host):
```bash
#!/bin/bash
# /opt/learng/backup.sh

CONTAINER_NAME=learng
BACKUP_DIR=/opt/learng/backups
DATE=$(date +%Y%m%d)

docker exec $CONTAINER_NAME sqlite3 /app/data/learng.db ".backup '/app/data/backup-$DATE.db'"
docker cp $CONTAINER_NAME:/app/data/backup-$DATE.db $BACKUP_DIR/
docker exec $CONTAINER_NAME rm /app/data/backup-$DATE.db

# Keep only last 30 days
find $BACKUP_DIR -name "backup-*.db" -mtime +30 -delete
```

**Crontab**:
```bash
0 2 * * * /opt/learng/backup.sh
```

---

## 13. Success Metrics (MVP)

### 13.1 Admin Efficiency
- **Time to create 15-word scenario**: <30 minutes (including media recording/upload)
- **Media upload success rate**: >98% (accounting for network issues)
- **Admin satisfaction**: Survey score >4/5 for ease of use

### 13.2 Learner Engagement
- **Scenario completion rate**: >60% of started scenarios
- **Quiz attempt rate**: >70% of learners who complete scenario take quiz
- **Quiz pass rate**: >65% (indicates content difficulty appropriate)
- **Repeat visit rate**: >40% return within 7 days

### 13.3 Technical Performance
- **API latency (p95)**: <300ms
- **Media load time**: <2s for images, <1s for audio
- **Uptime**: >99% (target 99.5%)
- **Bug rate**: <3 critical bugs per sprint

### 13.4 Cost (Pre-AI)
- **Hosting**: <$50/month (VPS + CDN)
- **Storage**: ~100MB per 100-word journey (images + audio)
- **Bandwidth**: <5GB/month for 100 active learners

---

## 14. Open Questions & Decisions Needed

| Question | Options | Recommendation | Status |
|----------|---------|----------------|--------|
| Web framework (Go) | Gin vs Echo vs Chi | **Echo** ✅ (clean API, middleware-friendly) | ✅ Decided |
| Component library | shadcn/ui vs Material-UI vs Chakra | **shadcn/ui** (Tailwind native, customizable) | Before Sprint 2 |
| Audio format | MP3 vs WebM vs WAV | **MP3** (universal browser support) | Before Sprint 3 |
| Max words per scenario | 10, 15, or 20 | **15** (balances session length) | Before Sprint 4 |
| Quiz auto-generation | Random 5 from scenario words | **Yes**, with manual override option | Before Sprint 5 |
| Deployment platform | Docker on VPS vs Cloud Run vs Fly.io | **Docker on VPS** ✅ (cost control, flexibility) | ✅ Decided |

---

## 15. Related Documents

- **REQUIREMENT.md**: Business objectives and user personas
- **.github/copilot-instructions.md**: AI agent context and strategic guidance
- **Future Documents**:
  - `API_REFERENCE.md`: OpenAPI/Swagger spec
  - `DEPLOYMENT_GUIDE.md`: Step-by-step production deployment
  - `AI_INTEGRATION.md`: Phase 2 AI provider implementation guide
  - `USER_STORIES.md`: Sprint backlog and acceptance criteria

---

## 16. Revision History

| Version | Date       | Author | Changes |
|---------|------------|--------|---------|
| 1.0     | 2025-10-05 | Product Team | Initial functional spec for Phase 1 MVP with Go + React + SQLite stack |

---

**Next Steps**:
1. ✅ Stakeholder review and approval of this spec
2. ⏳ Initialize Go and React projects (Sprint 1 kickoff)
3. ⏳ Create detailed user stories for Sprint 1
4. ⏳ Setup development environment and CI/CD pipeline

**Document Owner**: Technical Lead  
**Approval Required From**: Product Lead, CTO  
**Target Approval Date**: 2025-10-10
