package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DatabasePath string
	JWTSecret    string
	UploadDir    string
	StaticDir    string // Frontend build directory (empty in dev)
	MaxImageSize int64  // bytes
	MaxAudioSize int64  // bytes
}

func Load() (*Config, error) {
	// Load .env file if it exists (ignore error if not found)
	_ = godotenv.Load()

	cfg := &Config{
		Port:         getEnv("PORT", "8080"),
		DatabasePath: getEnv("DB_PATH", "./learng.db"),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		UploadDir:    getEnv("UPLOAD_DIR", "./uploads"),
		StaticDir:    getEnv("STATIC_DIR", ""),                   // Empty in dev, set in production
		MaxImageSize: getEnvInt64("MAX_IMAGE_SIZE", 5*1024*1024), // 5MB default
		MaxAudioSize: getEnvInt64("MAX_AUDIO_SIZE", 2*1024*1024), // 2MB default
	}

	// Validate required fields
	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return fallback
}
