package services

import (
	"errors"
	"time"

	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/repository"
	"github.com/learng/backend/internal/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// RegisterRequest contains the data needed to register a new user
type RegisterRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
}

// LoginRequest contains the data needed to log in
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse contains the user and token after authentication
type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

// Register creates a new user account
func (s *AuthService) Register(req RegisterRequest) (*AuthResponse, error) {
	// Validate input
	if !utils.ValidateEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	valid, msg := utils.ValidatePassword(req.Password)
	if !valid {
		return nil, errors.New(msg)
	}

	if !utils.ValidateRole(req.Role) {
		return nil, errors.New("invalid role")
	}

	// Check if user already exists
	exists, err := s.userRepo.Exists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		DisplayName:  req.DisplayName,
		Role:         req.Role,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// Don't return password hash
	user.PasswordHash = ""

	return &AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// Login authenticates a user and returns a token
func (s *AuthService) Login(req LoginRequest) (*AuthResponse, error) {
	// Validate input
	if !utils.ValidateEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Role, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// Don't return password hash
	user.PasswordHash = ""

	return &AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Don't return password hash
	user.PasswordHash = ""

	return user, nil
}
