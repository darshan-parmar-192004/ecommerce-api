package services

import (
	"context"
	"fmt"
	"time"

	"backend/internal/cache"
	"backend/internal/config"
	"backend/internal/constants"
	"backend/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	CustomerID string `json:"customer_id"`
	Email      string `json:"email"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo   *models.CustomerRepository
	cache  cache.Service
	secret string
	expiry time.Duration
	issuer string
}

func NewAuthService(repo *models.CustomerRepository, cacheSvc cache.Service) *AuthService {
	cfg, _ := config.Load()
	return &AuthService{
		repo:   repo,
		cache:  cacheSvc,
		secret: cfg.JWTSecret,
		expiry: time.Duration(cfg.JWTExpiry) * time.Hour,
		issuer: cfg.JWTIssuer,
	}
}

func (s *AuthService) Register(ctx context.Context, customerID, email, name, password, country, phone string) (*models.Customer, error) {
	if len(password) < 8 {
		return nil, fmt.Errorf(constants.ErrAuthWeakPassword)
	}

	existing, err := s.repo.FindByEmail(ctx, email)
	if err == nil && existing != nil {
		return nil, fmt.Errorf(constants.ErrAuthEmailExists)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	customer := models.Customer{
		CustomerID:   customerID,
		Email:        email,
		Name:         name,
		Country:      country,
		Phone:        phone,
		PasswordHash: string(hashed),
		CreatedAt:    time.Now(),
		Status:       "active",
	}

	created, err := s.repo.Create(ctx, customer)
	if err != nil {
		return nil, err
	}

	created.PasswordHash = ""
	return created, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, *models.Customer, error) {
	customer, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return "", nil, fmt.Errorf(constants.ErrAuthInvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password)); err != nil {
		return "", nil, fmt.Errorf(constants.ErrAuthInvalidCredentials)
	}

	now := time.Now()
	jti := uuid.New().String()
	claims := &Claims{
		CustomerID: customer.CustomerID,
		Email:      customer.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    s.issuer,
			ID:        jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", nil, fmt.Errorf("failed to sign token: %w", err)
	}

	customer.PasswordHash = ""
	return tokenString, customer, nil
}

func (s *AuthService) Logout(ctx context.Context, tokenString string) error {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return err
	}

	ttl := time.Until(claims.ExpiresAt.Time)
	if ttl <= 0 {
		return nil
	}

	key := fmt.Sprintf(constants.CacheKeyTokenBlacklist, claims.ID)
	return s.cache.Set(ctx, key, "1", ttl)
}

func (s *AuthService) ValidateToken(ctx context.Context, tokenString string) (*Claims, error) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf(constants.CacheKeyTokenBlacklist, claims.ID)
	if _, err := s.cache.Get(ctx, key); err == nil {
		return nil, fmt.Errorf(constants.ErrAuthTokenExpired)
	}

	return claims, nil
}

func (s *AuthService) parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf(constants.ErrAuthTokenInvalid)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf(constants.ErrAuthTokenInvalid)
	}

	return claims, nil
}
