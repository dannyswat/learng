# learng Frontend

React + TypeScript + Vite frontend for the learng language learning platform.

## 🚀 Quick Start

### Prerequisites
- Node.js 18+ and npm
- Backend API running on `http://localhost:8080`

### Installation

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The app will be available at `http://localhost:5173`

## 📁 Project Structure

```
frontend/
├── src/
│   ├── components/          # React components
│   │   ├── admin/          # Admin-specific components
│   │   ├── learner/        # Learner-specific components
│   │   └── shared/         # Shared/reusable components
│   ├── pages/              # Page components
│   ├── hooks/              # Custom React hooks
│   ├── services/           # API service layer
│   ├── contexts/           # React contexts (auth, etc.)
│   ├── types/              # TypeScript type definitions
│   ├── utils/              # Utility functions
│   ├── App.tsx             # Main app component
│   ├── main.tsx            # Entry point
│   └── index.css           # Global styles
├── public/                 # Static assets
├── index.html              # HTML template
├── vite.config.ts          # Vite configuration
├── tailwind.config.js      # Tailwind CSS configuration
└── tsconfig.json           # TypeScript configuration
```

## 🛠️ Development

### Available Scripts

```bash
# Development server with hot reload
npm run dev

# Type checking (no build)
npm run type-check

# Build for production
npm run build

# Preview production build locally
npm run preview

# Lint code
npm run lint
```

### API Proxy Configuration

The dev server automatically proxies API requests to avoid CORS issues:

- `/api/*` → `http://localhost:8080/api/*`
- `/uploads/*` → `http://localhost:8080/uploads/*`

This is configured in `vite.config.ts` and means:
- No CORS configuration needed
- Same URLs work in dev and production
- Frontend appears to be on same origin as backend

### Adding New Components

Components are organized by domain:

- **Admin components**: `/src/components/admin/` - Journey editor, word management, media upload
- **Learner components**: `/src/components/learner/` - Word cards, quiz UI, progress tracking
- **Shared components**: `/src/components/shared/` - Buttons, inputs, modals, etc.

Example component structure:
```tsx
import React from 'react';

interface MyComponentProps {
  title: string;
  onAction: () => void;
}

export const MyComponent: React.FC<MyComponentProps> = ({ title, onAction }) => {
  return (
    <div className="card">
      <h2>{title}</h2>
      <button onClick={onAction}>Click me</button>
    </div>
  );
};
```

### Using the API Service

All API calls go through service modules in `/src/services/`:

```tsx
import { journeyService } from '@/services/journey.service';

// In a component
const fetchJourneys = async () => {
  const data = await journeyService.getJourneys({ status: 'published' });
  setJourneys(data.journeys);
};
```

Authentication is handled automatically via Axios interceptors (JWT token from localStorage).

### Custom Hooks

Pre-built hooks for common tasks:

```tsx
import { useMediaUpload } from '@/hooks/useMediaUpload';
import { useAudioRecorder } from '@/hooks/useAudioRecorder';

// Upload image
const { uploadImage, isUploading, error } = useMediaUpload();
const url = await uploadImage(file, wordId);

// Record audio
const { startRecording, stopRecording, audioBlob } = useAudioRecorder();
```

## 🎨 Styling

### Tailwind CSS

This project uses Tailwind CSS for styling. Common patterns:

```tsx
// Button variants (from shared components)
<Button variant="primary">Save</Button>
<Button variant="outline">Cancel</Button>
<Button variant="danger">Delete</Button>

// Custom Tailwind classes
<div className="flex items-center justify-between p-4 bg-white rounded-lg shadow">
  {/* content */}
</div>
```

### Color Palette

Primary colors are defined in `tailwind.config.js`:

- `primary-50` to `primary-900` - Main brand colors (blue)
- Use semantic classes: `bg-primary-600`, `text-primary-700`, `border-primary-500`

### Pre-built Utility Classes

Defined in `index.css`:

- `.btn` - Base button styles
- `.btn-primary`, `.btn-secondary`, `.btn-outline` - Button variants
- `.input` - Input field styles
- `.card` - Card container styles

## 🔐 Authentication

### Using AuthContext

