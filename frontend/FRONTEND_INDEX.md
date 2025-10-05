# learng Frontend - Project Index

> **Navigation guide for the React/TypeScript frontend codebase**

## 📋 Quick Reference

- **Purpose**: Web-based UI for admin content management and learner experience
- **Tech Stack**: React 18, TypeScript, Vite, Tailwind CSS, React Router
- **Dev Server**: `npm run dev` → http://localhost:5173
- **Build Output**: `/dist` directory (served by Go backend in production)

---

## 🗂️ Directory Structure

```
frontend/
├── src/
│   ├── components/          # React components
│   │   ├── admin/          # Admin-specific components (to be built)
│   │   ├── learner/        # Learner-specific components (to be built)
│   │   └── shared/         # Shared/reusable components
│   │       ├── Button.tsx  # ✅ Reusable button with variants
│   │       └── Input.tsx   # ✅ Form input with label/error handling
│   ├── pages/              # Page components
│   │   └── LoginPage.tsx   # ✅ Login form with email/password
│   ├── hooks/              # Custom React hooks
│   │   ├── useMediaUpload.ts    # ✅ Image/audio upload hook
│   │   └── useAudioRecorder.ts  # ✅ Browser audio recording hook
│   ├── services/           # API service layer
│   │   ├── api.ts                # ✅ Axios instance with interceptors
│   │   ├── auth.service.ts       # ✅ Login, register, logout
│   │   ├── journey.service.ts    # ✅ Journey CRUD operations
│   │   └── media.service.ts      # ✅ Image/audio upload
│   ├── contexts/           # React contexts
│   │   └── AuthContext.tsx # ✅ Authentication state management
│   ├── types/              # TypeScript definitions
│   │   └── api.types.ts    # ✅ Complete API type definitions
│   ├── App.tsx             # ✅ Main app with routing
│   ├── main.tsx            # ✅ Application entry point
│   └── index.css           # ✅ Global styles (Tailwind directives)
├── public/                 # Static assets (to be populated)
├── index.html              # ✅ HTML template
├── package.json            # ✅ Dependencies and scripts
├── vite.config.ts          # ✅ Vite config with API proxy
├── tailwind.config.js      # ✅ Tailwind theme customization
├── tsconfig.json           # ✅ TypeScript configuration
├── .eslintrc.cjs           # ✅ ESLint rules
├── .gitignore              # ✅ Git ignore patterns
├── .env.example            # ✅ Environment variable template
├── setup.sh                # ✅ Automated setup script
└── README.md               # ✅ Comprehensive documentation
```

---

## 🎯 Implementation Status

### ✅ Complete (Phase 0: Foundation)

**Configuration & Build**
- [x] `package.json` - Dependencies (React 18, TypeScript, Vite, Tailwind, React Router, Axios)
- [x] `vite.config.ts` - Dev server with proxy (`/api/*` and `/uploads/*` → backend)
- [x] `tsconfig.json` - TypeScript strict mode with path aliases (`@/*`)
- [x] `tailwind.config.js` - Custom color palette (primary blues)
- [x] `postcss.config.js` - Tailwind CSS processing
- [x] `.eslintrc.cjs` - ESLint with TypeScript and React rules

**Type Definitions** (`src/types/api.types.ts`)
- [x] User types: `User`, `AuthResponse`, `LoginRequest`, `RegisterRequest`
- [x] Journey types: `Journey`, `JourneyListResponse`, `CreateJourneyRequest`
- [x] Scenario types: `Scenario`, `CreateScenarioRequest`
- [x] Word types: `Word`, `CreateWordRequest`, `UpdateWordRequest`
- [x] Media types: `MediaUploadResponse`
- [x] Quiz types: `Quiz`, `QuizQuestion`, `QuizAttempt`, `SubmitQuizRequest`
- [x] Progress types: `LearnerProgress`, `JourneyProgress`, `LearnerJourney`

