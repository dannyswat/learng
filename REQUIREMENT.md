# learng – Business Requirements Document

## 1. Executive Summary
**learng** is a web application designed to help children learn a new language through an engaging, AI-powered multimedia experience. The platform presents vocabulary through pictures/videos with pronunciation audio, organizes content into themed scenarios, tests retention through quizzes, and tracks individual learning progress. An intuitive admin interface allows educators to create content by simply entering words—the system automatically generates supporting media using generative AI.

## 2. Business Objectives
| Objective | Success Metric | Timeline |
|-----------|----------------|----------|
| Reduce content creation time by 80% vs manual curation | Admin can publish 50+ word journey in <2 hours | MVP |
| Engage learners with high retention | 70%+ scenario completion rate | Month 3 |
| Scale content library efficiently | Support 5+ language pairs within 6 months | Phase 2 |
| Provide measurable learning outcomes | Track mastery progression per learner | MVP |
| Maintain cost-effective operations | AI generation cost <$0.10 per word (image+audio) | Ongoing |

## 3. Target Users & Personas

### 3.1 Primary Users (Learners)
- **Age Range**: 5-12 years old (primary), with potential expansion to teens
- **Learning Context**: Self-paced home learning, classroom supplement, travel preparation
- **Device Usage**: Primarily tablets/computers; responsive web interface
- **Attention Span**: 10-15 minute sessions typical
- **Needs**: Visual learning, clear audio, immediate feedback, sense of progress

### 3.2 Secondary Users (Admins/Educators)
- **Profile**: Teachers, tutors, parents creating custom curricula
- **Technical Skill**: Non-technical; require intuitive interface
- **Time Constraints**: Limited prep time; need rapid content creation
- **Needs**: Bulk operations, content reuse, learner progress visibility

### 3.3 Tertiary Users (Parents - Future Phase)
- **Profile**: Monitoring child's progress
- **Needs**: High-level dashboards, time limits, content appropriateness assurance

## 4. Core User Journeys

### 4.1 Learner Journey
1. **Discovery**: Select a journey by theme (e.g., "At the Park", "Food & Dining")
2. **Learning**: View cards sequentially—each showing:
   - Large, colorful image or short video
   - Target language word/phrase (large, readable font)
   - Optional native language translation
   - Pronunciation audio (auto-play or tap-to-play)
3. **Navigation**: Move forward/backward through cards via buttons, swipe, or keyboard
4. **Assessment**: After completing a scenario set, take a short quiz (5-10 questions)
5. **Feedback**: See immediate results, identify words needing review
6. **Progress**: View completion status and mastery levels

### 4.2 Admin Journey
1. **Content Planning**: Create a journey structure (title, description, target language)
2. **Scenario Setup**: Define scenarios within the journey (e.g., "Colors", "Family Members")
3. **Word Entry**: Input words/phrases via form or bulk paste from spreadsheet
4. **AI Generation**: System automatically generates image and audio for each word (asynchronous)
5. **Review & Adjust**: Preview generated media; regenerate if needed; optionally upload custom media
6. **Publishing**: Activate journey for learners
7. **Monitoring**: Review aggregated analytics (completion rates, difficult words, learner progress)

## 5. Key Features (Business Level)

### 5.1 Learner Experience
- **Rich Multimedia Cards**: Image/video + pronunciation for every vocabulary item
- **Intuitive Navigation**: Clear next/previous controls; progress indicator
- **Interactive Quizzes**: Multiple choice, listening comprehension, image matching
- **Progress Tracking**: Visual indicators of journey completion and word mastery
- **Accessibility**: Keyboard navigation, alt text, adjustable text size

### 5.2 Admin/Educator Tools
- **Simple Content Creation**: Text-only input; AI handles media generation
- **Bulk Operations**: Import word lists from CSV or clipboard
- **Content Organization**: Hierarchical structure (Journey → Scenario → Words)
- **Media Management**: Review, regenerate, or manually replace AI-generated assets
- **Analytics Dashboard**: Track learner engagement, quiz performance, content effectiveness
- **Draft/Publish Workflow**: Test content before making it live

