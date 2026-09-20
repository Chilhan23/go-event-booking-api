package auth

import (
	"context"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) (*UserResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error)
}

type service struct {
	repo      Repository
	jwtSecret string
}

func NewService(repo Repository, jwtSecret string) Service {
	return &service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *service) generateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), 
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.jwtSecret))
}

// Register handles user registration
func (s *service) Register(ctx context.Context, req *RegisterRequest) (*UserResponse, error) {
	// Check if username is already taken
	existUser, _ := s.repo.FindByUsername(ctx, req.Username)
	if existUser != nil {
		return nil, errors.New("username already taken")
	}

	// Check if email is already registered
	existEmail, _ := s.repo.FindByEmail(ctx, req.Email)
	if existEmail != nil {
		return nil, errors.New("email already registered")
	}

	// Hash plain text password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashPassword),
	}

	// Persist user to database
	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// Login authenticates a user and returns a signed JWT token
func (s *service) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	// Find user by username
	user, err := s.repo.FindByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return nil, errors.New("invalid username or password")
	}

	// Verify hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User: UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}