**API Services**
- [x] `api.ts` - Axios instance with auth token injection + 401 redirect
- [x] `auth.service.ts` - `login()`, `register()`, `logout()`, `getCurrentUser()`, `isAuthenticated()`
- [x] `journey.service.ts` - `getJourneys()`, `getJourneyById()`, `createJourney()`, `updateJourney()`, `deleteJourney()`, `publishJourney()`
- [x] `media.service.ts` - `uploadImage()`, `uploadAudio()`

**Contexts**
- [x] `AuthContext.tsx` - Global auth state with `user`, `isAuthenticated`, `login()`, `register()`, `logout()`

**Custom Hooks**
- [x] `useMediaUpload.ts` - `uploadImage()`, `uploadAudio()` with loading/error states
- [x] `useAudioRecorder.ts` - `startRecording()`, `stopRecording()`, `pauseRecording()`, `resumeRecording()` with MediaRecorder API

**Shared Components**
- [x] `Button.tsx` - Variants: primary/secondary/outline/danger, sizes: sm/md/lg, loading state
- [x] `Input.tsx` - Label, error message, helper text support

**Pages**
- [x] `LoginPage.tsx` - Email/password form with error handling

**App Structure**
- [x] `App.tsx` - Routing with protected routes, role-based redirects (admin→`/admin`, learner→`/learn`)
- [x] `main.tsx` - React 18 root mounting

**Styling**
- [x] `index.css` - Tailwind directives + custom utility classes (`.btn`, `.btn-primary`, `.input`, `.card`)

**Documentation**
- [x] `README.md` - Getting started, structure, API usage, styling guide, troubleshooting
- [x] `setup.sh` - Automated dependency installation and verification

### 🚧 Pending (Sprint Implementation)

**Sprint 1: Authentication UI**
- [ ] `RegisterPage.tsx` - User registration form
- [ ] Error toast/notification component
- [ ] Loading spinner component
- [ ] Form validation utilities

**Sprint 2: Admin Journey Management**
- [ ] `admin/DashboardPage.tsx` - Overview of journeys and stats
- [ ] `admin/JourneyListPage.tsx` - Table of all journeys with filters
- [ ] `admin/JourneyEditPage.tsx` - Create/edit journey form
- [ ] `admin/JourneyForm.tsx` - Reusable journey form component
- [ ] `admin/ScenarioForm.tsx` - Scenario creation within journey
- [ ] `admin/WordForm.tsx` - Word creation within scenario
- [ ] Journey service integration in components

**Sprint 3: Media Handling**
- [ ] `admin/MediaUploader.tsx` - Drag-drop image upload with preview
- [ ] `admin/AudioRecorder.tsx` - Audio recording UI with playback
- [ ] `admin/ImagePreview.tsx` - Image cropping/editing (optional)
- [ ] Media validation (MIME type, file size, dimensions)
- [ ] Upload progress indicators

**Sprint 4: Learner Experience**
- [ ] `learner/JourneysPage.tsx` - Browse published journeys
- [ ] `learner/JourneyCard.tsx` - Journey preview with progress
- [ ] `learner/ScenarioViewPage.tsx` - View scenario words as cards
- [ ] `learner/WordCard.tsx` - Card with image, text, audio playback
- [ ] `learner/ProgressBar.tsx` - Visual progress indicator
- [ ] Card navigation (swipe gestures, keyboard arrows)
- [ ] Progress tracking (view count, mastery update)

**Sprint 5: Quiz System**
- [ ] `learner/QuizPage.tsx` - Quiz taking interface
- [ ] `learner/QuizView.tsx` - Question display with answer options
- [ ] `learner/QuizResults.tsx` - Score and feedback after submission
- [ ] Quiz answer validation
- [ ] Quiz service integration

**Sprint 6: Polish & Features**
- [ ] `shared/Modal.tsx` - Reusable modal/dialog
- [ ] `shared/Navigation.tsx` - Header navigation with user menu
- [ ] `shared/Card.tsx` - Generic card container
- [ ] Responsive design improvements (tablet, mobile)
- [ ] Accessibility enhancements (keyboard nav, screen reader support)
- [ ] Loading skeletons
- [ ] Error boundary component
- [ ] 404 page

---

## 🔑 Key Files Explained

### Configuration Files