### 5.3 AI-Powered Automation
- **Automatic Image Generation**: Context-appropriate, kid-safe illustrations
- **Pronunciation Audio**: Natural-sounding TTS in target language
- **Safety Filtering**: Automated moderation to ensure child-appropriate content
- **Cost Optimization**: Caching and reuse of similar prompts

### 5.4 Progress & Learning Intelligence
- **Individual Tracking**: Per-learner mastery levels for each word
- **Adaptive Review**: Flag words with low quiz performance for review
- **Journey History**: Record of completed scenarios and quiz scores
- **Future**: Spaced repetition scheduling, personalized difficulty adjustment

## 6. Success Criteria

### 6.1 User Engagement (Learners)
- 70%+ of started journeys reach first quiz
- Average session duration: 10-15 minutes
- 60%+ learners return within 7 days
- Quiz pass rate (≥70% correct): 75%+ of attempts

### 6.2 Content Efficiency (Admins)
- Journey creation time: <2 hours for 50-word scenario set
- AI generation success rate: >95% (ready without manual intervention)
- Admin satisfaction with media quality: >80% approve on first generation
- Time to publish new journey: <1 day from concept to learner availability

### 6.3 Technical Performance
- Card load time: <2 seconds (after initial journey load)
- Audio playback latency: <200ms from button press
- System availability: 99%+ uptime
- AI generation queue time: <30 seconds median per asset

### 6.4 Business Viability
- AI generation cost: <$0.10 per word (combined image + audio)
- Scalability: Support 1,000+ concurrent learners
- Content reuse rate: 40%+ of words appear in multiple scenarios (reducing marginal cost)

## 7. Business Constraints & Considerations

### 7.1 Budget
- MVP development: Bootstrap/limited initial investment
- Ongoing AI costs: Pay-per-use model initially; migrate to hybrid self-hosted if scale justifies
- Infrastructure: Cloud-based, auto-scaling to match usage

### 7.2 Compliance & Safety
- **Child Safety**: All generated content must pass automated moderation
- **Data Privacy**: Minimal PII collection; comply with PDPO (Hong Kong), prepare for COPPA/GDPR-K if expanding
- **Parental Consent**: Implement consent workflow if required by jurisdiction
- **Content Licensing**: Ensure commercial usage rights for all AI-generated media

### 7.3 Market Positioning
- **Differentiator**: AI-powered rapid content creation (vs. manual curation platforms)
- **Target Market**: Initially Hong Kong/Asia-Pacific; English UI teaching Cantonese, Mandarin, or English as second language
- **Competition**: Duolingo (gamified but limited kid focus), Memrise, Rosetta Stone Kids (expensive)
- **Pricing Model (Future)**: Freemium (basic journeys free) or institutional licenses (schools/tutoring centers)

### 7.4 Regional Considerations (Hong Kong Base)
- **Hosting**: Prefer Asia-Pacific regions (Hong Kong, Singapore) for low latency
- **AI Providers**: Use globally accessible providers (Azure, Stability, ElevenLabs) with HK/Singapore endpoints
- **Languages**: Initial focus on Cantonese, Mandarin, English (bidirectional teaching pairs)
- **Cultural Sensitivity**: Ensure prompts and content respect local cultural norms

## 8. Out of Scope (Initial Release)

The following are explicitly deferred to post-MVP phases:
- Native mobile apps (iOS/Android) – MVP is responsive web only
- Real-time multiplayer or classroom collaboration features
- Advanced speech recognition for pronunciation practice
- Parental dashboard and controls
- Gamification (badges, leaderboards, rewards)
- Spaced repetition algorithms (basic review triggers only in MVP)
- Video generation (optional placeholder images for MVP)
- Advanced analytics (ML-driven insights, A/B testing)
- Multi-tenancy / white-label deployments

