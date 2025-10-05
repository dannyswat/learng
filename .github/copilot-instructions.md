# AI Coding Agent Instructions for `learng`

## Project Mission
Build an AI-powered language learning web app for kids (ages 5-12) that automatically generates rich multimedia content (images + pronunciation audio) from simple word/phrase inputs, organizes vocabulary into themed journeys, and tracks individual learner mastery.

## Current State (Pre-Implementation)
- **Phase**: Requirements defined; tech stack selection pending
- **Documents**: `REQUIREMENT.md` (business requirements), `AI ANALYSIS.md` (AI provider comparison)
- **Region**: Hong Kong-based; target Asia-Pacific latency; initial languages: Cantonese/Mandarin/English pairs
- **No code yet**: First implementation pass should establish stack, directory structure, and scaffold MVP features

## Strategic Context (Read First)
- **Differentiator**: Admin enters words → AI auto-generates kid-safe images + TTS audio → learners consume via card-based UI
- **Safety-Critical**: All generated media must pass moderation (child audience)
- **Cost-Sensitive**: Target <$0.10/word generation cost; plan caching + eventual hybrid self-hosting
- **Agile**: Functional specs evolve iteratively; this doc + REQUIREMENT.md are strategic anchors

## Architecture Principles (Intended)
1. **Modular AI Adapters**: Abstract image/audio/LLM providers behind interfaces to enable hot-swapping (Azure → Stability, ElevenLabs → Azure TTS, etc.)
2. **Async-First Generation**: Queue AI jobs; never block admin UI on generation latency
3. **Safety Pipeline**: Input validation → Prompt refinement (LLM) → Generation → Moderation → Storage
4. **Progressive Enhancement**: Responsive web (MVP) → PWA offline (Phase 2) → native apps (future)
5. **Privacy-by-Design**: Minimal PII; learner progress tracked with pseudonymous IDs

## Recommended Tech Stack (Pending Confirmation)
**Awaiting stakeholder decision; default assumption for scaffolding:**
- **Frontend**: Next.js (App Router) + React + TypeScript + Tailwind CSS
- **Backend**: Node.js + Express/Fastify (or Next.js API routes) + TypeScript
- **Database**: PostgreSQL (via Prisma ORM) for relational data + audit log
- **Queue**: BullMQ (Redis-backed) for async AI job processing
- **Storage**: Cloud object storage (Azure Blob or S3-compatible) + CDN
- **AI Providers (MVP)**: Azure OpenAI (image + LLM), Azure TTS or ElevenLabs (audio), Azure Content Safety (moderation)
- **Deployment**: Docker-compose (local) → Azure Container Apps or Vercel + separate worker service (production)

## Directory Structure (Proposed Once Stack Chosen)
```
learng/
├── apps/
│   ├── web/            # Next.js frontend (learner + admin UI)
│   └── worker/         # AI generation job processor
├── packages/
│   ├── ai-adapters/    # Provider abstraction layer
│   ├── db/             # Prisma schema + migrations
│   └── shared/         # Common types, utils, constants
├── docs/               # Architecture decisions, API specs
├── scripts/            # Setup, seed data, cost analysis
└── .github/            # CI workflows, this file
```

## Core Data Model (High-Level)
- **Journey** (1) → **Scenarios** (N) → **Words** (N) → **MediaAssets** (N per word)
- **User** (learner/admin roles) → **Mastery** (user × word progress)
- **QuizSession** → **QuizQuestions** (quiz attempts + scoring)
- **Events** (append-only audit log)
- **GenerationJobs** (queue tasks: pending → processing → completed|failed)

## AI Integration Patterns (Critical)
**All AI calls must follow this flow:**
1. **Prompt Refinement**: Admin word → LLM generates descriptive, kid-safe visual prompt + alt text
2. **Enqueue Jobs**: Create `generation_jobs` records (image + audio) with status=pending
3. **Worker Polls Queue**: Fetch job → call provider adapter → store result/error
4. **Moderation Gate**: Before marking asset ready, run safety check; flag/quarantine unsafe outputs
5. **Status Updates**: Emit events for admin UI live status (polling or WebSocket)