#### `vite.config.ts`
**Purpose**: Vite build tool configuration  
**Key Features**:
- Dev server on port 5173
- **Proxy configuration**: `/api/*` and `/uploads/*` forwarded to `http://localhost:8080`
  - **Why**: Avoids CORS issues in development (frontend and backend appear on same origin)
- Path alias: `@/` → `src/`
- Build output: `dist/` directory

**Critical for**: Development workflow (no CORS), same-origin API calls

#### `tailwind.config.js`
**Purpose**: Tailwind CSS theme customization  
**Key Features**:
- Custom color palette: `primary-50` through `primary-900` (blue theme)
- Font family: Inter as default sans-serif
- Content paths: Scans `index.html` and all `.tsx`/`.jsx` files for class names

**Critical for**: Consistent styling, custom brand colors

#### `tsconfig.json`
**Purpose**: TypeScript compiler configuration  
**Key Features**:
- Strict mode enabled (catches type errors early)
- Path alias: `@/*` maps to `src/*` (cleaner imports)
- JSX mode: `react-jsx` (React 18 new JSX transform)

**Critical for**: Type safety, IDE autocomplete, import aliases

---

### Core Application Files

#### `src/main.tsx`
**Purpose**: Application entry point  
**Flow**: Mounts React app to `#root` element in `index.html`  
**Dependencies**: React 18's `createRoot`, StrictMode wrapper

#### `src/App.tsx`
**Purpose**: Root component with routing logic  
**Key Features**:
- `AuthProvider` wraps entire app (global auth state)
- `ProtectedRoute` wrapper enforces authentication + role-based access
- `RootRedirect` redirects `/` to `/admin` or `/learn` based on user role
- Route structure:
  - `/login` - Public login page
  - `/register` - Public registration page
  - `/admin/*` - Admin-only routes
  - `/learn/*` - Learner-only routes

**Critical for**: Routing, authentication flow, role separation

#### `src/index.css`
**Purpose**: Global styles and Tailwind directives  
**Key Features**:
- Tailwind base, components, utilities layers
- Custom utility classes:
  - `.btn` - Base button styles
  - `.btn-primary`, `.btn-secondary`, `.btn-outline` - Button variants
  - `.input` - Form input styles
  - `.card` - Card container styles
- Global body styles (gray background)

**Critical for**: Consistent styling, reusable utility classes

---

### Type Definitions

#### `src/types/api.types.ts`
**Purpose**: Central repository for all TypeScript types matching backend API  
**Coverage**:
- Request/response types for all API endpoints
- Domain models (User, Journey, Scenario, Word, Quiz, Progress)
- Enums (user roles, mastery levels, quiz types, journey status)

**Why Important**:
- Type safety across entire app
- Autocomplete in IDE
- Catches API contract mismatches at compile time

**Example Usage**:
```tsx
import { Journey, CreateJourneyRequest } from '@/types/api.types';

const createJourney = async (data: CreateJourneyRequest): Promise<Journey> => {
  const response = await api.post<Journey>('/api/v1/journeys', data);
  return response.data;
};
```

---

### API Services

#### `src/services/api.ts`
**Purpose**: Axios HTTP client with global configuration  
**Key Features**:
- **Request interceptor**: Injects JWT token from `localStorage` into `Authorization` header
- **Response interceptor**: Detects 401 Unauthorized → clears auth → redirects to `/login`
- Base URL: Empty string (relies on Vite proxy in dev, same-origin in production)
- Timeout: 10 seconds

**Critical for**: Automatic auth token management, error handling

#### `src/services/auth.service.ts`
**Purpose**: Authentication-specific API calls  
**Methods**:
- `login(credentials)` - POST `/api/v1/auth/login` → stores token + user in `localStorage`
- `register(data)` - POST `/api/v1/auth/register` → stores token + user
- `logout()` - Clears `localStorage`
- `getCurrentUser()` - Retrieves user from `localStorage`
- `isAuthenticated()` - Checks if token exists

**Usage**:
```tsx
import { authService } from '@/services/auth.service';

await authService.login({ email, password });
// Token + user automatically stored
```