## 9. Risks & Mitigation Strategies

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| AI-generated content inappropriate | High (child safety) | Medium | Multi-layer safety filters + admin approval queue option |
| High AI costs exceed budget | High (viability) | Medium | Implement caching, prompt optimization, hybrid self-hosting plan |
| Slow AI generation frustrates admins | Medium (adoption) | Medium | Async queue + placeholder content + status notifications |
| Low learner engagement/retention | High (product failure) | Medium | User testing with target age group; iterate on UI/UX |
| Regional AI service restrictions | Medium (availability) | Low-Medium | Multi-provider abstraction layer; fallback options |
| Data privacy violations | High (legal/reputation) | Low | Privacy-by-design; minimal data collection; compliance audit |
| Scope creep delays MVP | Medium (timeline) | High | Strict MVP feature fence; agile sprint discipline |

## 10. Roadmap & Phasing

### Phase 1: MVP (Months 1-3)
- Core learner journey (browse, view cards, take quizzes)
- Admin content creation (words, scenarios, journeys)
- AI generation (images via DALL·E/SDXL, audio via Azure TTS or ElevenLabs)
- Basic progress tracking (views, quiz scores, mastery levels)
- Single language pair demonstration
- Responsive web interface

### Phase 2: Enhancement (Months 4-6)
- Additional language pairs (3-5 total)
- Video generation for select scenarios
- Improved analytics dashboard
- Spaced repetition triggers
- Parent view (read-only progress monitoring)
- Performance optimizations (caching, CDN)

### Phase 3: Scale (Months 7-12)
- Speech recognition (pronunciation practice scoring)
- Gamification features (badges, streaks)
- Offline-first PWA
- Institutional admin panel (class management)
- Self-hosted AI models for cost reduction
- Advanced personalization (adaptive difficulty)

## 11. Open Questions Requiring Stakeholder Input

1. **Target Languages**: Which language pairs should MVP support? (Recommend: English↔Cantonese as pilot)
2. **Age Range Refinement**: Focus on 5-8 or expand to 5-12? (Affects UI complexity, content tone)
3. **Deployment Timeline**: Hard launch date or iterative soft launch?
4. **Budget Cap**: Monthly AI spend limit for MVP phase?
5. **Content Ownership**: Who provides initial seed journeys? (Internal creation vs. educator partnerships?)
6. **Monetization**: Free pilot, freemium, or institutional sales from day one?
7. **Compliance**: COPPA/GDPR-K requirements if launching beyond Hong Kong?

## 12. Agile Development Approach

**Functional specifications and technical details will be defined iteratively** through:
- **User Stories**: Defined per sprint with acceptance criteria
- **Sprint Planning**: 2-week cycles with prioritized backlog
- **Technical Specs**: Created just-in-time for implementation
- **Continuous Feedback**: User testing with target demographic after each major feature
- **Architecture Decisions Records (ADRs)**: Document key technical choices separately

This document serves as the **strategic anchor**; detailed requirements evolve in the backlog and sprint artifacts.

## 13. Appendices

### A. Related Documents
- `AI ANALYSIS.md` – Detailed provider evaluation and regional considerations
- `.github/copilot-instructions.md` – AI coding agent context
- Future: `ARCHITECTURE.md`, `API_SPEC.md`, `USER_STORIES.md` (created as needed)

### B. Glossary
- **Journey**: A curated learning path consisting of multiple scenarios
- **Scenario**: A thematic grouping of words (e.g., "Colors", "Family")
- **Card**: A single vocabulary item presentation (image + word + audio)
- **Mastery Level**: Learner's proficiency with a word (NEW, LEARNING, REVIEW, MASTERED)
- **Admin**: Content creator/educator role
- **Learner**: End user (child) consuming content

---

**Document Ownership**: Product Lead  
**Last Updated**: 2025-10-05  
**Next Review**: After MVP user testing completion