**Provider Adapter Interface (conceptual TypeScript):**
```typescript
interface ImageGenAdapter {
  generate(req: ImageGenRequest): Promise<ImageGenResult>;
}
interface AudioTTSAdapter {
  synthesize(req: AudioRequest): Promise<AudioResult>;
}
// Implementations: AzureImageAdapter, StabilityImageAdapter, ElevenLabsTTSAdapter, AzureTTSAdapter
```

## Testing Strategy (MVP Baseline)
- **Unit**: AI adapter logic, prompt templates, mastery calculation
- **Integration**: API endpoints (journey CRUD, quiz scoring), job queue processing
- **E2E (manual initially)**: Admin creates word → verify AI generation → learner views card → takes quiz
- **Safety Tests**: Submit edge-case prompts; verify moderation catches inappropriate outputs

## Security & Compliance (Non-Negotiable)
- Store API keys in environment variables / Azure Key Vault (never commit)
- Hash passwords (bcrypt); use JWT or session cookies for auth
- Rate-limit admin regeneration endpoints to prevent cost abuse
- Log all AI requests (prompt + provider + cost estimate) for audit
- Anonymize learner analytics (no direct PII in aggregate reports)

## Conventions (Evolving)
**Once code is written, update this section with discovered patterns:**
- Error handling: Structured errors with codes (e.g., `AI_GENERATION_FAILED`, `MODERATION_FLAGGED`)
- Logging: Structured JSON (trace IDs per request/job)
- Config: `.env.example` with all required vars; validate on startup with Zod/Joi
- Naming: `camelCase` (TS/JS), `snake_case` (DB columns), `kebab-case` (files/folders)

## When Adding New Features
1. Check if AI provider abstraction is affected → update interface first
2. Add DB migration if schema changes (version journeys to avoid breaking published content)
3. Update `REQUIREMENT.md` if business scope shifts
4. Add cost telemetry if new AI call introduced (track tokens/chars/images)
5. Consider safety implications (new user input? → sanitize; new AI output? → moderate)

## Cost Control Techniques (Apply Everywhere)
- Cache AI outputs: hash(prompt+style) → check DB before calling provider
- Batch operations: Admin bulk-adds 50 words → enqueue all jobs, respect rate limits
- Fallback tiers: Expensive model fails/quotas exhausted → cheaper fallback (e.g., GPT-4o → GPT-4o-mini)
- Monitoring: Daily cost counter; soft alert at 80% budget; hard stop at 100%

## Open Questions to Resolve Before Deep Coding
- Exact language pairs for MVP? (Recommend: English↔Cantonese pilot)
- Azure subscription + OpenAI quota approved?
- Admin auth: OAuth (Google) or simple email/password?
- Hosting preference: Vercel, Azure Container Apps, or other?
- Budget ceiling for MVP phase AI spend?

## Development Workflow (Once Stack Established)
```bash
# Local setup (example commands, adjust per stack)
npm install                     # Install dependencies
npx prisma migrate dev          # Run DB migrations
npm run dev                     # Start Next.js + worker (concurrently)
npm run test                    # Run test suite
npm run lint                    # ESLint + Prettier
```

## CI/CD Expectations
- **PR Checks**: Lint, type-check, unit tests, build verification
- **Main Branch**: Auto-deploy to staging environment
- **Production**: Manual approval gate (until stable)

## Related Documents (Read for Context)
- `REQUIREMENT.md` – Business goals, user journeys, success metrics, phasing
- `AI ANALYSIS.md` – Provider comparison, safety pipeline, cost strategies, regional constraints
- Future: `ARCHITECTURE.md` (system design), `API_SPEC.md` (OpenAPI/GraphQL schema), `USER_STORIES.md` (sprint backlog)

## AI Agent Best Practices for This Project
- **Safety First**: When generating prompts or handling user input, always apply sanitization/validation
- **Cost Awareness**: Prefer cheaper models (e.g., gpt-4o-mini over gpt-4o) unless quality explicitly requires premium
- **Async by Default**: Never block HTTP responses on AI generation; use job queue + status polling
- **Regional Respect**: Suggest Asia-Pacific regions (Hong Kong, Singapore) for deployments/AI endpoints
- **Test AI Paths**: When adding adapter implementations, include mock/stub tests to avoid live API costs in CI

---
**Last Updated**: 2025-10-05 | **Next Update Trigger**: After tech stack decision + initial scaffold commit