#### `src/services/journey.service.ts`
**Purpose**: Journey CRUD operations  
**Methods**:
- `getJourneys(params)` - GET with filters (status, page, limit)
- `getJourneyById(id)` - GET single journey with scenarios
- `createJourney(data)` - POST new journey
- `updateJourney(id, data)` - PUT partial update
- `deleteJourney(id)` - DELETE
- `publishJourney(id)` - Convenience method (updates status to 'published')

#### `src/services/media.service.ts`
**Purpose**: File upload handling  
**Methods**:
- `uploadImage(file, wordId?)` - POST multipart/form-data to `/api/v1/media/upload/image`
- `uploadAudio(file, wordId?)` - POST multipart/form-data to `/api/v1/media/upload/audio`

**Why Separate from Other Services**: Different content-type (`multipart/form-data` vs `application/json`)

---

### React Contexts

#### `src/contexts/AuthContext.tsx`
**Purpose**: Global authentication state management  
**State**:
- `user: User | null` - Current logged-in user
- `isAuthenticated: boolean` - Derived from user existence
- `isLoading: boolean` - True during initial load (checking `localStorage`)

**Methods**:
- `login(email, password)` - Calls auth service, updates state
- `register(...)` - Calls auth service, updates state
- `logout()` - Clears auth, resets state

**Usage**:
```tsx
import { useAuth } from '@/contexts/AuthContext';

const MyComponent = () => {
  const { user, logout } = useAuth();
  
  return (
    <div>
      <p>Welcome, {user?.displayName}</p>
      <button onClick={logout}>Logout</button>
    </div>
  );
};
```

**Critical for**: Avoiding prop drilling, centralized auth logic

---

### Custom Hooks

#### `src/hooks/useMediaUpload.ts`
**Purpose**: Encapsulates image/audio upload logic with state  
**State**:
- `isUploading: boolean`
- `progress: number` (0-100, currently placeholder)
- `error: string | null`
- `url: string | null` (result URL after upload)

**Methods**:
- `uploadImage(file, wordId?)` - Async upload, updates state
- `uploadAudio(file, wordId?)` - Async upload, updates state
- `reset()` - Clear state

**Usage**:
```tsx
const { uploadImage, isUploading, url, error } = useMediaUpload();

const handleFileChange = async (e) => {
  const file = e.target.files[0];
  await uploadImage(file, wordId);
  if (url) {
    console.log('Upload successful:', url);
  }
};
```

#### `src/hooks/useAudioRecorder.ts`
**Purpose**: Browser audio recording using MediaRecorder API  
**State**:
- `isRecording: boolean`
- `isPaused: boolean`
- `recordingTime: number` (seconds elapsed)
- `audioBlob: Blob | null` (recorded audio data)
- `audioUrl: string | null` (object URL for playback)
- `error: string | null`

**Methods**:
- `startRecording()` - Requests mic permission, starts recording
- `stopRecording()` - Stops recording, generates Blob + URL
- `pauseRecording()` - Pauses recording (can resume)
- `resumeRecording()` - Resumes paused recording
- `reset()` - Clear all state, revoke object URL

**Usage**:
```tsx
const { startRecording, stopRecording, audioBlob, audioUrl } = useAudioRecorder();

<button onClick={startRecording}>Start</button>
<button onClick={stopRecording}>Stop</button>
{audioUrl && <audio src={audioUrl} controls />}
```

**Browser Compatibility**: Requires MediaRecorder API (Chrome 49+, Firefox 25+, Safari 14.1+)

---

### Shared Components

#### `src/components/shared/Button.tsx`
**Props**:
- `variant?: 'primary' | 'secondary' | 'outline' | 'danger'` (default: 'primary')
- `size?: 'sm' | 'md' | 'lg'` (default: 'md')
- `isLoading?: boolean` - Shows spinner, disables button
- Extends `React.ButtonHTMLAttributes<HTMLButtonElement>`

**Usage**:
```tsx
<Button variant="primary" size="lg" onClick={handleSave}>
  Save Journey
</Button>

<Button variant="danger" isLoading={isDeleting} onClick={handleDelete}>
  Delete
</Button>
```

