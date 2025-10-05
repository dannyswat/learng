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
	"github.com/learng/backend/internal/models"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	_, err = initDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

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

	api.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "pong",
		})
	})

	// TODO: Register route handlers here
	// Example:
	// authHandler := handlers.NewAuthHandler(db, cfg)
	// api.POST("/auth/register", authHandler.Register)
	// api.POST("/auth/login", authHandler.Login)

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