```tsx
import { useAuth } from '@/contexts/AuthContext';

const MyComponent = () => {
  const { user, isAuthenticated, login, logout } = useAuth();
  
  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }
  
  return <div>Welcome, {user?.displayName}!</div>;
};
```

### Protected Routes

```tsx
<Route
  path="/admin/*"
  element={
    <ProtectedRoute requireRole="admin">
      <AdminDashboard />
    </ProtectedRoute>
  }
/>
```

## 🧪 Testing

(To be implemented in Sprint 6)

Testing strategy:
- Unit tests: React Testing Library for components
- Integration tests: User flow scenarios
- E2E tests: Playwright for critical paths

## 📦 Building for Production

### Build

```bash
npm run build
```

This creates optimized static files in `/dist` directory:
- Minified JavaScript bundles
- Optimized images and assets
- Generated `index.html`

### Preview Production Build

```bash
npm run preview
```

Serves the production build locally at `http://localhost:4173`

### Deployment

The production build is designed to be served by the Go backend:

1. Frontend build output (`/dist`) is copied to backend's `/static` directory
2. Go backend serves these files via Echo's static file handler
3. Single-domain deployment (no CORS issues)

See `design/CORE.md` for full deployment architecture.

## 🔧 Configuration

### Environment Variables

Create `.env.local` for local development (not committed):

```bash
# Not currently needed - all config is in vite.config.ts
# Future: API_BASE_URL=http://localhost:8080
```

### TypeScript Path Aliases

`@/*` is aliased to `/src/*`:

```tsx
import { Button } from '@/components/shared/Button';
import { useAuth } from '@/contexts/AuthContext';
import type { User } from '@/types/api.types';
```

## 📚 Key Dependencies

- **React 18** - UI library
- **TypeScript** - Type safety
- **Vite** - Build tool and dev server
- **React Router v6** - Client-side routing
- **Axios** - HTTP client with interceptors
- **Tailwind CSS** - Utility-first styling
- **clsx** - Conditional class names
- **lucide-react** - Icon library

## 🐛 Troubleshooting

### Port 5173 already in use

```bash
# Kill the process
lsof -ti:5173 | xargs kill -9

# Or change port in vite.config.ts
server: { port: 5174 }
```

### Backend connection refused

Ensure Go backend is running on port 8080:
```bash
cd ../backend
make run
```

### Type errors after npm install

```bash
npm run type-check
```

If errors persist, delete `node_modules` and reinstall:
```bash
rm -rf node_modules package-lock.json
npm install
```

### CORS errors

This shouldn't happen with Vite proxy, but if you see CORS errors:
1. Verify proxy config in `vite.config.ts`
2. Check backend is running on `http://localhost:8080`
3. Restart both frontend and backend servers

## 🚦 Development Workflow

### Typical Session

```bash
# Terminal 1: Backend
cd backend
make run

# Terminal 2: Frontend
cd frontend
npm run dev
```

Access app at `http://localhost:5173`

### Making Changes

1. Edit component files in `/src/components/`
2. Changes hot-reload automatically (no page refresh needed)
3. Type errors show in terminal and editor
4. Check network tab for API calls

### Adding New Features

1. Define TypeScript types in `/src/types/api.types.ts`
2. Add API service methods in `/src/services/`
3. Create component in appropriate folder
4. Wire up routing in `App.tsx` if needed
5. Test in browser

## 📖 Next Steps

### Sprint 1 (Current Phase)
- ✅ Project structure created
- ✅ Authentication flow implemented
- ⏳ Complete login/register pages
- ⏳ Add error handling and loading states

### Sprint 2
- Implement admin journey management UI
- Create scenario and word forms
- Build journey list and detail pages

### Sprint 3
- Media upload component
- Audio recorder component
- Image preview and cropping

### Sprint 4
- Learner journey browsing
- WordCard component with animations
- Card navigation (swipe, keyboard)

### Sprint 5
- Quiz UI implementation
- Quiz results and feedback
- Progress tracking visualization

## 🤝 Contributing

(Internal team project - follow sprint board on GitHub Projects)

## 📄 License

Proprietary - learng Team 2025