#### `src/components/shared/Input.tsx`
**Props**:
- `label?: string` - Displayed above input
- `error?: string` - Red error message below input
- `helperText?: string` - Gray helper text below input
- Extends `React.InputHTMLAttributes<HTMLInputElement>`

**Usage**:
```tsx
<Input
  label="Email address"
  type="email"
  value={email}
  onChange={(e) => setEmail(e.target.value)}
  error={emailError}
  helperText="We'll never share your email."
/>
```

---

### Pages

#### `src/pages/LoginPage.tsx`
**Purpose**: User login form  
**Features**:
- Email and password inputs
- Form validation (browser native `required`)
- Error display from API
- Loading state during submission
- Link to register page

**Flow**:
1. User enters credentials
2. Submit → calls `authService.login()`
3. On success → `AuthContext` updates → navigate to `/` → `RootRedirect` sends to `/admin` or `/learn`
4. On error → display error message

**Next Steps**: Add client-side validation, "remember me" checkbox, "forgot password" link

---

## 🔄 Data Flow Patterns

### Authentication Flow

```
User enters credentials
    ↓
LoginPage.handleSubmit()
    ↓
authService.login({ email, password })
    ↓
POST /api/v1/auth/login
    ↓
Backend returns { user, token }
    ↓
authService stores token + user in localStorage
    ↓
AuthContext.login() updates state
    ↓
useAuth() hook triggers re-render
    ↓
ProtectedRoute allows access
    ↓
RootRedirect sends to /admin or /learn
```

### API Request Flow

```
Component calls service method
    ↓
Service method calls api.get/post/put/delete
    ↓
Axios request interceptor adds JWT token
    ↓
Vite proxy forwards to backend (dev) OR same origin (prod)
    ↓
Backend processes request
    ↓
Response returns
    ↓
Axios response interceptor checks for 401
    ↓
If 401: clear auth, redirect to /login
    ↓
If success: return data to service
    ↓
Service returns typed data to component
    ↓
Component updates state, re-renders
```

### Media Upload Flow

```
User selects file
    ↓
useMediaUpload.uploadImage(file, wordId)
    ↓
mediaService.uploadImage(file, wordId)
    ↓
FormData created with file + wordId
    ↓
POST /api/v1/media/upload/image (multipart/form-data)
    ↓
Backend validates file (type, size)
    ↓
Backend saves to /uploads/images/
    ↓
Backend returns { url, filename, size, mimeType }
    ↓
useMediaUpload state updates: { url, isUploading: false }
    ↓
Component displays image preview or updates word
```

---

## 🎨 Styling Architecture

### Tailwind Utility-First Approach

**Example Component**:
```tsx
<div className="flex items-center justify-between p-4 bg-white rounded-lg shadow-sm border border-gray-200">
  <h2 className="text-lg font-semibold text-gray-900">Journey Title</h2>
  <Button variant="primary" size="sm">Edit</Button>
</div>
```

**Breakdown**:
- `flex items-center justify-between` - Flexbox layout
- `p-4` - Padding 1rem (16px) all sides
- `bg-white` - White background
- `rounded-lg` - Large border radius
- `shadow-sm` - Small drop shadow
- `border border-gray-200` - 1px gray border
- `text-lg font-semibold text-gray-900` - Large, bold, dark gray text

### Custom Utility Classes (in `index.css`)

**Why**: Commonly used patterns abstracted for consistency

```css
.btn {
  /* Base button styles: flex, rounded, transitions, focus ring */
}

.btn-primary {
  /* Blue background, white text, hover effects */
}

.input {
  /* Standard input styles: border, padding, focus ring */
}

.card {
  /* White background, rounded, border, shadow */
}
```

**Usage**: Apply directly or compose with Tailwind classes
```tsx
<button className="btn btn-primary">Save</button>
<input className="input" />
<div className="card">Content</div>
```

### Responsive Design

Tailwind breakpoints (mobile-first):
- `sm:` - 640px and up
- `md:` - 768px and up
- `lg:` - 1024px and up
- `xl:` - 1280px and up

