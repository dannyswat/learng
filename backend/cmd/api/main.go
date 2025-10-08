package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/learng/backend/internal/config"
	"github.com/learng/backend/internal/handlers"
	customMiddleware "github.com/learng/backend/internal/middleware"
	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/repository"
	"github.com/learng/backend/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	journeyRepo := repository.NewJourneyRepository(db)
	scenarioRepo := repository.NewScenarioRepository(db)
	wordRepo := repository.NewWordRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	journeyService := services.NewJourneyService(journeyRepo, scenarioRepo)
	scenarioService := services.NewScenarioService(scenarioRepo, journeyRepo)
	wordService := services.NewWordService(wordRepo, scenarioRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	journeyHandler := handlers.NewJourneyHandler(journeyService)
	scenarioHandler := handlers.NewScenarioHandler(scenarioService)
	wordHandler := handlers.NewWordHandler(wordService)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "healthy",
			"version": "1.0.0",
		})
	})

	// API routes (v1)
	api := e.Group("/api/v1")

	// Public routes (no authentication required)
	api.POST("/auth/register", authHandler.Register)
	api.POST("/auth/login", authHandler.Login)

	// Protected routes (authentication required)
	protected := api.Group("")
	protected.Use(customMiddleware.AuthMiddleware(cfg.JWTSecret))
	protected.GET("/auth/me", authHandler.GetMe)

	// Journey routes (admin only for create/update/delete)
	protected.GET("/journeys", journeyHandler.GetJourneys)
	protected.GET("/journeys/:id", journeyHandler.GetJourneyByID)
	protected.POST("/journeys", journeyHandler.CreateJourney)
	protected.PUT("/journeys/:id", journeyHandler.UpdateJourney)
	protected.DELETE("/journeys/:id", journeyHandler.DeleteJourney)

	// Scenario routes
	protected.POST("/scenarios", scenarioHandler.CreateScenario)
	protected.GET("/scenarios/:id", scenarioHandler.GetScenarioByID)
	protected.PUT("/scenarios/:id", scenarioHandler.UpdateScenario)
	protected.DELETE("/scenarios/:id", scenarioHandler.DeleteScenario)

	// Word routes
	protected.POST("/words", wordHandler.CreateWord)
	protected.GET("/words/:id", wordHandler.GetWordByID)
	protected.PUT("/words/:id", wordHandler.UpdateWord)
	protected.DELETE("/words/:id", wordHandler.DeleteWord)

	// Serve uploaded media files
	e.Static("/uploads", cfg.UploadDir)

	// Serve frontend static files (production only)
	if cfg.StaticDir != "" {
		log.Println("Serving static files from:", cfg.StaticDir)
		e.Static("/assets", cfg.StaticDir+"/assets")
		e.File("/favicon.ico", cfg.StaticDir+"/favicon.ico")
		// SPA fallback: serve index.html for all non-API/non-uploads routes
		e.File("/*", cfg.StaticDir+"/index.html")
	} else {
		log.Println("Running in development mode (no static files served)")
	}

	// Start server
	port := cfg.Port
	log.Printf("Starting server on port %s...\n", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDatabase(cfg *config.Config) (*gorm.DB, error) {
	// Ensure database directory exists
	if err := os.MkdirAll(cfg.UploadDir+"/images", 0755); err != nil {
		return nil, fmt.Errorf("failed to create images directory: %w", err)
	}
	if err := os.MkdirAll(cfg.UploadDir+"/audio", 0755); err != nil {
		return nil, fmt.Errorf("failed to create audio directory: %w", err)
	}

	// Open database connection
	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate models
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.Journey{},
		&models.Scenario{},
		&models.Word{},
		&models.Quiz{},
		&models.QuizQuestion{},
		&models.LearnerProgress{},
		&models.QuizAttempt{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}
