// API Response Types
export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
}

// User Types
export interface User {
  id: string;
  email: string;
  displayName: string;
  role: 'admin' | 'learner';
  createdAt: string;
  updatedAt: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  displayName: string;
  role: 'admin' | 'learner';
}

// Journey Types
export interface Journey {
  id: string;
  title: string;
  description: string;
  sourceLanguage: string;
  targetLanguage: string;
  status: 'draft' | 'published' | 'archived';
  createdBy: string;
  createdAt: string;
  updatedAt: string;
  scenarioCount?: number;
  wordCount?: number;
}

export interface JourneyListResponse {
  journeys: Journey[];
  total: number;
  page: number;
  limit: number;
}

export interface CreateJourneyRequest {
  title: string;
  description: string;
  sourceLanguage: string;
  targetLanguage: string;
}

// Scenario Types
export interface Scenario {
  id: string;
  journeyId: string;
  title: string;
  description: string;
  displayOrder: number;
  createdAt: string;
  updatedAt: string;
  wordCount?: number;
  words?: Word[];
  quiz?: Quiz;
}

export interface CreateScenarioRequest {
  journeyId: string;
  title: string;
  description: string;
  displayOrder: number;
}

// Word Types
export interface Word {
  id: string;
  scenarioId: string;
  targetText: string;
  sourceText: string;
  displayOrder: number;
  imageUrl: string | null;
  audioUrl: string | null;
  generationMethod: 'manual' | 'ai_image' | 'ai_audio' | 'ai_both';
  createdAt: string;
  updatedAt: string;
  masteryLevel?: 'new' | 'learning' | 'review' | 'mastered';
}

export interface CreateWordRequest {
  scenarioId: string;
  targetText: string;
  sourceText: string;
  displayOrder: number;
}

export interface UpdateWordRequest {
  targetText?: string;
  sourceText?: string;
  displayOrder?: number;
  imageUrl?: string;
  audioUrl?: string;
}

// Media Types
export interface MediaUploadResponse {
  url: string;
  filename: string;
  size: number;
  mimeType: string;
  duration?: number;
}

// Quiz Types
export interface Quiz {
  id: string;
  scenarioId: string;
  title: string;
  passThreshold: number;
  createdAt: string;
  updatedAt: string;
  questions?: QuizQuestion[];
}

export interface QuizQuestion {
  id: string;
  quizId: string;
  wordId: string;
  questionType: 'multiple_choice' | 'audio_match' | 'image_match';
  questionText: string;
  correctAnswer: string;
  options: string[];
  displayOrder: number;
}

export interface QuizAttempt {
  id: string;
  userId: string;
  quizId: string;
  score: number;
  totalQuestions: number;
  correctAnswers: number;
  answers: QuizAnswer[];
  completedAt: string;
}

export interface QuizAnswer {
  questionId: string;
  answer: string;
  isCorrect: boolean;
  correctAnswer?: string;
}

export interface SubmitQuizRequest {
  answers: {
    questionId: string;
    answer: string;
  }[];
}

export interface SubmitQuizResponse {
  attemptId: string;
  score: number;
  totalQuestions: number;
  correctAnswers: number;
  passed: boolean;
  feedback: QuizAnswer[];
}

// Progress Types
export interface LearnerProgress {
  id: string;
  userId: string;
  wordId: string;
  masteryLevel: 'new' | 'learning' | 'review' | 'mastered';
  viewCount: number;
  lastViewedAt: string;
  createdAt: string;
  updatedAt: string;
}

export interface UpdateProgressRequest {
  wordId: string;
}

export interface JourneyProgress {
  totalWords: number;
  viewedWords: number;
  masteredWords: number;
  completionPercentage: number;
}

export interface LearnerJourney extends Journey {
  progress: JourneyProgress;
}

export interface LearnerJourneysResponse {
  journeys: LearnerJourney[];
}