**Example**:
```tsx
<div className="px-4 sm:px-6 lg:px-8">
  {/* Horizontal padding increases with screen size */}
</div>

<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
  {/* 1 column mobile, 2 tablet, 3 desktop */}
</div>
```

---

## 🧪 Development Workflow

### Starting Development

```bash
# Terminal 1: Backend
cd backend
make run

# Terminal 2: Frontend
cd frontend
npm run dev
```

**Access**: http://localhost:5173

### Making Changes

1. **Edit component file** (e.g., `src/components/admin/JourneyForm.tsx`)
2. **Save** → Vite hot-reloads instantly
3. **Check console** for errors
4. **Test in browser**

### Type Checking

```bash
npm run type-check
```

**Why**: Catches type errors without building (faster than full build)

### Building for Production

```bash
npm run build
```

**Output**: `/dist` directory with optimized files

**Deployment**: Copy `/dist/*` to backend's `/static` directory

---

## 🐛 Common Issues & Solutions

### Import Errors (Red Squiggles)

**Symptom**: `Cannot find module 'react'` errors in VSCode  
**Cause**: Dependencies not installed yet  
**Solution**: Run `npm install`

### CORS Errors

**Symptom**: Network errors with CORS policy violations  
**Cause**: Backend not running OR proxy misconfigured  
**Solution**:
1. Verify backend is on http://localhost:8080
2. Check `vite.config.ts` proxy settings
3. Restart both servers

### Type Errors in Components

**Symptom**: `Parameter implicitly has 'any' type`  
**Cause**: Missing type annotations  
**Solution**: Add explicit types
```tsx
// Before
const handleChange = (e) => setEmail(e.target.value);

// After
const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value);
```

### 401 Unauthorized Errors

**Symptom**: API calls fail with 401 status  
**Cause**: Token expired or invalid  
**Solution**: Axios interceptor auto-redirects to `/login` - just log in again

---

## 📚 Next Implementation Steps

### Sprint 1 Priorities

1. **Complete RegisterPage.tsx**
   - Copy structure from `LoginPage.tsx`
   - Add `displayName` field
   - Add `role` selection (admin/learner)
   - Wire up `authService.register()`

2. **Add Error Handling**
   - Create `Toast.tsx` component for notifications
   - Add to `App.tsx` layout
   - Show success/error messages

3. **Loading States**
   - Create `Spinner.tsx` component
   - Use in protected routes while `isLoading`
   - Use in buttons during async operations

4. **Form Validation**
   - Create `src/utils/validation.ts`
   - Email format validation
   - Password strength validation (min 8 chars, 1 number)
   - Display errors in real-time

### Sprint 2 Priorities

1. **Admin Dashboard**
   - Create `admin/DashboardPage.tsx`
   - Display journey count, word count, recent activity
   - Navigation to journey list

2. **Journey Management**
   - Create `admin/JourneyListPage.tsx` - table with filters
   - Create `admin/JourneyEditPage.tsx` - form to create/edit
   - Integrate `journeyService` methods
   - Add pagination

3. **Scenario & Word Forms**
   - Nested forms within journey editor
   - Dynamic add/remove scenario sections
   - Dynamic add/remove word sections

---

## 🔗 Related Documentation

- **Backend**: `/backend/README.md` - Go API server documentation
- **Functional Spec**: `/design/CORE.md` - Complete MVP specification
- **Requirements**: `/REQUIREMENT.md` - Business objectives and user stories
- **AI Strategy**: `.github/copilot-instructions.md` - AI coding agent guidelines

---

## 📞 Quick Help

**Issue**: Dev server won't start  
**Check**: Port 5173 available? Run `lsof -ti:5173 | xargs kill -9`

**Issue**: Can't connect to backend  
**Check**: Backend running on 8080? Run `cd ../backend && make run`

**Issue**: Type errors everywhere  
**Check**: Dependencies installed? Run `npm install`

**Issue**: Build fails  
**Check**: Type errors? Run `npm run type-check` first

---

**Last Updated**: 2025-10-05 (Sprint 0 - Foundation Complete)  
**Maintained By**: learng Development Team  
**Status**: ✅ Ready for Sprint 1 (Authentication UI)
